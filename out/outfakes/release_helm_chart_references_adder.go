// Code generated by counterfeiter. DO NOT EDIT.
package outfakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v4"
)

type ReleaseHelmChartReferencesAdder struct {
	AddReleaseHelmChartReferencesStub        func(release pivnet.Release) error
	addReleaseHelmChartReferencesMutex       sync.RWMutex
	addReleaseHelmChartReferencesArgsForCall []struct {
		release pivnet.Release
	}
	addReleaseHelmChartReferencesReturns struct {
		result1 error
	}
	addReleaseHelmChartReferencesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ReleaseHelmChartReferencesAdder) AddReleaseHelmChartReferences(release pivnet.Release) error {
	fake.addReleaseHelmChartReferencesMutex.Lock()
	ret, specificReturn := fake.addReleaseHelmChartReferencesReturnsOnCall[len(fake.addReleaseHelmChartReferencesArgsForCall)]
	fake.addReleaseHelmChartReferencesArgsForCall = append(fake.addReleaseHelmChartReferencesArgsForCall, struct {
		release pivnet.Release
	}{release})
	fake.recordInvocation("AddReleaseHelmChartReferences", []interface{}{release})
	fake.addReleaseHelmChartReferencesMutex.Unlock()
	if fake.AddReleaseHelmChartReferencesStub != nil {
		return fake.AddReleaseHelmChartReferencesStub(release)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addReleaseHelmChartReferencesReturns.result1
}

func (fake *ReleaseHelmChartReferencesAdder) AddReleaseHelmChartReferencesCallCount() int {
	fake.addReleaseHelmChartReferencesMutex.RLock()
	defer fake.addReleaseHelmChartReferencesMutex.RUnlock()
	return len(fake.addReleaseHelmChartReferencesArgsForCall)
}

func (fake *ReleaseHelmChartReferencesAdder) AddReleaseHelmChartReferencesArgsForCall(i int) pivnet.Release {
	fake.addReleaseHelmChartReferencesMutex.RLock()
	defer fake.addReleaseHelmChartReferencesMutex.RUnlock()
	return fake.addReleaseHelmChartReferencesArgsForCall[i].release
}

func (fake *ReleaseHelmChartReferencesAdder) AddReleaseHelmChartReferencesReturns(result1 error) {
	fake.AddReleaseHelmChartReferencesStub = nil
	fake.addReleaseHelmChartReferencesReturns = struct {
		result1 error
	}{result1}
}

func (fake *ReleaseHelmChartReferencesAdder) AddReleaseHelmChartReferencesReturnsOnCall(i int, result1 error) {
	fake.AddReleaseHelmChartReferencesStub = nil
	if fake.addReleaseHelmChartReferencesReturnsOnCall == nil {
		fake.addReleaseHelmChartReferencesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addReleaseHelmChartReferencesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ReleaseHelmChartReferencesAdder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addReleaseHelmChartReferencesMutex.RLock()
	defer fake.addReleaseHelmChartReferencesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ReleaseHelmChartReferencesAdder) recordInvocation(key string, args []interface{}) {
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
