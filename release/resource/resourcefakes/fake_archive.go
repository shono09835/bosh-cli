// Code generated by counterfeiter. DO NOT EDIT.
package resourcefakes

import (
	"sync"

	"github.com/shono09835/bosh-cli/v7/release/resource"
)

type FakeArchive struct {
	BuildStub        func(string) (string, string, error)
	buildMutex       sync.RWMutex
	buildArgsForCall []struct {
		arg1 string
	}
	buildReturns struct {
		result1 string
		result2 string
		result3 error
	}
	buildReturnsOnCall map[int]struct {
		result1 string
		result2 string
		result3 error
	}
	CleanUpStub        func(string)
	cleanUpMutex       sync.RWMutex
	cleanUpArgsForCall []struct {
		arg1 string
	}
	FingerprintStub        func() (string, error)
	fingerprintMutex       sync.RWMutex
	fingerprintArgsForCall []struct {
	}
	fingerprintReturns struct {
		result1 string
		result2 error
	}
	fingerprintReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeArchive) Build(arg1 string) (string, string, error) {
	fake.buildMutex.Lock()
	ret, specificReturn := fake.buildReturnsOnCall[len(fake.buildArgsForCall)]
	fake.buildArgsForCall = append(fake.buildArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.BuildStub
	fakeReturns := fake.buildReturns
	fake.recordInvocation("Build", []interface{}{arg1})
	fake.buildMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeArchive) BuildCallCount() int {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	return len(fake.buildArgsForCall)
}

func (fake *FakeArchive) BuildCalls(stub func(string) (string, string, error)) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = stub
}

func (fake *FakeArchive) BuildArgsForCall(i int) string {
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	argsForCall := fake.buildArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeArchive) BuildReturns(result1 string, result2 string, result3 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	fake.buildReturns = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeArchive) BuildReturnsOnCall(i int, result1 string, result2 string, result3 error) {
	fake.buildMutex.Lock()
	defer fake.buildMutex.Unlock()
	fake.BuildStub = nil
	if fake.buildReturnsOnCall == nil {
		fake.buildReturnsOnCall = make(map[int]struct {
			result1 string
			result2 string
			result3 error
		})
	}
	fake.buildReturnsOnCall[i] = struct {
		result1 string
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeArchive) CleanUp(arg1 string) {
	fake.cleanUpMutex.Lock()
	fake.cleanUpArgsForCall = append(fake.cleanUpArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CleanUpStub
	fake.recordInvocation("CleanUp", []interface{}{arg1})
	fake.cleanUpMutex.Unlock()
	if stub != nil {
		fake.CleanUpStub(arg1)
	}
}

func (fake *FakeArchive) CleanUpCallCount() int {
	fake.cleanUpMutex.RLock()
	defer fake.cleanUpMutex.RUnlock()
	return len(fake.cleanUpArgsForCall)
}

func (fake *FakeArchive) CleanUpCalls(stub func(string)) {
	fake.cleanUpMutex.Lock()
	defer fake.cleanUpMutex.Unlock()
	fake.CleanUpStub = stub
}

func (fake *FakeArchive) CleanUpArgsForCall(i int) string {
	fake.cleanUpMutex.RLock()
	defer fake.cleanUpMutex.RUnlock()
	argsForCall := fake.cleanUpArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeArchive) Fingerprint() (string, error) {
	fake.fingerprintMutex.Lock()
	ret, specificReturn := fake.fingerprintReturnsOnCall[len(fake.fingerprintArgsForCall)]
	fake.fingerprintArgsForCall = append(fake.fingerprintArgsForCall, struct {
	}{})
	stub := fake.FingerprintStub
	fakeReturns := fake.fingerprintReturns
	fake.recordInvocation("Fingerprint", []interface{}{})
	fake.fingerprintMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeArchive) FingerprintCallCount() int {
	fake.fingerprintMutex.RLock()
	defer fake.fingerprintMutex.RUnlock()
	return len(fake.fingerprintArgsForCall)
}

func (fake *FakeArchive) FingerprintCalls(stub func() (string, error)) {
	fake.fingerprintMutex.Lock()
	defer fake.fingerprintMutex.Unlock()
	fake.FingerprintStub = stub
}

func (fake *FakeArchive) FingerprintReturns(result1 string, result2 error) {
	fake.fingerprintMutex.Lock()
	defer fake.fingerprintMutex.Unlock()
	fake.FingerprintStub = nil
	fake.fingerprintReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeArchive) FingerprintReturnsOnCall(i int, result1 string, result2 error) {
	fake.fingerprintMutex.Lock()
	defer fake.fingerprintMutex.Unlock()
	fake.FingerprintStub = nil
	if fake.fingerprintReturnsOnCall == nil {
		fake.fingerprintReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.fingerprintReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeArchive) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.buildMutex.RLock()
	defer fake.buildMutex.RUnlock()
	fake.cleanUpMutex.RLock()
	defer fake.cleanUpMutex.RUnlock()
	fake.fingerprintMutex.RLock()
	defer fake.fingerprintMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeArchive) recordInvocation(key string, args []interface{}) {
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

var _ resource.Archive = new(FakeArchive)
