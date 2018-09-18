// Code generated by counterfeiter. DO NOT EDIT.
package credsfakes

import (
	"sync"

	"github.com/concourse/concourse/atc/creds"
)

type FakeVariablesFactory struct {
	NewVariablesStub        func(string, string) creds.Variables
	newVariablesMutex       sync.RWMutex
	newVariablesArgsForCall []struct {
		arg1 string
		arg2 string
	}
	newVariablesReturns struct {
		result1 creds.Variables
	}
	newVariablesReturnsOnCall map[int]struct {
		result1 creds.Variables
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVariablesFactory) NewVariables(arg1 string, arg2 string) creds.Variables {
	fake.newVariablesMutex.Lock()
	ret, specificReturn := fake.newVariablesReturnsOnCall[len(fake.newVariablesArgsForCall)]
	fake.newVariablesArgsForCall = append(fake.newVariablesArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("NewVariables", []interface{}{arg1, arg2})
	fake.newVariablesMutex.Unlock()
	if fake.NewVariablesStub != nil {
		return fake.NewVariablesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.newVariablesReturns.result1
}

func (fake *FakeVariablesFactory) NewVariablesCallCount() int {
	fake.newVariablesMutex.RLock()
	defer fake.newVariablesMutex.RUnlock()
	return len(fake.newVariablesArgsForCall)
}

func (fake *FakeVariablesFactory) NewVariablesArgsForCall(i int) (string, string) {
	fake.newVariablesMutex.RLock()
	defer fake.newVariablesMutex.RUnlock()
	return fake.newVariablesArgsForCall[i].arg1, fake.newVariablesArgsForCall[i].arg2
}

func (fake *FakeVariablesFactory) NewVariablesReturns(result1 creds.Variables) {
	fake.NewVariablesStub = nil
	fake.newVariablesReturns = struct {
		result1 creds.Variables
	}{result1}
}

func (fake *FakeVariablesFactory) NewVariablesReturnsOnCall(i int, result1 creds.Variables) {
	fake.NewVariablesStub = nil
	if fake.newVariablesReturnsOnCall == nil {
		fake.newVariablesReturnsOnCall = make(map[int]struct {
			result1 creds.Variables
		})
	}
	fake.newVariablesReturnsOnCall[i] = struct {
		result1 creds.Variables
	}{result1}
}

func (fake *FakeVariablesFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newVariablesMutex.RLock()
	defer fake.newVariablesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVariablesFactory) recordInvocation(key string, args []interface{}) {
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

var _ creds.VariablesFactory = new(FakeVariablesFactory)