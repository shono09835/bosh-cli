// Code generated by counterfeiter. DO NOT EDIT.
package cmdfakes

import (
	"sync"

	"github.com/shono09835/bosh-cli/v7/cmd"
	"github.com/shono09835/bosh-cli/v7/cmd/opts"
	"github.com/shono09835/bosh-cli/v7/release"
)

type FakeReleaseCreatingCmd struct {
	RunStub        func(opts.CreateReleaseOpts) (release.Release, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 opts.CreateReleaseOpts
	}
	runReturns struct {
		result1 release.Release
		result2 error
	}
	runReturnsOnCall map[int]struct {
		result1 release.Release
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReleaseCreatingCmd) Run(arg1 opts.CreateReleaseOpts) (release.Release, error) {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 opts.CreateReleaseOpts
	}{arg1})
	stub := fake.RunStub
	fakeReturns := fake.runReturns
	fake.recordInvocation("Run", []interface{}{arg1})
	fake.runMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReleaseCreatingCmd) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeReleaseCreatingCmd) RunCalls(stub func(opts.CreateReleaseOpts) (release.Release, error)) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakeReleaseCreatingCmd) RunArgsForCall(i int) opts.CreateReleaseOpts {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	argsForCall := fake.runArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReleaseCreatingCmd) RunReturns(result1 release.Release, result2 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 release.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseCreatingCmd) RunReturnsOnCall(i int, result1 release.Release, result2 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 release.Release
			result2 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 release.Release
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseCreatingCmd) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReleaseCreatingCmd) recordInvocation(key string, args []interface{}) {
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

var _ cmd.ReleaseCreatingCmd = new(FakeReleaseCreatingCmd)
