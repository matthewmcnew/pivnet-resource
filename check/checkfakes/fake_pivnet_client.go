// This file was generated by counterfeiter
package checkfakes

import (
	"sync"

	go_pivnet "github.com/pivotal-cf/go-pivnet/v2"
)

type FakePivnetClient struct {
	ReleaseTypesStub        func() ([]go_pivnet.ReleaseType, error)
	releaseTypesMutex       sync.RWMutex
	releaseTypesArgsForCall []struct{}
	releaseTypesReturns     struct {
		result1 []go_pivnet.ReleaseType
		result2 error
	}
	ReleasesForProductSlugStub        func(string) ([]go_pivnet.Release, error)
	releasesForProductSlugMutex       sync.RWMutex
	releasesForProductSlugArgsForCall []struct {
		arg1 string
	}
	releasesForProductSlugReturns struct {
		result1 []go_pivnet.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePivnetClient) ReleaseTypes() ([]go_pivnet.ReleaseType, error) {
	fake.releaseTypesMutex.Lock()
	fake.releaseTypesArgsForCall = append(fake.releaseTypesArgsForCall, struct{}{})
	fake.recordInvocation("ReleaseTypes", []interface{}{})
	fake.releaseTypesMutex.Unlock()
	if fake.ReleaseTypesStub != nil {
		return fake.ReleaseTypesStub()
	} else {
		return fake.releaseTypesReturns.result1, fake.releaseTypesReturns.result2
	}
}

func (fake *FakePivnetClient) ReleaseTypesCallCount() int {
	fake.releaseTypesMutex.RLock()
	defer fake.releaseTypesMutex.RUnlock()
	return len(fake.releaseTypesArgsForCall)
}

func (fake *FakePivnetClient) ReleaseTypesReturns(result1 []go_pivnet.ReleaseType, result2 error) {
	fake.ReleaseTypesStub = nil
	fake.releaseTypesReturns = struct {
		result1 []go_pivnet.ReleaseType
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) ReleasesForProductSlug(arg1 string) ([]go_pivnet.Release, error) {
	fake.releasesForProductSlugMutex.Lock()
	fake.releasesForProductSlugArgsForCall = append(fake.releasesForProductSlugArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ReleasesForProductSlug", []interface{}{arg1})
	fake.releasesForProductSlugMutex.Unlock()
	if fake.ReleasesForProductSlugStub != nil {
		return fake.ReleasesForProductSlugStub(arg1)
	} else {
		return fake.releasesForProductSlugReturns.result1, fake.releasesForProductSlugReturns.result2
	}
}

func (fake *FakePivnetClient) ReleasesForProductSlugCallCount() int {
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	return len(fake.releasesForProductSlugArgsForCall)
}

func (fake *FakePivnetClient) ReleasesForProductSlugArgsForCall(i int) string {
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	return fake.releasesForProductSlugArgsForCall[i].arg1
}

func (fake *FakePivnetClient) ReleasesForProductSlugReturns(result1 []go_pivnet.Release, result2 error) {
	fake.ReleasesForProductSlugStub = nil
	fake.releasesForProductSlugReturns = struct {
		result1 []go_pivnet.Release
		result2 error
	}{result1, result2}
}

func (fake *FakePivnetClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.releaseTypesMutex.RLock()
	defer fake.releaseTypesMutex.RUnlock()
	fake.releasesForProductSlugMutex.RLock()
	defer fake.releasesForProductSlugMutex.RUnlock()
	return fake.invocations
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
