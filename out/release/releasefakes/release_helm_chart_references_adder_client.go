// Code generated by counterfeiter. DO NOT EDIT.
package releasefakes

import (
	"sync"

	pivnet "github.com/pivotal-cf/go-pivnet/v4"
)

type ReleaseHelmChartReferencesAdderClient struct {
	AddHelmChartReferenceStub        func(string, int, int) error
	addHelmChartReferenceMutex       sync.RWMutex
	addHelmChartReferenceArgsForCall []struct {
		arg1 string
		arg2 int
		arg3 int
	}
	addHelmChartReferenceReturns struct {
		result1 error
	}
	addHelmChartReferenceReturnsOnCall map[int]struct {
		result1 error
	}
	CreateHelmChartReferenceStub        func(pivnet.CreateHelmChartReferenceConfig) (pivnet.HelmChartReference, error)
	createHelmChartReferenceMutex       sync.RWMutex
	createHelmChartReferenceArgsForCall []struct {
		arg1 pivnet.CreateHelmChartReferenceConfig
	}
	createHelmChartReferenceReturns struct {
		result1 pivnet.HelmChartReference
		result2 error
	}
	createHelmChartReferenceReturnsOnCall map[int]struct {
		result1 pivnet.HelmChartReference
		result2 error
	}
	DeleteHelmChartReferenceStub        func(string, int) (pivnet.HelmChartReference, error)
	deleteHelmChartReferenceMutex       sync.RWMutex
	deleteHelmChartReferenceArgsForCall []struct {
		arg1 string
		arg2 int
	}
	deleteHelmChartReferenceReturns struct {
		result1 pivnet.HelmChartReference
		result2 error
	}
	deleteHelmChartReferenceReturnsOnCall map[int]struct {
		result1 pivnet.HelmChartReference
		result2 error
	}
	GetHelmChartReferenceStub        func(string, int) (pivnet.HelmChartReference, error)
	getHelmChartReferenceMutex       sync.RWMutex
	getHelmChartReferenceArgsForCall []struct {
		arg1 string
		arg2 int
	}
	getHelmChartReferenceReturns struct {
		result1 pivnet.HelmChartReference
		result2 error
	}
	getHelmChartReferenceReturnsOnCall map[int]struct {
		result1 pivnet.HelmChartReference
		result2 error
	}
	HelmChartReferencesStub        func(string) ([]pivnet.HelmChartReference, error)
	helmChartReferencesMutex       sync.RWMutex
	helmChartReferencesArgsForCall []struct {
		arg1 string
	}
	helmChartReferencesReturns struct {
		result1 []pivnet.HelmChartReference
		result2 error
	}
	helmChartReferencesReturnsOnCall map[int]struct {
		result1 []pivnet.HelmChartReference
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ReleaseHelmChartReferencesAdderClient) AddHelmChartReference(arg1 string, arg2 int, arg3 int) error {
	fake.addHelmChartReferenceMutex.Lock()
	ret, specificReturn := fake.addHelmChartReferenceReturnsOnCall[len(fake.addHelmChartReferenceArgsForCall)]
	fake.addHelmChartReferenceArgsForCall = append(fake.addHelmChartReferenceArgsForCall, struct {
		arg1 string
		arg2 int
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("AddHelmChartReference", []interface{}{arg1, arg2, arg3})
	fake.addHelmChartReferenceMutex.Unlock()
	if fake.AddHelmChartReferenceStub != nil {
		return fake.AddHelmChartReferenceStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.addHelmChartReferenceReturns
	return fakeReturns.result1
}

func (fake *ReleaseHelmChartReferencesAdderClient) AddHelmChartReferenceCallCount() int {
	fake.addHelmChartReferenceMutex.RLock()
	defer fake.addHelmChartReferenceMutex.RUnlock()
	return len(fake.addHelmChartReferenceArgsForCall)
}

func (fake *ReleaseHelmChartReferencesAdderClient) AddHelmChartReferenceCalls(stub func(string, int, int) error) {
	fake.addHelmChartReferenceMutex.Lock()
	defer fake.addHelmChartReferenceMutex.Unlock()
	fake.AddHelmChartReferenceStub = stub
}

func (fake *ReleaseHelmChartReferencesAdderClient) AddHelmChartReferenceArgsForCall(i int) (string, int, int) {
	fake.addHelmChartReferenceMutex.RLock()
	defer fake.addHelmChartReferenceMutex.RUnlock()
	argsForCall := fake.addHelmChartReferenceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *ReleaseHelmChartReferencesAdderClient) AddHelmChartReferenceReturns(result1 error) {
	fake.addHelmChartReferenceMutex.Lock()
	defer fake.addHelmChartReferenceMutex.Unlock()
	fake.AddHelmChartReferenceStub = nil
	fake.addHelmChartReferenceReturns = struct {
		result1 error
	}{result1}
}

func (fake *ReleaseHelmChartReferencesAdderClient) AddHelmChartReferenceReturnsOnCall(i int, result1 error) {
	fake.addHelmChartReferenceMutex.Lock()
	defer fake.addHelmChartReferenceMutex.Unlock()
	fake.AddHelmChartReferenceStub = nil
	if fake.addHelmChartReferenceReturnsOnCall == nil {
		fake.addHelmChartReferenceReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addHelmChartReferenceReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *ReleaseHelmChartReferencesAdderClient) CreateHelmChartReference(arg1 pivnet.CreateHelmChartReferenceConfig) (pivnet.HelmChartReference, error) {
	fake.createHelmChartReferenceMutex.Lock()
	ret, specificReturn := fake.createHelmChartReferenceReturnsOnCall[len(fake.createHelmChartReferenceArgsForCall)]
	fake.createHelmChartReferenceArgsForCall = append(fake.createHelmChartReferenceArgsForCall, struct {
		arg1 pivnet.CreateHelmChartReferenceConfig
	}{arg1})
	fake.recordInvocation("CreateHelmChartReference", []interface{}{arg1})
	fake.createHelmChartReferenceMutex.Unlock()
	if fake.CreateHelmChartReferenceStub != nil {
		return fake.CreateHelmChartReferenceStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createHelmChartReferenceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ReleaseHelmChartReferencesAdderClient) CreateHelmChartReferenceCallCount() int {
	fake.createHelmChartReferenceMutex.RLock()
	defer fake.createHelmChartReferenceMutex.RUnlock()
	return len(fake.createHelmChartReferenceArgsForCall)
}

func (fake *ReleaseHelmChartReferencesAdderClient) CreateHelmChartReferenceCalls(stub func(pivnet.CreateHelmChartReferenceConfig) (pivnet.HelmChartReference, error)) {
	fake.createHelmChartReferenceMutex.Lock()
	defer fake.createHelmChartReferenceMutex.Unlock()
	fake.CreateHelmChartReferenceStub = stub
}

func (fake *ReleaseHelmChartReferencesAdderClient) CreateHelmChartReferenceArgsForCall(i int) pivnet.CreateHelmChartReferenceConfig {
	fake.createHelmChartReferenceMutex.RLock()
	defer fake.createHelmChartReferenceMutex.RUnlock()
	argsForCall := fake.createHelmChartReferenceArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ReleaseHelmChartReferencesAdderClient) CreateHelmChartReferenceReturns(result1 pivnet.HelmChartReference, result2 error) {
	fake.createHelmChartReferenceMutex.Lock()
	defer fake.createHelmChartReferenceMutex.Unlock()
	fake.CreateHelmChartReferenceStub = nil
	fake.createHelmChartReferenceReturns = struct {
		result1 pivnet.HelmChartReference
		result2 error
	}{result1, result2}
}

func (fake *ReleaseHelmChartReferencesAdderClient) CreateHelmChartReferenceReturnsOnCall(i int, result1 pivnet.HelmChartReference, result2 error) {
	fake.createHelmChartReferenceMutex.Lock()
	defer fake.createHelmChartReferenceMutex.Unlock()
	fake.CreateHelmChartReferenceStub = nil
	if fake.createHelmChartReferenceReturnsOnCall == nil {
		fake.createHelmChartReferenceReturnsOnCall = make(map[int]struct {
			result1 pivnet.HelmChartReference
			result2 error
		})
	}
	fake.createHelmChartReferenceReturnsOnCall[i] = struct {
		result1 pivnet.HelmChartReference
		result2 error
	}{result1, result2}
}

func (fake *ReleaseHelmChartReferencesAdderClient) DeleteHelmChartReference(arg1 string, arg2 int) (pivnet.HelmChartReference, error) {
	fake.deleteHelmChartReferenceMutex.Lock()
	ret, specificReturn := fake.deleteHelmChartReferenceReturnsOnCall[len(fake.deleteHelmChartReferenceArgsForCall)]
	fake.deleteHelmChartReferenceArgsForCall = append(fake.deleteHelmChartReferenceArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("DeleteHelmChartReference", []interface{}{arg1, arg2})
	fake.deleteHelmChartReferenceMutex.Unlock()
	if fake.DeleteHelmChartReferenceStub != nil {
		return fake.DeleteHelmChartReferenceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deleteHelmChartReferenceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ReleaseHelmChartReferencesAdderClient) DeleteHelmChartReferenceCallCount() int {
	fake.deleteHelmChartReferenceMutex.RLock()
	defer fake.deleteHelmChartReferenceMutex.RUnlock()
	return len(fake.deleteHelmChartReferenceArgsForCall)
}

func (fake *ReleaseHelmChartReferencesAdderClient) DeleteHelmChartReferenceCalls(stub func(string, int) (pivnet.HelmChartReference, error)) {
	fake.deleteHelmChartReferenceMutex.Lock()
	defer fake.deleteHelmChartReferenceMutex.Unlock()
	fake.DeleteHelmChartReferenceStub = stub
}

func (fake *ReleaseHelmChartReferencesAdderClient) DeleteHelmChartReferenceArgsForCall(i int) (string, int) {
	fake.deleteHelmChartReferenceMutex.RLock()
	defer fake.deleteHelmChartReferenceMutex.RUnlock()
	argsForCall := fake.deleteHelmChartReferenceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ReleaseHelmChartReferencesAdderClient) DeleteHelmChartReferenceReturns(result1 pivnet.HelmChartReference, result2 error) {
	fake.deleteHelmChartReferenceMutex.Lock()
	defer fake.deleteHelmChartReferenceMutex.Unlock()
	fake.DeleteHelmChartReferenceStub = nil
	fake.deleteHelmChartReferenceReturns = struct {
		result1 pivnet.HelmChartReference
		result2 error
	}{result1, result2}
}

func (fake *ReleaseHelmChartReferencesAdderClient) DeleteHelmChartReferenceReturnsOnCall(i int, result1 pivnet.HelmChartReference, result2 error) {
	fake.deleteHelmChartReferenceMutex.Lock()
	defer fake.deleteHelmChartReferenceMutex.Unlock()
	fake.DeleteHelmChartReferenceStub = nil
	if fake.deleteHelmChartReferenceReturnsOnCall == nil {
		fake.deleteHelmChartReferenceReturnsOnCall = make(map[int]struct {
			result1 pivnet.HelmChartReference
			result2 error
		})
	}
	fake.deleteHelmChartReferenceReturnsOnCall[i] = struct {
		result1 pivnet.HelmChartReference
		result2 error
	}{result1, result2}
}

func (fake *ReleaseHelmChartReferencesAdderClient) GetHelmChartReference(arg1 string, arg2 int) (pivnet.HelmChartReference, error) {
	fake.getHelmChartReferenceMutex.Lock()
	ret, specificReturn := fake.getHelmChartReferenceReturnsOnCall[len(fake.getHelmChartReferenceArgsForCall)]
	fake.getHelmChartReferenceArgsForCall = append(fake.getHelmChartReferenceArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("GetHelmChartReference", []interface{}{arg1, arg2})
	fake.getHelmChartReferenceMutex.Unlock()
	if fake.GetHelmChartReferenceStub != nil {
		return fake.GetHelmChartReferenceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getHelmChartReferenceReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ReleaseHelmChartReferencesAdderClient) GetHelmChartReferenceCallCount() int {
	fake.getHelmChartReferenceMutex.RLock()
	defer fake.getHelmChartReferenceMutex.RUnlock()
	return len(fake.getHelmChartReferenceArgsForCall)
}

func (fake *ReleaseHelmChartReferencesAdderClient) GetHelmChartReferenceCalls(stub func(string, int) (pivnet.HelmChartReference, error)) {
	fake.getHelmChartReferenceMutex.Lock()
	defer fake.getHelmChartReferenceMutex.Unlock()
	fake.GetHelmChartReferenceStub = stub
}

func (fake *ReleaseHelmChartReferencesAdderClient) GetHelmChartReferenceArgsForCall(i int) (string, int) {
	fake.getHelmChartReferenceMutex.RLock()
	defer fake.getHelmChartReferenceMutex.RUnlock()
	argsForCall := fake.getHelmChartReferenceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ReleaseHelmChartReferencesAdderClient) GetHelmChartReferenceReturns(result1 pivnet.HelmChartReference, result2 error) {
	fake.getHelmChartReferenceMutex.Lock()
	defer fake.getHelmChartReferenceMutex.Unlock()
	fake.GetHelmChartReferenceStub = nil
	fake.getHelmChartReferenceReturns = struct {
		result1 pivnet.HelmChartReference
		result2 error
	}{result1, result2}
}

func (fake *ReleaseHelmChartReferencesAdderClient) GetHelmChartReferenceReturnsOnCall(i int, result1 pivnet.HelmChartReference, result2 error) {
	fake.getHelmChartReferenceMutex.Lock()
	defer fake.getHelmChartReferenceMutex.Unlock()
	fake.GetHelmChartReferenceStub = nil
	if fake.getHelmChartReferenceReturnsOnCall == nil {
		fake.getHelmChartReferenceReturnsOnCall = make(map[int]struct {
			result1 pivnet.HelmChartReference
			result2 error
		})
	}
	fake.getHelmChartReferenceReturnsOnCall[i] = struct {
		result1 pivnet.HelmChartReference
		result2 error
	}{result1, result2}
}

func (fake *ReleaseHelmChartReferencesAdderClient) HelmChartReferences(arg1 string) ([]pivnet.HelmChartReference, error) {
	fake.helmChartReferencesMutex.Lock()
	ret, specificReturn := fake.helmChartReferencesReturnsOnCall[len(fake.helmChartReferencesArgsForCall)]
	fake.helmChartReferencesArgsForCall = append(fake.helmChartReferencesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("HelmChartReferences", []interface{}{arg1})
	fake.helmChartReferencesMutex.Unlock()
	if fake.HelmChartReferencesStub != nil {
		return fake.HelmChartReferencesStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.helmChartReferencesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ReleaseHelmChartReferencesAdderClient) HelmChartReferencesCallCount() int {
	fake.helmChartReferencesMutex.RLock()
	defer fake.helmChartReferencesMutex.RUnlock()
	return len(fake.helmChartReferencesArgsForCall)
}

func (fake *ReleaseHelmChartReferencesAdderClient) HelmChartReferencesCalls(stub func(string) ([]pivnet.HelmChartReference, error)) {
	fake.helmChartReferencesMutex.Lock()
	defer fake.helmChartReferencesMutex.Unlock()
	fake.HelmChartReferencesStub = stub
}

func (fake *ReleaseHelmChartReferencesAdderClient) HelmChartReferencesArgsForCall(i int) string {
	fake.helmChartReferencesMutex.RLock()
	defer fake.helmChartReferencesMutex.RUnlock()
	argsForCall := fake.helmChartReferencesArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ReleaseHelmChartReferencesAdderClient) HelmChartReferencesReturns(result1 []pivnet.HelmChartReference, result2 error) {
	fake.helmChartReferencesMutex.Lock()
	defer fake.helmChartReferencesMutex.Unlock()
	fake.HelmChartReferencesStub = nil
	fake.helmChartReferencesReturns = struct {
		result1 []pivnet.HelmChartReference
		result2 error
	}{result1, result2}
}

func (fake *ReleaseHelmChartReferencesAdderClient) HelmChartReferencesReturnsOnCall(i int, result1 []pivnet.HelmChartReference, result2 error) {
	fake.helmChartReferencesMutex.Lock()
	defer fake.helmChartReferencesMutex.Unlock()
	fake.HelmChartReferencesStub = nil
	if fake.helmChartReferencesReturnsOnCall == nil {
		fake.helmChartReferencesReturnsOnCall = make(map[int]struct {
			result1 []pivnet.HelmChartReference
			result2 error
		})
	}
	fake.helmChartReferencesReturnsOnCall[i] = struct {
		result1 []pivnet.HelmChartReference
		result2 error
	}{result1, result2}
}

func (fake *ReleaseHelmChartReferencesAdderClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addHelmChartReferenceMutex.RLock()
	defer fake.addHelmChartReferenceMutex.RUnlock()
	fake.createHelmChartReferenceMutex.RLock()
	defer fake.createHelmChartReferenceMutex.RUnlock()
	fake.deleteHelmChartReferenceMutex.RLock()
	defer fake.deleteHelmChartReferenceMutex.RUnlock()
	fake.getHelmChartReferenceMutex.RLock()
	defer fake.getHelmChartReferenceMutex.RUnlock()
	fake.helmChartReferencesMutex.RLock()
	defer fake.helmChartReferencesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ReleaseHelmChartReferencesAdderClient) recordInvocation(key string, args []interface{}) {
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
