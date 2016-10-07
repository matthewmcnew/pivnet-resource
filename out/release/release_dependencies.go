package release

import (
	"fmt"
	"log"

	pivnet "github.com/pivotal-cf/go-pivnet"
	"github.com/pivotal-cf/pivnet-resource/metadata"
)

type ReleaseDependenciesAdder struct {
	logger      *log.Logger
	pivnet      releaseDependenciesAdderClient
	metadata    metadata.Metadata
	productSlug string
}

func NewReleaseDependenciesAdder(
	logger *log.Logger,
	pivnetClient releaseDependenciesAdderClient,
	metadata metadata.Metadata,
	productSlug string,
) ReleaseDependenciesAdder {
	return ReleaseDependenciesAdder{
		logger:      logger,
		pivnet:      pivnetClient,
		metadata:    metadata,
		productSlug: productSlug,
	}
}

//go:generate counterfeiter --fake-name ReleaseDependenciesAdderClient . releaseDependenciesAdderClient
type releaseDependenciesAdderClient interface {
	AddReleaseDependency(productSlug string, releaseID int, dependentReleaseID int) error
	GetRelease(productSlug string, releaseVersion string) (pivnet.Release, error)
}

func (rf ReleaseDependenciesAdder) AddReleaseDependencies(release pivnet.Release) error {
	rf.logger.Printf(
		"Adding release dependencies to %s/%d",
		rf.productSlug,
		release.ID,
	)

	for i, d := range rf.metadata.Dependencies {
		dependentReleaseID := d.Release.ID
		if dependentReleaseID == 0 {
			if d.Release.Version == "" || d.Release.Product.Slug == "" {
				return fmt.Errorf(
					"Either ReleaseID or release version and product slug must be provided for dependency[%d]",
					i,
				)
			}

			rf.logger.Printf(
				"Looking up dependent release ID for: %s/%d",
				d.Release.Product.Slug,
				d.Release.Version,
			)
			r, err := rf.pivnet.GetRelease(d.Release.Product.Slug, d.Release.Version)
			if err != nil {
				return err
			}
			dependentReleaseID = r.ID
		}

		rf.logger.Printf(
			"Adding dependent release with ID: %d to newly-created release",
			dependentReleaseID,
		)
		err := rf.pivnet.AddReleaseDependency(rf.productSlug, release.ID, dependentReleaseID)
		if err != nil {
			return err
		}
	}

	return nil
}
