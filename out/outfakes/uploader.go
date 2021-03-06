// Code generated by counterfeiter. DO NOT EDIT.
package outfakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v4"
)

type Uploader struct {
	UploadStub        func(release pivnet.Release, exactGlobs []string) error
	uploadMutex       sync.RWMutex
	uploadArgsForCall []struct {
		release    pivnet.Release
		exactGlobs []string
	}
	uploadReturns struct {
		result1 error
	}
	uploadReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Uploader) Upload(release pivnet.Release, exactGlobs []string) error {
	var exactGlobsCopy []string
	if exactGlobs != nil {
		exactGlobsCopy = make([]string, len(exactGlobs))
		copy(exactGlobsCopy, exactGlobs)
	}
	fake.uploadMutex.Lock()
	ret, specificReturn := fake.uploadReturnsOnCall[len(fake.uploadArgsForCall)]
	fake.uploadArgsForCall = append(fake.uploadArgsForCall, struct {
		release    pivnet.Release
		exactGlobs []string
	}{release, exactGlobsCopy})
	fake.recordInvocation("Upload", []interface{}{release, exactGlobsCopy})
	fake.uploadMutex.Unlock()
	if fake.UploadStub != nil {
		return fake.UploadStub(release, exactGlobs)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.uploadReturns.result1
}

func (fake *Uploader) UploadCallCount() int {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return len(fake.uploadArgsForCall)
}

func (fake *Uploader) UploadArgsForCall(i int) (pivnet.Release, []string) {
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	return fake.uploadArgsForCall[i].release, fake.uploadArgsForCall[i].exactGlobs
}

func (fake *Uploader) UploadReturns(result1 error) {
	fake.UploadStub = nil
	fake.uploadReturns = struct {
		result1 error
	}{result1}
}

func (fake *Uploader) UploadReturnsOnCall(i int, result1 error) {
	fake.UploadStub = nil
	if fake.uploadReturnsOnCall == nil {
		fake.uploadReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uploadReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Uploader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.uploadMutex.RLock()
	defer fake.uploadMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Uploader) recordInvocation(key string, args []interface{}) {
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
