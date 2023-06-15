// Code generated by counterfeiter. DO NOT EDIT.
package releasedirfakes

import (
	"sync"

	"github.com/shono09835/bosh-cli/v7/release"
	"github.com/shono09835/bosh-cli/v7/release/manifest"
	"github.com/shono09835/bosh-cli/v7/releasedir"
	"github.com/cppforlife/go-semi-semantic/version"
)

type FakeReleaseIndex struct {
	AddStub        func(manifest.Manifest) error
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 manifest.Manifest
	}
	addReturns struct {
		result1 error
	}
	addReturnsOnCall map[int]struct {
		result1 error
	}
	ContainsStub        func(release.Release) (bool, error)
	containsMutex       sync.RWMutex
	containsArgsForCall []struct {
		arg1 release.Release
	}
	containsReturns struct {
		result1 bool
		result2 error
	}
	containsReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	LastVersionStub        func(string) (*version.Version, error)
	lastVersionMutex       sync.RWMutex
	lastVersionArgsForCall []struct {
		arg1 string
	}
	lastVersionReturns struct {
		result1 *version.Version
		result2 error
	}
	lastVersionReturnsOnCall map[int]struct {
		result1 *version.Version
		result2 error
	}
	ManifestPathStub        func(string, string) string
	manifestPathMutex       sync.RWMutex
	manifestPathArgsForCall []struct {
		arg1 string
		arg2 string
	}
	manifestPathReturns struct {
		result1 string
	}
	manifestPathReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReleaseIndex) Add(arg1 manifest.Manifest) error {
	fake.addMutex.Lock()
	ret, specificReturn := fake.addReturnsOnCall[len(fake.addArgsForCall)]
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		arg1 manifest.Manifest
	}{arg1})
	stub := fake.AddStub
	fakeReturns := fake.addReturns
	fake.recordInvocation("Add", []interface{}{arg1})
	fake.addMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseIndex) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *FakeReleaseIndex) AddCalls(stub func(manifest.Manifest) error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = stub
}

