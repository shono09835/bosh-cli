// Code generated by counterfeiter. DO NOT EDIT.
package pkgfakes

import (
	"sync"

	"github.com/shono09835/bosh-cli/v7/release/pkg"
)

type FakeCompilable struct {
	ArchiveDigestStub        func() string
	archiveDigestMutex       sync.RWMutex
	archiveDigestArgsForCall []struct {
	}
	archiveDigestReturns struct {
		result1 string
	}
	archiveDigestReturnsOnCall map[int]struct {
		result1 string
	}
	ArchivePathStub        func() string
	archivePathMutex       sync.RWMutex
	archivePathArgsForCall []struct {
	}
	archivePathReturns struct {
		result1 string
	}
	archivePathReturnsOnCall map[int]struct {
		result1 string
	}
	DepsStub        func() []pkg.Compilable
	depsMutex       sync.RWMutex
	depsArgsForCall []struct {
	}
	depsReturns struct {
		result1 []pkg.Compilable
	}
	depsReturnsOnCall map[int]struct {
		result1 []pkg.Compilable
	}
	FingerprintStub        func() string
	fingerprintMutex       sync.RWMutex
	fingerprintArgsForCall []struct {
	}
	fingerprintReturns struct {
		result1 string
	}
	fingerprintReturnsOnCall map[int]struct {
		result1 string
	}
	IsCompiledStub        func() bool
	isCompiledMutex       sync.RWMutex
	isCompiledArgsForCall []struct {
	}
	isCompiledReturns struct {
		result1 bool
	}
	isCompiledReturnsOnCall map[int]struct {
		result1 bool
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCompilable) ArchiveDigest() string {
	fake.archiveDigestMutex.Lock()
	ret, specificReturn := fake.archiveDigestReturnsOnCall[len(fake.archiveDigestArgsForCall)]
	fake.archiveDigestArgsForCall = append(fake.archiveDigestArgsForCall, struct {
	}{})
	stub := fake.ArchiveDigestStub
	fakeReturns := fake.archiveDigestReturns
	fake.recordInvocation("ArchiveDigest", []interface{}{})
	fake.archiveDigestMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCompilable) ArchiveDigestCallCount() int {
	fake.archiveDigestMutex.RLock()
	defer fake.archiveDigestMutex.RUnlock()
	return len(fake.archiveDigestArgsForCall)
}

func (fake *FakeCompilable) ArchiveDigestCalls(stub func() string) {
	fake.archiveDigestMutex.Lock()
	defer fake.archiveDigestMutex.Unlock()
	fake.ArchiveDigestStub = stub
}

func (fake *FakeCompilable) ArchiveDigestReturns(result1 string) {
	fake.archiveDigestMutex.Lock()
	defer fake.archiveDigestMutex.Unlock()
	fake.ArchiveDigestStub = nil
	fake.archiveDigestReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCompilable) ArchiveDigestReturnsOnCall(i int, result1 string) {
	fake.archiveDigestMutex.Lock()
	defer fake.archiveDigestMutex.Unlock()
	fake.ArchiveDigestStub = nil
	if fake.archiveDigestReturnsOnCall == nil {
		fake.archiveDigestReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.archiveDigestReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCompilable) ArchivePath() string {
	fake.archivePathMutex.Lock()
	ret, specificReturn := fake.archivePathReturnsOnCall[len(fake.archivePathArgsForCall)]
	fake.archivePathArgsForCall = append(fake.archivePathArgsForCall, struct {
	}{})
	stub := fake.ArchivePathStub
	fakeReturns := fake.archivePathReturns
	fake.recordInvocation("ArchivePath", []interface{}{})
	fake.archivePathMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCompilable) ArchivePathCallCount() int {
	fake.archivePathMutex.RLock()
	defer fake.archivePathMutex.RUnlock()
	return len(fake.archivePathArgsForCall)
}

func (fake *FakeCompilable) ArchivePathCalls(stub func() string) {
	fake.archivePathMutex.Lock()
	defer fake.archivePathMutex.Unlock()
	fake.ArchivePathStub = stub
}

func (fake *FakeCompilable) ArchivePathReturns(result1 string) {
	fake.archivePathMutex.Lock()
	defer fake.archivePathMutex.Unlock()
	fake.ArchivePathStub = nil
	fake.archivePathReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCompilable) ArchivePathReturnsOnCall(i int, result1 string) {
	fake.archivePathMutex.Lock()
	defer fake.archivePathMutex.Unlock()
	fake.ArchivePathStub = nil
	if fake.archivePathReturnsOnCall == nil {
		fake.archivePathReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.archivePathReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCompilable) Deps() []pkg.Compilable {
	fake.depsMutex.Lock()
	ret, specificReturn := fake.depsReturnsOnCall[len(fake.depsArgsForCall)]
	fake.depsArgsForCall = append(fake.depsArgsForCall, struct {
	}{})
	stub := fake.DepsStub
	fakeReturns := fake.depsReturns
	fake.recordInvocation("Deps", []interface{}{})
	fake.depsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCompilable) DepsCallCount() int {
	fake.depsMutex.RLock()
	defer fake.depsMutex.RUnlock()
	return len(fake.depsArgsForCall)
}

func (fake *FakeCompilable) DepsCalls(stub func() []pkg.Compilable) {
	fake.depsMutex.Lock()
	defer fake.depsMutex.Unlock()
	fake.DepsStub = stub
}

func (fake *FakeCompilable) DepsReturns(result1 []pkg.Compilable) {
	fake.depsMutex.Lock()
	defer fake.depsMutex.Unlock()
	fake.DepsStub = nil
	fake.depsReturns = struct {
		result1 []pkg.Compilable
	}{result1}
}

func (fake *FakeCompilable) DepsReturnsOnCall(i int, result1 []pkg.Compilable) {
	fake.depsMutex.Lock()
	defer fake.depsMutex.Unlock()
	fake.DepsStub = nil
	if fake.depsReturnsOnCall == nil {
		fake.depsReturnsOnCall = make(map[int]struct {
			result1 []pkg.Compilable
		})
	}
	fake.depsReturnsOnCall[i] = struct {
		result1 []pkg.Compilable
	}{result1}
}

func (fake *FakeCompilable) Fingerprint() string {
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
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCompilable) FingerprintCallCount() int {
	fake.fingerprintMutex.RLock()
	defer fake.fingerprintMutex.RUnlock()
	return len(fake.fingerprintArgsForCall)
}

func (fake *FakeCompilable) FingerprintCalls(stub func() string) {
	fake.fingerprintMutex.Lock()
	defer fake.fingerprintMutex.Unlock()
	fake.FingerprintStub = stub
}

func (fake *FakeCompilable) FingerprintReturns(result1 string) {
	fake.fingerprintMutex.Lock()
	defer fake.fingerprintMutex.Unlock()
	fake.FingerprintStub = nil
	fake.fingerprintReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCompilable) FingerprintReturnsOnCall(i int, result1 string) {
	fake.fingerprintMutex.Lock()
	defer fake.fingerprintMutex.Unlock()
	fake.FingerprintStub = nil
	if fake.fingerprintReturnsOnCall == nil {
		fake.fingerprintReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.fingerprintReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCompilable) IsCompiled() bool {
	fake.isCompiledMutex.Lock()
	ret, specificReturn := fake.isCompiledReturnsOnCall[len(fake.isCompiledArgsForCall)]
	fake.isCompiledArgsForCall = append(fake.isCompiledArgsForCall, struct {
	}{})
	stub := fake.IsCompiledStub
	fakeReturns := fake.isCompiledReturns
	fake.recordInvocation("IsCompiled", []interface{}{})
	fake.isCompiledMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCompilable) IsCompiledCallCount() int {
	fake.isCompiledMutex.RLock()
	defer fake.isCompiledMutex.RUnlock()
	return len(fake.isCompiledArgsForCall)
}

func (fake *FakeCompilable) IsCompiledCalls(stub func() bool) {
	fake.isCompiledMutex.Lock()
	defer fake.isCompiledMutex.Unlock()
	fake.IsCompiledStub = stub
}

func (fake *FakeCompilable) IsCompiledReturns(result1 bool) {
	fake.isCompiledMutex.Lock()
	defer fake.isCompiledMutex.Unlock()
	fake.IsCompiledStub = nil
	fake.isCompiledReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCompilable) IsCompiledReturnsOnCall(i int, result1 bool) {
	fake.isCompiledMutex.Lock()
	defer fake.isCompiledMutex.Unlock()
	fake.IsCompiledStub = nil
	if fake.isCompiledReturnsOnCall == nil {
		fake.isCompiledReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isCompiledReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeCompilable) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	stub := fake.NameStub
	fakeReturns := fake.nameReturns
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCompilable) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeCompilable) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeCompilable) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeCompilable) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeCompilable) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.archiveDigestMutex.RLock()
	defer fake.archiveDigestMutex.RUnlock()
	fake.archivePathMutex.RLock()
	defer fake.archivePathMutex.RUnlock()
	fake.depsMutex.RLock()
	defer fake.depsMutex.RUnlock()
	fake.fingerprintMutex.RLock()
	defer fake.fingerprintMutex.RUnlock()
	fake.isCompiledMutex.RLock()
	defer fake.isCompiledMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCompilable) recordInvocation(key string, args []interface{}) {
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

var _ pkg.Compilable = new(FakeCompilable)
