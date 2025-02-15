// Code generated by counterfeiter. DO NOT EDIT.
package uaafakes

import (
	"sync"

	"github.com/shono09835/bosh-cli/v7/uaa"
)

type FakeToken struct {
	IsValidStub        func() bool
	isValidMutex       sync.RWMutex
	isValidArgsForCall []struct {
	}
	isValidReturns struct {
		result1 bool
	}
	isValidReturnsOnCall map[int]struct {
		result1 bool
	}
	TypeStub        func() string
	typeMutex       sync.RWMutex
	typeArgsForCall []struct {
	}
	typeReturns struct {
		result1 string
	}
	typeReturnsOnCall map[int]struct {
		result1 string
	}
	ValueStub        func() string
	valueMutex       sync.RWMutex
	valueArgsForCall []struct {
	}
	valueReturns struct {
		result1 string
	}
	valueReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeToken) IsValid() bool {
	fake.isValidMutex.Lock()
	ret, specificReturn := fake.isValidReturnsOnCall[len(fake.isValidArgsForCall)]
	fake.isValidArgsForCall = append(fake.isValidArgsForCall, struct {
	}{})
	stub := fake.IsValidStub
	fakeReturns := fake.isValidReturns
	fake.recordInvocation("IsValid", []interface{}{})
	fake.isValidMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeToken) IsValidCallCount() int {
	fake.isValidMutex.RLock()
	defer fake.isValidMutex.RUnlock()
	return len(fake.isValidArgsForCall)
}

func (fake *FakeToken) IsValidCalls(stub func() bool) {
	fake.isValidMutex.Lock()
	defer fake.isValidMutex.Unlock()
	fake.IsValidStub = stub
}

func (fake *FakeToken) IsValidReturns(result1 bool) {
	fake.isValidMutex.Lock()
	defer fake.isValidMutex.Unlock()
	fake.IsValidStub = nil
	fake.isValidReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeToken) IsValidReturnsOnCall(i int, result1 bool) {
	fake.isValidMutex.Lock()
	defer fake.isValidMutex.Unlock()
	fake.IsValidStub = nil
	if fake.isValidReturnsOnCall == nil {
		fake.isValidReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isValidReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeToken) Type() string {
	fake.typeMutex.Lock()
	ret, specificReturn := fake.typeReturnsOnCall[len(fake.typeArgsForCall)]
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct {
	}{})
	stub := fake.TypeStub
	fakeReturns := fake.typeReturns
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeToken) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeToken) TypeCalls(stub func() string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = stub
}

func (fake *FakeToken) TypeReturns(result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeToken) TypeReturnsOnCall(i int, result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	if fake.typeReturnsOnCall == nil {
		fake.typeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.typeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeToken) Value() string {
	fake.valueMutex.Lock()
	ret, specificReturn := fake.valueReturnsOnCall[len(fake.valueArgsForCall)]
	fake.valueArgsForCall = append(fake.valueArgsForCall, struct {
	}{})
	stub := fake.ValueStub
	fakeReturns := fake.valueReturns
	fake.recordInvocation("Value", []interface{}{})
	fake.valueMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeToken) ValueCallCount() int {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	return len(fake.valueArgsForCall)
}

func (fake *FakeToken) ValueCalls(stub func() string) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = stub
}

func (fake *FakeToken) ValueReturns(result1 string) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	fake.valueReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeToken) ValueReturnsOnCall(i int, result1 string) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	if fake.valueReturnsOnCall == nil {
		fake.valueReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.valueReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeToken) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.isValidMutex.RLock()
	defer fake.isValidMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeToken) recordInvocation(key string, args []interface{}) {
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

var _ uaa.Token = new(FakeToken)