func (fake *FakeReleaseIndex) AddArgsForCall(i int) manifest.Manifest {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	argsForCall := fake.addArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReleaseIndex) AddReturns(result1 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseIndex) AddReturnsOnCall(i int, result1 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	if fake.addReturnsOnCall == nil {
		fake.addReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeReleaseIndex) Contains(arg1 release.Release) (bool, error) {
	fake.containsMutex.Lock()
	ret, specificReturn := fake.containsReturnsOnCall[len(fake.containsArgsForCall)]
	fake.containsArgsForCall = append(fake.containsArgsForCall, struct {
		arg1 release.Release
	}{arg1})
	stub := fake.ContainsStub
	fakeReturns := fake.containsReturns
	fake.recordInvocation("Contains", []interface{}{arg1})
	fake.containsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReleaseIndex) ContainsCallCount() int {
	fake.containsMutex.RLock()
	defer fake.containsMutex.RUnlock()
	return len(fake.containsArgsForCall)
}

func (fake *FakeReleaseIndex) ContainsCalls(stub func(release.Release) (bool, error)) {
	fake.containsMutex.Lock()
	defer fake.containsMutex.Unlock()
	fake.ContainsStub = stub
}

func (fake *FakeReleaseIndex) ContainsArgsForCall(i int) release.Release {
	fake.containsMutex.RLock()
	defer fake.containsMutex.RUnlock()
	argsForCall := fake.containsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReleaseIndex) ContainsReturns(result1 bool, result2 error) {
	fake.containsMutex.Lock()
	defer fake.containsMutex.Unlock()
	fake.ContainsStub = nil
	fake.containsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseIndex) ContainsReturnsOnCall(i int, result1 bool, result2 error) {
	fake.containsMutex.Lock()
	defer fake.containsMutex.Unlock()
	fake.ContainsStub = nil
	if fake.containsReturnsOnCall == nil {
		fake.containsReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.containsReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseIndex) LastVersion(arg1 string) (*version.Version, error) {
	fake.lastVersionMutex.Lock()
	ret, specificReturn := fake.lastVersionReturnsOnCall[len(fake.lastVersionArgsForCall)]
	fake.lastVersionArgsForCall = append(fake.lastVersionArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.LastVersionStub
	fakeReturns := fake.lastVersionReturns
	fake.recordInvocation("LastVersion", []interface{}{arg1})
	fake.lastVersionMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeReleaseIndex) LastVersionCallCount() int {
	fake.lastVersionMutex.RLock()
	defer fake.lastVersionMutex.RUnlock()
	return len(fake.lastVersionArgsForCall)
}

func (fake *FakeReleaseIndex) LastVersionCalls(stub func(string) (*version.Version, error)) {
	fake.lastVersionMutex.Lock()
	defer fake.lastVersionMutex.Unlock()
	fake.LastVersionStub = stub
}

func (fake *FakeReleaseIndex) LastVersionArgsForCall(i int) string {
	fake.lastVersionMutex.RLock()
	defer fake.lastVersionMutex.RUnlock()
	argsForCall := fake.lastVersionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReleaseIndex) LastVersionReturns(result1 *version.Version, result2 error) {
	fake.lastVersionMutex.Lock()
	defer fake.lastVersionMutex.Unlock()
	fake.LastVersionStub = nil
	fake.lastVersionReturns = struct {
		result1 *version.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseIndex) LastVersionReturnsOnCall(i int, result1 *version.Version, result2 error) {
	fake.lastVersionMutex.Lock()
	defer fake.lastVersionMutex.Unlock()
	fake.LastVersionStub = nil
	if fake.lastVersionReturnsOnCall == nil {
		fake.lastVersionReturnsOnCall = make(map[int]struct {
			result1 *version.Version
			result2 error
		})
	}
	fake.lastVersionReturnsOnCall[i] = struct {
		result1 *version.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseIndex) ManifestPath(arg1 string, arg2 string) string {
	fake.manifestPathMutex.Lock()
	ret, specificReturn := fake.manifestPathReturnsOnCall[len(fake.manifestPathArgsForCall)]
	fake.manifestPathArgsForCall = append(fake.manifestPathArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.ManifestPathStub
	fakeReturns := fake.manifestPathReturns
	fake.recordInvocation("ManifestPath", []interface{}{arg1, arg2})
	fake.manifestPathMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeReleaseIndex) ManifestPathCallCount() int {
	fake.manifestPathMutex.RLock()
	defer fake.manifestPathMutex.RUnlock()
	return len(fake.manifestPathArgsForCall)
}

func (fake *FakeReleaseIndex) ManifestPathCalls(stub func(string, string) string) {
	fake.manifestPathMutex.Lock()
	defer fake.manifestPathMutex.Unlock()
	fake.ManifestPathStub = stub
}

func (fake *FakeReleaseIndex) ManifestPathArgsForCall(i int) (string, string) {
	fake.manifestPathMutex.RLock()
	defer fake.manifestPathMutex.RUnlock()
	argsForCall := fake.manifestPathArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeReleaseIndex) ManifestPathReturns(result1 string) {
	fake.manifestPathMutex.Lock()
	defer fake.manifestPathMutex.Unlock()
	fake.ManifestPathStub = nil
	fake.manifestPathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeReleaseIndex) ManifestPathReturnsOnCall(i int, result1 string) {
	fake.manifestPathMutex.Lock()
	defer fake.manifestPathMutex.Unlock()
	fake.ManifestPathStub = nil
	if fake.manifestPathReturnsOnCall == nil {
		fake.manifestPathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.manifestPathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeReleaseIndex) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.containsMutex.RLock()
	defer fake.containsMutex.RUnlock()
	fake.lastVersionMutex.RLock()
	defer fake.lastVersionMutex.RUnlock()
	fake.manifestPathMutex.RLock()
	defer fake.manifestPathMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReleaseIndex) recordInvocation(key string, args []interface{}) {
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

var _ releasedir.ReleaseIndex = new(FakeReleaseIndex)
