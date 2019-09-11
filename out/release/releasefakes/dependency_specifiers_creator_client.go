// This file was generated by counterfeiter
package releasefakes

import (
	"sync"

	go_pivnet "github.com/pivotal-cf/go-pivnet/v2"
)

type DependencySpecifiersCreatorClient struct {
	CreateDependencySpecifierStub        func(productSlug string, releaseID int, dependentProductSlug string, specifier string) (go_pivnet.DependencySpecifier, error)
	createDependencySpecifierMutex       sync.RWMutex
	createDependencySpecifierArgsForCall []struct {
		productSlug          string
		releaseID            int
		dependentProductSlug string
		specifier            string
	}
	createDependencySpecifierReturns struct {
		result1 go_pivnet.DependencySpecifier
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DependencySpecifiersCreatorClient) CreateDependencySpecifier(productSlug string, releaseID int, dependentProductSlug string, specifier string) (go_pivnet.DependencySpecifier, error) {
	fake.createDependencySpecifierMutex.Lock()
	fake.createDependencySpecifierArgsForCall = append(fake.createDependencySpecifierArgsForCall, struct {
		productSlug          string
		releaseID            int
		dependentProductSlug string
		specifier            string
	}{productSlug, releaseID, dependentProductSlug, specifier})
	fake.recordInvocation("CreateDependencySpecifier", []interface{}{productSlug, releaseID, dependentProductSlug, specifier})
	fake.createDependencySpecifierMutex.Unlock()
	if fake.CreateDependencySpecifierStub != nil {
		return fake.CreateDependencySpecifierStub(productSlug, releaseID, dependentProductSlug, specifier)
	} else {
		return fake.createDependencySpecifierReturns.result1, fake.createDependencySpecifierReturns.result2
	}
}

func (fake *DependencySpecifiersCreatorClient) CreateDependencySpecifierCallCount() int {
	fake.createDependencySpecifierMutex.RLock()
	defer fake.createDependencySpecifierMutex.RUnlock()
	return len(fake.createDependencySpecifierArgsForCall)
}

func (fake *DependencySpecifiersCreatorClient) CreateDependencySpecifierArgsForCall(i int) (string, int, string, string) {
	fake.createDependencySpecifierMutex.RLock()
	defer fake.createDependencySpecifierMutex.RUnlock()
	return fake.createDependencySpecifierArgsForCall[i].productSlug, fake.createDependencySpecifierArgsForCall[i].releaseID, fake.createDependencySpecifierArgsForCall[i].dependentProductSlug, fake.createDependencySpecifierArgsForCall[i].specifier
}

func (fake *DependencySpecifiersCreatorClient) CreateDependencySpecifierReturns(result1 go_pivnet.DependencySpecifier, result2 error) {
	fake.CreateDependencySpecifierStub = nil
	fake.createDependencySpecifierReturns = struct {
		result1 go_pivnet.DependencySpecifier
		result2 error
	}{result1, result2}
}

func (fake *DependencySpecifiersCreatorClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createDependencySpecifierMutex.RLock()
	defer fake.createDependencySpecifierMutex.RUnlock()
	return fake.invocations
}

func (fake *DependencySpecifiersCreatorClient) recordInvocation(key string, args []interface{}) {
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
