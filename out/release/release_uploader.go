package release

import (
	"log"
	"path/filepath"

	"github.com/pivotal-cf-experimental/pivnet-resource/metadata"
	"github.com/pivotal-cf-experimental/pivnet-resource/pivnet"
)

type ReleaseUploader struct {
	s3          s3client
	pivnet      uploadClient
	logger      logging
	md5Summer   md5Summer
	metadata    metadata.Metadata
	skipUpload  bool
	sourcesDir  string
	productSlug string
}

type uploadClient interface {
	FindProductForSlug(slug string) (pivnet.Product, error)
	CreateProductFile(pivnet.CreateProductFileConfig) (pivnet.ProductFile, error)
	AddProductFile(productID int, releaseID int, productFileID int) error
}

type s3client interface {
	UploadFile(string) (string, error)
}

type md5Summer interface {
	SumFile(filepath string) (string, error)
}

func NewReleaseUploader(s3 s3client, pivnet uploadClient, logger logging, md5Summer md5Summer, metadata metadata.Metadata, skip bool, sourcesDir, productSlug string) ReleaseUploader {
	return ReleaseUploader{
		s3:          s3,
		pivnet:      pivnet,
		logger:      logger,
		md5Summer:   md5Summer,
		metadata:    metadata,
		skipUpload:  skip,
		sourcesDir:  sourcesDir,
		productSlug: productSlug,
	}
}

func (u ReleaseUploader) Upload(release pivnet.Release, exactGlobs []string) error {
	if u.skipUpload {
		u.logger.Debugf("File glob and s3_filepath_prefix not provided - skipping upload to s3")
		return nil
	}

	for _, exactGlob := range exactGlobs {
		fullFilepath := filepath.Join(u.sourcesDir, exactGlob)
		fileContentsMD5, err := u.md5Summer.SumFile(fullFilepath)
		if err != nil {
			log.Fatalln(err)
		}

		remotePath, err := u.s3.UploadFile(exactGlob)
		if err != nil {
			return err
		}

		product, err := u.pivnet.FindProductForSlug(u.productSlug)
		if err != nil {
			log.Fatalln(err)
		}

		filename := filepath.Base(exactGlob)

		var description string
		uploadAs := filename
		for _, f := range u.metadata.ProductFiles {
			if f.File == exactGlob {
				u.logger.Debugf("exact glob '%s' matches metadata file: '%s'\n", exactGlob, f.File)
				description = f.Description
				if f.UploadAs != "" {
					u.logger.Debugf("upload_as provided for exact glob: '%s' - uploading to remote filename: '%s' instead\n", exactGlob, f.UploadAs)
					uploadAs = f.UploadAs
				}
			} else {
				u.logger.Debugf("exact glob %s does not match metadata file: %s\n", exactGlob, f.File)
			}
		}

		u.logger.Debugf(
			"Creating product file: {product_slug: %s, filename: %s, aws_object_key: %s, file_version: %s, description: %s}\n",
			u.productSlug,
			uploadAs,
			remotePath,
			release.Version,
			description,
		)

		productFile, err := u.pivnet.CreateProductFile(pivnet.CreateProductFileConfig{
			ProductSlug:  u.productSlug,
			Name:         uploadAs,
			AWSObjectKey: remotePath,
			FileVersion:  release.Version,
			MD5:          fileContentsMD5,
			Description:  description,
		})
		if err != nil {
			return err
		}

		u.logger.Debugf(
			"Adding product file: {product_slug: %s, product_id: %d, filename: %s, product_file_id: %d, release_id: %d}\n",
			u.productSlug,
			product.ID,
			filename,
			productFile.ID,
			release.ID,
		)

		err = u.pivnet.AddProductFile(product.ID, release.ID, productFile.ID)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return nil
}