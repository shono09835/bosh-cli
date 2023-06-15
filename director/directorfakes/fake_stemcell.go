// Code generated by counterfeiter. DO NOT EDIT.
package directorfakes

import (
	"sync"

	"github.com/shono09835/bosh-cli/v7/director"
	"github.com/cppforlife/go-semi-semantic/version"
)

type FakeStemcell struct {
	CIDStub        func() string
	cIDMutex       sync.RWMutex
	cIDArgsForCall []struct {
	}
	cIDReturns struct {
		result1 string
	}
	cIDReturnsOnCall map[int]struct {
		result1 string
	}
	CPIStub        func() string
	cPIMutex       sync.RWMutex
	cPIArgsForCall []struct {
	}
	cPIReturns struct {
		result1 string
	}
	cPIReturnsOnCall map[int]struct {
		result1 string
	}
	DeleteStub        func(bool) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 bool
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
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
	OSNameStub        func() string
	oSNameMutex       sync.RWMutex
	oSNameArgsForCall []struct {
	}
	oSNameReturns struct {
		result1 string
	}
	oSNameReturnsOnCall map[int]struct {
		result1 string
	}
	VersionStub        func() version.Version
	versionMutex       sync.RWMutex
	versionArgsForCall []struct {
	}
	versionReturns struct {
		result1 version.Version
	}
	versionReturnsOnCall map[int]struct {
		result1 version.Version
	}
	VersionMarkStub        func(string) string
	versionMarkMutex       sync.RWMutex
	versionMarkArgsForCall []struct {
		arg1 string
	}
	versionMarkReturns struct {
		result1 string
	}
	versionMarkReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStemcell) CID() string {
	fake.cIDMutex.Lock()
	ret, specificReturn := fake.cIDReturnsOnCall[len(fake.cIDArgsForCall)]
	fake.cIDArgsForCall = append(fake.cIDArgsForCall, struct {
	}{})
	stub := fake.CIDStub
	fakeReturns := fake.cIDReturns
	fake.recordInvocation("CID", []interface{}{})
	fake.cIDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStemcell) CIDCallCount() int {
	fake.cIDMutex.RLock()
	defer fake.cIDMutex.RUnlock()
	return len(fake.cIDArgsForCall)
}

func (fake *FakeStemcell) CIDCalls(stub func() string) {
	fake.cIDMutex.Lock()
	defer fake.cIDMutex.Unlock()
	fake.CIDStub = stub
}

