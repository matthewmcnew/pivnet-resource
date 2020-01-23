// Code generated by counterfeiter. DO NOT EDIT.
package infakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v4"
)

type FakePivnetClient struct {
	GetReleaseStub        func(productSlug string, version string) (pivnet.Release, error)
	getReleaseMutex       sync.RWMutex
	getReleaseArgsForCall []struct {
		productSlug string
		version     string
	}
	getReleaseReturns struct {
		result1 pivnet.Release
		result2 error
	}
	getReleaseReturnsOnCall map[int]struct {
		result1 pivnet.Release
		result2 error
	}
	AcceptEULAStub        func(productSlug string, releaseID int) error
	acceptEULAMutex       sync.RWMutex
	acceptEULAArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	acceptEULAReturns struct {
		result1 error
	}
	acceptEULAReturnsOnCall map[int]struct {
		result1 error
	}
	FileGroupsForReleaseStub        func(productSlug string, releaseID int) ([]pivnet.FileGroup, error)
	fileGroupsForReleaseMutex       sync.RWMutex
	fileGroupsForReleaseArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	fileGroupsForReleaseReturns struct {
		result1 []pivnet.FileGroup
		result2 error
	}
	fileGroupsForReleaseReturnsOnCall map[int]struct {
		result1 []pivnet.FileGroup
		result2 error
	}
	ImageReferencesForReleaseStub        func(productSlug string, releaseID int) ([]pivnet.ImageReference, error)
	imageReferencesForReleaseMutex       sync.RWMutex
	imageReferencesForReleaseArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	imageReferencesForReleaseReturns struct {
		result1 []pivnet.ImageReference
		result2 error
	}
	imageReferencesForReleaseReturnsOnCall map[int]struct {
		result1 []pivnet.ImageReference
		result2 error
	}
	HelmChartReferencesForReleaseStub        func(productSlug string, releaseID int) ([]pivnet.HelmChartReference, error)
	helmChartReferencesForReleaseMutex       sync.RWMutex
	helmChartReferencesForReleaseArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	helmChartReferencesForReleaseReturns struct {
		result1 []pivnet.HelmChartReference
		result2 error
	}
	helmChartReferencesForReleaseReturnsOnCall map[int]struct {
		result1 []pivnet.HelmChartReference
		result2 error
	}
	ProductFilesForReleaseStub        func(productSlug string, releaseID int) ([]pivnet.ProductFile, error)
	productFilesForReleaseMutex       sync.RWMutex
	productFilesForReleaseArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	productFilesForReleaseReturns struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	productFilesForReleaseReturnsOnCall map[int]struct {
		result1 []pivnet.ProductFile
		result2 error
	}
	ProductFileForReleaseStub        func(productSlug string, releaseID int, productFileID int) (pivnet.ProductFile, error)
	productFileForReleaseMutex       sync.RWMutex
	productFileForReleaseArgsForCall []struct {
		productSlug   string
		releaseID     int
		productFileID int
	}
	productFileForReleaseReturns struct {
		result1 pivnet.ProductFile
		result2 error
	}
	productFileForReleaseReturnsOnCall map[int]struct {
		result1 pivnet.ProductFile
		result2 error
	}
	ReleaseDependenciesStub        func(productSlug string, releaseID int) ([]pivnet.ReleaseDependency, error)
	releaseDependenciesMutex       sync.RWMutex
	releaseDependenciesArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	releaseDependenciesReturns struct {
		result1 []pivnet.ReleaseDependency
		result2 error
	}
	releaseDependenciesReturnsOnCall map[int]struct {
		result1 []pivnet.ReleaseDependency
		result2 error
	}
	DependencySpecifiersStub        func(productSlug string, releaseID int) ([]pivnet.DependencySpecifier, error)
	dependencySpecifiersMutex       sync.RWMutex
	dependencySpecifiersArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	dependencySpecifiersReturns struct {
		result1 []pivnet.DependencySpecifier
		result2 error
	}
	dependencySpecifiersReturnsOnCall map[int]struct {
		result1 []pivnet.DependencySpecifier
		result2 error
	}
	ReleaseUpgradePathsStub        func(productSlug string, releaseID int) ([]pivnet.ReleaseUpgradePath, error)
	releaseUpgradePathsMutex       sync.RWMutex
	releaseUpgradePathsArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	releaseUpgradePathsReturns struct {
		result1 []pivnet.ReleaseUpgradePath
		result2 error
	}
	releaseUpgradePathsReturnsOnCall map[int]struct {
		result1 []pivnet.ReleaseUpgradePath
		result2 error
	}
	UpgradePathSpecifiersStub        func(productSlug string, releaseID int) ([]pivnet.UpgradePathSpecifier, error)
	upgradePathSpecifiersMutex       sync.RWMutex
	upgradePathSpecifiersArgsForCall []struct {
		productSlug string
		releaseID   int
	}
	upgradePathSpecifiersReturns struct {
		result1 []pivnet.UpgradePathSpecifier
		result2 error
	}
	upgradePathSpecifiersReturnsOnCall map[int]struct {
		result1 []pivnet.UpgradePathSpecifier
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) GetRelease(productSlug string, version string) (pivnet.Release, error) {
	fake.getReleaseMutex.Lock()
	ret, specificReturn := fake.getReleaseReturnsOnCall[len(fake.getReleaseArgsForCall)]
	fake.getReleaseArgsForCall = append(fake.getReleaseArgsForCall, struct {
		productSlug string
		version     string
	}{productSlug, version})
	fake.recordInvocation("GetRelease", []interface{}{productSlug, version})
	fake.getReleaseMutex.Unlock()
	if fake.GetReleaseStub != nil {
		return fake.GetReleaseStub(productSlug, version)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getReleaseReturns.result1, fake.getReleaseReturns.result2
}

func (fake *FakePivnetClient) GetReleaseCallCount() int {
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	return len(fake.getReleaseArgsForCall)
}

func (fake *FakePivnetClient) GetReleaseArgsForCall(i int) (string, string) {
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	return fake.getReleaseArgsForCall[i].productSlug, fake.getReleaseArgsForCall[i].version
}

func (fake *FakePivnetClient) GetReleaseReturns(result1 pivnet.Release, result2 error) {
	fake.GetReleaseStub = nil
	fake.getReleaseReturns = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) GetReleaseReturnsOnCall(i int, result1 pivnet.Release, result2 error) {
	fake.GetReleaseStub = nil
	if fake.getReleaseReturnsOnCall == nil {
		fake.getReleaseReturnsOnCall = make(map[int]struct {
			result1 pivnet.Release
			result2 error
		})
	}
	fake.getReleaseReturnsOnCall[i] = struct {
		result1 pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) AcceptEULA(productSlug string, releaseID int) error {
	fake.acceptEULAMutex.Lock()
	ret, specificReturn := fake.acceptEULAReturnsOnCall[len(fake.acceptEULAArgsForCall)]
	fake.acceptEULAArgsForCall = append(fake.acceptEULAArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("AcceptEULA", []interface{}{productSlug, releaseID})
	fake.acceptEULAMutex.Unlock()
	if fake.AcceptEULAStub != nil {
		return fake.AcceptEULAStub(productSlug, releaseID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.acceptEULAReturns.result1
}

func (fake *FakePivnetClient) AcceptEULACallCount() int {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	return len(fake.acceptEULAArgsForCall)
}

func (fake *FakePivnetClient) AcceptEULAArgsForCall(i int) (string, int) {
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	return fake.acceptEULAArgsForCall[i].productSlug, fake.acceptEULAArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) AcceptEULAReturns(result1 error) {
	fake.AcceptEULAStub = nil
	fake.acceptEULAReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) AcceptEULAReturnsOnCall(i int, result1 error) {
	fake.AcceptEULAStub = nil
	if fake.acceptEULAReturnsOnCall == nil {
		fake.acceptEULAReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.acceptEULAReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakePivnetClient) FileGroupsForRelease(productSlug string, releaseID int) ([]pivnet.FileGroup, error) {
	fake.fileGroupsForReleaseMutex.Lock()
	ret, specificReturn := fake.fileGroupsForReleaseReturnsOnCall[len(fake.fileGroupsForReleaseArgsForCall)]
	fake.fileGroupsForReleaseArgsForCall = append(fake.fileGroupsForReleaseArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("FileGroupsForRelease", []interface{}{productSlug, releaseID})
	fake.fileGroupsForReleaseMutex.Unlock()
	if fake.FileGroupsForReleaseStub != nil {
		return fake.FileGroupsForReleaseStub(productSlug, releaseID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.fileGroupsForReleaseReturns.result1, fake.fileGroupsForReleaseReturns.result2
}

func (fake *FakePivnetClient) FileGroupsForReleaseCallCount() int {
	fake.fileGroupsForReleaseMutex.RLock()
	defer fake.fileGroupsForReleaseMutex.RUnlock()
	return len(fake.fileGroupsForReleaseArgsForCall)
}

func (fake *FakePivnetClient) FileGroupsForReleaseArgsForCall(i int) (string, int) {
	fake.fileGroupsForReleaseMutex.RLock()
	defer fake.fileGroupsForReleaseMutex.RUnlock()
	return fake.fileGroupsForReleaseArgsForCall[i].productSlug, fake.fileGroupsForReleaseArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) FileGroupsForReleaseReturns(result1 []pivnet.FileGroup, result2 error) {
	fake.FileGroupsForReleaseStub = nil
	fake.fileGroupsForReleaseReturns = struct {
		result1 []pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) FileGroupsForReleaseReturnsOnCall(i int, result1 []pivnet.FileGroup, result2 error) {
	fake.FileGroupsForReleaseStub = nil
	if fake.fileGroupsForReleaseReturnsOnCall == nil {
		fake.fileGroupsForReleaseReturnsOnCall = make(map[int]struct {
			result1 []pivnet.FileGroup
			result2 error
		})
	}
	fake.fileGroupsForReleaseReturnsOnCall[i] = struct {
		result1 []pivnet.FileGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ImageReferencesForRelease(productSlug string, releaseID int) ([]pivnet.ImageReference, error) {
	fake.imageReferencesForReleaseMutex.Lock()
	ret, specificReturn := fake.imageReferencesForReleaseReturnsOnCall[len(fake.imageReferencesForReleaseArgsForCall)]
	fake.imageReferencesForReleaseArgsForCall = append(fake.imageReferencesForReleaseArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("ImageReferencesForRelease", []interface{}{productSlug, releaseID})
	fake.imageReferencesForReleaseMutex.Unlock()
	if fake.ImageReferencesForReleaseStub != nil {
		return fake.ImageReferencesForReleaseStub(productSlug, releaseID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.imageReferencesForReleaseReturns.result1, fake.imageReferencesForReleaseReturns.result2
}

func (fake *FakePivnetClient) ImageReferencesForReleaseCallCount() int {
	fake.imageReferencesForReleaseMutex.RLock()
	defer fake.imageReferencesForReleaseMutex.RUnlock()
	return len(fake.imageReferencesForReleaseArgsForCall)
}

func (fake *FakePivnetClient) ImageReferencesForReleaseArgsForCall(i int) (string, int) {
	fake.imageReferencesForReleaseMutex.RLock()
	defer fake.imageReferencesForReleaseMutex.RUnlock()
	return fake.imageReferencesForReleaseArgsForCall[i].productSlug, fake.imageReferencesForReleaseArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) ImageReferencesForReleaseReturns(result1 []pivnet.ImageReference, result2 error) {
	fake.ImageReferencesForReleaseStub = nil
	fake.imageReferencesForReleaseReturns = struct {
		result1 []pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ImageReferencesForReleaseReturnsOnCall(i int, result1 []pivnet.ImageReference, result2 error) {
	fake.ImageReferencesForReleaseStub = nil
	if fake.imageReferencesForReleaseReturnsOnCall == nil {
		fake.imageReferencesForReleaseReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ImageReference
			result2 error
		})
	}
	fake.imageReferencesForReleaseReturnsOnCall[i] = struct {
		result1 []pivnet.ImageReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) HelmChartReferencesForRelease(productSlug string, releaseID int) ([]pivnet.HelmChartReference, error) {
	fake.helmChartReferencesForReleaseMutex.Lock()
	ret, specificReturn := fake.helmChartReferencesForReleaseReturnsOnCall[len(fake.helmChartReferencesForReleaseArgsForCall)]
	fake.helmChartReferencesForReleaseArgsForCall = append(fake.helmChartReferencesForReleaseArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("HelmChartReferencesForRelease", []interface{}{productSlug, releaseID})
	fake.helmChartReferencesForReleaseMutex.Unlock()
	if fake.HelmChartReferencesForReleaseStub != nil {
		return fake.HelmChartReferencesForReleaseStub(productSlug, releaseID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.helmChartReferencesForReleaseReturns.result1, fake.helmChartReferencesForReleaseReturns.result2
}

func (fake *FakePivnetClient) HelmChartReferencesForReleaseCallCount() int {
	fake.helmChartReferencesForReleaseMutex.RLock()
	defer fake.helmChartReferencesForReleaseMutex.RUnlock()
	return len(fake.helmChartReferencesForReleaseArgsForCall)
}

func (fake *FakePivnetClient) HelmChartReferencesForReleaseArgsForCall(i int) (string, int) {
	fake.helmChartReferencesForReleaseMutex.RLock()
	defer fake.helmChartReferencesForReleaseMutex.RUnlock()
	return fake.helmChartReferencesForReleaseArgsForCall[i].productSlug, fake.helmChartReferencesForReleaseArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) HelmChartReferencesForReleaseReturns(result1 []pivnet.HelmChartReference, result2 error) {
	fake.HelmChartReferencesForReleaseStub = nil
	fake.helmChartReferencesForReleaseReturns = struct {
		result1 []pivnet.HelmChartReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) HelmChartReferencesForReleaseReturnsOnCall(i int, result1 []pivnet.HelmChartReference, result2 error) {
	fake.HelmChartReferencesForReleaseStub = nil
	if fake.helmChartReferencesForReleaseReturnsOnCall == nil {
		fake.helmChartReferencesForReleaseReturnsOnCall = make(map[int]struct {
			result1 []pivnet.HelmChartReference
			result2 error
		})
	}
	fake.helmChartReferencesForReleaseReturnsOnCall[i] = struct {
		result1 []pivnet.HelmChartReference
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductFilesForRelease(productSlug string, releaseID int) ([]pivnet.ProductFile, error) {
	fake.productFilesForReleaseMutex.Lock()
	ret, specificReturn := fake.productFilesForReleaseReturnsOnCall[len(fake.productFilesForReleaseArgsForCall)]
	fake.productFilesForReleaseArgsForCall = append(fake.productFilesForReleaseArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("ProductFilesForRelease", []interface{}{productSlug, releaseID})
	fake.productFilesForReleaseMutex.Unlock()
	if fake.ProductFilesForReleaseStub != nil {
		return fake.ProductFilesForReleaseStub(productSlug, releaseID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.productFilesForReleaseReturns.result1, fake.productFilesForReleaseReturns.result2
}

func (fake *FakePivnetClient) ProductFilesForReleaseCallCount() int {
	fake.productFilesForReleaseMutex.RLock()
	defer fake.productFilesForReleaseMutex.RUnlock()
	return len(fake.productFilesForReleaseArgsForCall)
}

func (fake *FakePivnetClient) ProductFilesForReleaseArgsForCall(i int) (string, int) {
	fake.productFilesForReleaseMutex.RLock()
	defer fake.productFilesForReleaseMutex.RUnlock()
	return fake.productFilesForReleaseArgsForCall[i].productSlug, fake.productFilesForReleaseArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) ProductFilesForReleaseReturns(result1 []pivnet.ProductFile, result2 error) {
	fake.ProductFilesForReleaseStub = nil
	fake.productFilesForReleaseReturns = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductFilesForReleaseReturnsOnCall(i int, result1 []pivnet.ProductFile, result2 error) {
	fake.ProductFilesForReleaseStub = nil
	if fake.productFilesForReleaseReturnsOnCall == nil {
		fake.productFilesForReleaseReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ProductFile
			result2 error
		})
	}
	fake.productFilesForReleaseReturnsOnCall[i] = struct {
		result1 []pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductFileForRelease(productSlug string, releaseID int, productFileID int) (pivnet.ProductFile, error) {
	fake.productFileForReleaseMutex.Lock()
	ret, specificReturn := fake.productFileForReleaseReturnsOnCall[len(fake.productFileForReleaseArgsForCall)]
	fake.productFileForReleaseArgsForCall = append(fake.productFileForReleaseArgsForCall, struct {
		productSlug   string
		releaseID     int
		productFileID int
	}{productSlug, releaseID, productFileID})
	fake.recordInvocation("ProductFileForRelease", []interface{}{productSlug, releaseID, productFileID})
	fake.productFileForReleaseMutex.Unlock()
	if fake.ProductFileForReleaseStub != nil {
		return fake.ProductFileForReleaseStub(productSlug, releaseID, productFileID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.productFileForReleaseReturns.result1, fake.productFileForReleaseReturns.result2
}

func (fake *FakePivnetClient) ProductFileForReleaseCallCount() int {
	fake.productFileForReleaseMutex.RLock()
	defer fake.productFileForReleaseMutex.RUnlock()
	return len(fake.productFileForReleaseArgsForCall)
}

func (fake *FakePivnetClient) ProductFileForReleaseArgsForCall(i int) (string, int, int) {
	fake.productFileForReleaseMutex.RLock()
	defer fake.productFileForReleaseMutex.RUnlock()
	return fake.productFileForReleaseArgsForCall[i].productSlug, fake.productFileForReleaseArgsForCall[i].releaseID, fake.productFileForReleaseArgsForCall[i].productFileID
}

func (fake *FakePivnetClient) ProductFileForReleaseReturns(result1 pivnet.ProductFile, result2 error) {
	fake.ProductFileForReleaseStub = nil
	fake.productFileForReleaseReturns = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ProductFileForReleaseReturnsOnCall(i int, result1 pivnet.ProductFile, result2 error) {
	fake.ProductFileForReleaseStub = nil
	if fake.productFileForReleaseReturnsOnCall == nil {
		fake.productFileForReleaseReturnsOnCall = make(map[int]struct {
			result1 pivnet.ProductFile
			result2 error
		})
	}
	fake.productFileForReleaseReturnsOnCall[i] = struct {
		result1 pivnet.ProductFile
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseDependencies(productSlug string, releaseID int) ([]pivnet.ReleaseDependency, error) {
	fake.releaseDependenciesMutex.Lock()
	ret, specificReturn := fake.releaseDependenciesReturnsOnCall[len(fake.releaseDependenciesArgsForCall)]
	fake.releaseDependenciesArgsForCall = append(fake.releaseDependenciesArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("ReleaseDependencies", []interface{}{productSlug, releaseID})
	fake.releaseDependenciesMutex.Unlock()
	if fake.ReleaseDependenciesStub != nil {
		return fake.ReleaseDependenciesStub(productSlug, releaseID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.releaseDependenciesReturns.result1, fake.releaseDependenciesReturns.result2
}

func (fake *FakePivnetClient) ReleaseDependenciesCallCount() int {
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	return len(fake.releaseDependenciesArgsForCall)
}

func (fake *FakePivnetClient) ReleaseDependenciesArgsForCall(i int) (string, int) {
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	return fake.releaseDependenciesArgsForCall[i].productSlug, fake.releaseDependenciesArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) ReleaseDependenciesReturns(result1 []pivnet.ReleaseDependency, result2 error) {
	fake.ReleaseDependenciesStub = nil
	fake.releaseDependenciesReturns = struct {
		result1 []pivnet.ReleaseDependency
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseDependenciesReturnsOnCall(i int, result1 []pivnet.ReleaseDependency, result2 error) {
	fake.ReleaseDependenciesStub = nil
	if fake.releaseDependenciesReturnsOnCall == nil {
		fake.releaseDependenciesReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ReleaseDependency
			result2 error
		})
	}
	fake.releaseDependenciesReturnsOnCall[i] = struct {
		result1 []pivnet.ReleaseDependency
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) DependencySpecifiers(productSlug string, releaseID int) ([]pivnet.DependencySpecifier, error) {
	fake.dependencySpecifiersMutex.Lock()
	ret, specificReturn := fake.dependencySpecifiersReturnsOnCall[len(fake.dependencySpecifiersArgsForCall)]
	fake.dependencySpecifiersArgsForCall = append(fake.dependencySpecifiersArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("DependencySpecifiers", []interface{}{productSlug, releaseID})
	fake.dependencySpecifiersMutex.Unlock()
	if fake.DependencySpecifiersStub != nil {
		return fake.DependencySpecifiersStub(productSlug, releaseID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.dependencySpecifiersReturns.result1, fake.dependencySpecifiersReturns.result2
}

func (fake *FakePivnetClient) DependencySpecifiersCallCount() int {
	fake.dependencySpecifiersMutex.RLock()
	defer fake.dependencySpecifiersMutex.RUnlock()
	return len(fake.dependencySpecifiersArgsForCall)
}

func (fake *FakePivnetClient) DependencySpecifiersArgsForCall(i int) (string, int) {
	fake.dependencySpecifiersMutex.RLock()
	defer fake.dependencySpecifiersMutex.RUnlock()
	return fake.dependencySpecifiersArgsForCall[i].productSlug, fake.dependencySpecifiersArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) DependencySpecifiersReturns(result1 []pivnet.DependencySpecifier, result2 error) {
	fake.DependencySpecifiersStub = nil
	fake.dependencySpecifiersReturns = struct {
		result1 []pivnet.DependencySpecifier
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) DependencySpecifiersReturnsOnCall(i int, result1 []pivnet.DependencySpecifier, result2 error) {
	fake.DependencySpecifiersStub = nil
	if fake.dependencySpecifiersReturnsOnCall == nil {
		fake.dependencySpecifiersReturnsOnCall = make(map[int]struct {
			result1 []pivnet.DependencySpecifier
			result2 error
		})
	}
	fake.dependencySpecifiersReturnsOnCall[i] = struct {
		result1 []pivnet.DependencySpecifier
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseUpgradePaths(productSlug string, releaseID int) ([]pivnet.ReleaseUpgradePath, error) {
	fake.releaseUpgradePathsMutex.Lock()
	ret, specificReturn := fake.releaseUpgradePathsReturnsOnCall[len(fake.releaseUpgradePathsArgsForCall)]
	fake.releaseUpgradePathsArgsForCall = append(fake.releaseUpgradePathsArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("ReleaseUpgradePaths", []interface{}{productSlug, releaseID})
	fake.releaseUpgradePathsMutex.Unlock()
	if fake.ReleaseUpgradePathsStub != nil {
		return fake.ReleaseUpgradePathsStub(productSlug, releaseID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.releaseUpgradePathsReturns.result1, fake.releaseUpgradePathsReturns.result2
}

func (fake *FakePivnetClient) ReleaseUpgradePathsCallCount() int {
	fake.releaseUpgradePathsMutex.RLock()
	defer fake.releaseUpgradePathsMutex.RUnlock()
	return len(fake.releaseUpgradePathsArgsForCall)
}

func (fake *FakePivnetClient) ReleaseUpgradePathsArgsForCall(i int) (string, int) {
	fake.releaseUpgradePathsMutex.RLock()
	defer fake.releaseUpgradePathsMutex.RUnlock()
	return fake.releaseUpgradePathsArgsForCall[i].productSlug, fake.releaseUpgradePathsArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) ReleaseUpgradePathsReturns(result1 []pivnet.ReleaseUpgradePath, result2 error) {
	fake.ReleaseUpgradePathsStub = nil
	fake.releaseUpgradePathsReturns = struct {
		result1 []pivnet.ReleaseUpgradePath
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleaseUpgradePathsReturnsOnCall(i int, result1 []pivnet.ReleaseUpgradePath, result2 error) {
	fake.ReleaseUpgradePathsStub = nil
	if fake.releaseUpgradePathsReturnsOnCall == nil {
		fake.releaseUpgradePathsReturnsOnCall = make(map[int]struct {
			result1 []pivnet.ReleaseUpgradePath
			result2 error
		})
	}
	fake.releaseUpgradePathsReturnsOnCall[i] = struct {
		result1 []pivnet.ReleaseUpgradePath
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UpgradePathSpecifiers(productSlug string, releaseID int) ([]pivnet.UpgradePathSpecifier, error) {
	fake.upgradePathSpecifiersMutex.Lock()
	ret, specificReturn := fake.upgradePathSpecifiersReturnsOnCall[len(fake.upgradePathSpecifiersArgsForCall)]
	fake.upgradePathSpecifiersArgsForCall = append(fake.upgradePathSpecifiersArgsForCall, struct {
		productSlug string
		releaseID   int
	}{productSlug, releaseID})
	fake.recordInvocation("UpgradePathSpecifiers", []interface{}{productSlug, releaseID})
	fake.upgradePathSpecifiersMutex.Unlock()
	if fake.UpgradePathSpecifiersStub != nil {
		return fake.UpgradePathSpecifiersStub(productSlug, releaseID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.upgradePathSpecifiersReturns.result1, fake.upgradePathSpecifiersReturns.result2
}

func (fake *FakePivnetClient) UpgradePathSpecifiersCallCount() int {
	fake.upgradePathSpecifiersMutex.RLock()
	defer fake.upgradePathSpecifiersMutex.RUnlock()
	return len(fake.upgradePathSpecifiersArgsForCall)
}

func (fake *FakePivnetClient) UpgradePathSpecifiersArgsForCall(i int) (string, int) {
	fake.upgradePathSpecifiersMutex.RLock()
	defer fake.upgradePathSpecifiersMutex.RUnlock()
	return fake.upgradePathSpecifiersArgsForCall[i].productSlug, fake.upgradePathSpecifiersArgsForCall[i].releaseID
}

func (fake *FakePivnetClient) UpgradePathSpecifiersReturns(result1 []pivnet.UpgradePathSpecifier, result2 error) {
	fake.UpgradePathSpecifiersStub = nil
	fake.upgradePathSpecifiersReturns = struct {
		result1 []pivnet.UpgradePathSpecifier
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) UpgradePathSpecifiersReturnsOnCall(i int, result1 []pivnet.UpgradePathSpecifier, result2 error) {
	fake.UpgradePathSpecifiersStub = nil
	if fake.upgradePathSpecifiersReturnsOnCall == nil {
		fake.upgradePathSpecifiersReturnsOnCall = make(map[int]struct {
			result1 []pivnet.UpgradePathSpecifier
			result2 error
		})
	}
	fake.upgradePathSpecifiersReturnsOnCall[i] = struct {
		result1 []pivnet.UpgradePathSpecifier
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getReleaseMutex.RLock()
	defer fake.getReleaseMutex.RUnlock()
	fake.acceptEULAMutex.RLock()
	defer fake.acceptEULAMutex.RUnlock()
	fake.fileGroupsForReleaseMutex.RLock()
	defer fake.fileGroupsForReleaseMutex.RUnlock()
	fake.imageReferencesForReleaseMutex.RLock()
	defer fake.imageReferencesForReleaseMutex.RUnlock()
	fake.helmChartReferencesForReleaseMutex.RLock()
	defer fake.helmChartReferencesForReleaseMutex.RUnlock()
	fake.productFilesForReleaseMutex.RLock()
	defer fake.productFilesForReleaseMutex.RUnlock()
	fake.productFileForReleaseMutex.RLock()
	defer fake.productFileForReleaseMutex.RUnlock()
	fake.releaseDependenciesMutex.RLock()
	defer fake.releaseDependenciesMutex.RUnlock()
	fake.dependencySpecifiersMutex.RLock()
	defer fake.dependencySpecifiersMutex.RUnlock()
	fake.releaseUpgradePathsMutex.RLock()
	defer fake.releaseUpgradePathsMutex.RUnlock()
	fake.upgradePathSpecifiersMutex.RLock()
	defer fake.upgradePathSpecifiersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePivnetClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
