package fakes

import (
	"sync"

	"github.com/shono09835/bosh-cli/v7/config"
	"github.com/shono09835/bosh-cli/v7/release"
)

// FakeReleaseRepo was generated by counterfeiter
type FakeReleaseRepo struct {
	ListStub        func() ([]config.ReleaseRecord, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct{}
	listReturns     struct {
		result1 []config.ReleaseRecord
		result2 error
	}
	UpdateStub        func([]release.Release) error
	updateMutex       sync.RWMutex
	updateArgsForCall []struct {
		arg1 []release.Release
	}
	updateReturns struct {
		result1 error
	}
}

func (fake *FakeReleaseRepo) List() ([]config.ReleaseRecord, error) {
	fake.listMutex.Lock()
	fake.listArgsForCall = append(fake.listArgsForCall, struct{}{})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub()
	}
	return fake.listReturns.result1, fake.listReturns.result2
}

func (fake *FakeReleaseRepo) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeReleaseRepo) ListReturns(result1 []config.ReleaseRecord, result2 error) {
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []config.ReleaseRecord
		result2 error
	}{result1, result2}
}

func (fake *FakeReleaseRepo) Update(arg1 []release.Release) error {
	fake.updateMutex.Lock()
	fake.updateArgsForCall = append(fake.updateArgsForCall, struct {
		arg1 []release.Release
	}{arg1})
	fake.updateMutex.Unlock()
	if fake.UpdateStub != nil {
		return fake.UpdateStub(arg1)
	}
	return fake.updateReturns.result1
}

func (fake *FakeReleaseRepo) UpdateCallCount() int {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return len(fake.updateArgsForCall)
}

func (fake *FakeReleaseRepo) UpdateArgsForCall(i int) []release.Release {
	fake.updateMutex.RLock()
	defer fake.updateMutex.RUnlock()
	return fake.updateArgsForCall[i].arg1
}

func (fake *FakeReleaseRepo) UpdateReturns(result1 error) {
	fake.UpdateStub = nil
	fake.updateReturns = struct {
		result1 error
	}{result1}
}

var _ config.ReleaseRepo = new(FakeReleaseRepo)