func (fake *FakeStemcell) CIDReturns(result1 string) {
	fake.cIDMutex.Lock()
	defer fake.cIDMutex.Unlock()
	fake.CIDStub = nil
	fake.cIDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStemcell) CIDReturnsOnCall(i int, result1 string) {
	fake.cIDMutex.Lock()
	defer fake.cIDMutex.Unlock()
	fake.CIDStub = nil
	if fake.cIDReturnsOnCall == nil {
		fake.cIDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cIDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeStemcell) CPI() string {
	fake.cPIMutex.Lock()
	ret, specificReturn := fake.cPIReturnsOnCall[len(fake.cPIArgsForCall)]
	fake.cPIArgsForCall = append(fake.cPIArgsForCall, struct {
	}{})
	stub := fake.CPIStub
	fakeReturns := fake.cPIReturns
	fake.recordInvocation("CPI", []interface{}{})
	fake.cPIMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStemcell) CPICallCount() int {
	fake.cPIMutex.RLock()
	defer fake.cPIMutex.RUnlock()
	return len(fake.cPIArgsForCall)
}

func (fake *FakeStemcell) CPICalls(stub func() string) {
	fake.cPIMutex.Lock()
	defer fake.cPIMutex.Unlock()
	fake.CPIStub = stub
}

func (fake *FakeStemcell) CPIReturns(result1 string) {
	fake.cPIMutex.Lock()
	defer fake.cPIMutex.Unlock()
	fake.CPIStub = nil
	fake.cPIReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStemcell) CPIReturnsOnCall(i int, result1 string) {
	fake.cPIMutex.Lock()
	defer fake.cPIMutex.Unlock()
	fake.CPIStub = nil
	if fake.cPIReturnsOnCall == nil {
		fake.cPIReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.cPIReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeStemcell) Delete(arg1 bool) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStemcell) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeStemcell) DeleteCalls(stub func(bool) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeStemcell) DeleteArgsForCall(i int) bool {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStemcell) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStemcell) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStemcell) Name() string {
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

func (fake *FakeStemcell) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeStemcell) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeStemcell) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStemcell) NameReturnsOnCall(i int, result1 string) {
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

func (fake *FakeStemcell) OSName() string {
	fake.oSNameMutex.Lock()
	ret, specificReturn := fake.oSNameReturnsOnCall[len(fake.oSNameArgsForCall)]
	fake.oSNameArgsForCall = append(fake.oSNameArgsForCall, struct {
	}{})
	stub := fake.OSNameStub
	fakeReturns := fake.oSNameReturns
	fake.recordInvocation("OSName", []interface{}{})
	fake.oSNameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStemcell) OSNameCallCount() int {
	fake.oSNameMutex.RLock()
	defer fake.oSNameMutex.RUnlock()
	return len(fake.oSNameArgsForCall)
}

func (fake *FakeStemcell) OSNameCalls(stub func() string) {
	fake.oSNameMutex.Lock()
	defer fake.oSNameMutex.Unlock()
	fake.OSNameStub = stub
}

func (fake *FakeStemcell) OSNameReturns(result1 string) {
	fake.oSNameMutex.Lock()
	defer fake.oSNameMutex.Unlock()
	fake.OSNameStub = nil
	fake.oSNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStemcell) OSNameReturnsOnCall(i int, result1 string) {
	fake.oSNameMutex.Lock()
	defer fake.oSNameMutex.Unlock()
	fake.OSNameStub = nil
	if fake.oSNameReturnsOnCall == nil {
		fake.oSNameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.oSNameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeStemcell) Version() version.Version {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct {
	}{})
	stub := fake.VersionStub
	fakeReturns := fake.versionReturns
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStemcell) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *FakeStemcell) VersionCalls(stub func() version.Version) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = stub
}

func (fake *FakeStemcell) VersionReturns(result1 version.Version) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 version.Version
	}{result1}
}

func (fake *FakeStemcell) VersionReturnsOnCall(i int, result1 version.Version) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 version.Version
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 version.Version
	}{result1}
}

func (fake *FakeStemcell) VersionMark(arg1 string) string {
	fake.versionMarkMutex.Lock()
	ret, specificReturn := fake.versionMarkReturnsOnCall[len(fake.versionMarkArgsForCall)]
	fake.versionMarkArgsForCall = append(fake.versionMarkArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.VersionMarkStub
	fakeReturns := fake.versionMarkReturns
	fake.recordInvocation("VersionMark", []interface{}{arg1})
	fake.versionMarkMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeStemcell) VersionMarkCallCount() int {
	fake.versionMarkMutex.RLock()
	defer fake.versionMarkMutex.RUnlock()
	return len(fake.versionMarkArgsForCall)
}

func (fake *FakeStemcell) VersionMarkCalls(stub func(string) string) {
	fake.versionMarkMutex.Lock()
	defer fake.versionMarkMutex.Unlock()
	fake.VersionMarkStub = stub
}

func (fake *FakeStemcell) VersionMarkArgsForCall(i int) string {
	fake.versionMarkMutex.RLock()
	defer fake.versionMarkMutex.RUnlock()
	argsForCall := fake.versionMarkArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeStemcell) VersionMarkReturns(result1 string) {
	fake.versionMarkMutex.Lock()
	defer fake.versionMarkMutex.Unlock()
	fake.VersionMarkStub = nil
	fake.versionMarkReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeStemcell) VersionMarkReturnsOnCall(i int, result1 string) {
	fake.versionMarkMutex.Lock()
	defer fake.versionMarkMutex.Unlock()
	fake.VersionMarkStub = nil
	if fake.versionMarkReturnsOnCall == nil {
		fake.versionMarkReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.versionMarkReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeStemcell) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cIDMutex.RLock()
	defer fake.cIDMutex.RUnlock()
	fake.cPIMutex.RLock()
	defer fake.cPIMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.oSNameMutex.RLock()
	defer fake.oSNameMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	fake.versionMarkMutex.RLock()
	defer fake.versionMarkMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStemcell) recordInvocation(key string, args []interface{}) {
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

var _ director.Stemcell = new(FakeStemcell)
