// Code generated by counterfeiter. DO NOT EDIT.
package taskfakes

import (
	"sync"

	"github.com/shono09835/bosh-cli/v7/ui/task"
)

type FakeReporter struct {
	TaskFinishedStub        func(int, string)
	taskFinishedMutex       sync.RWMutex
	taskFinishedArgsForCall []struct {
		arg1 int
		arg2 string
	}
	TaskOutputChunkStub        func(int, []byte)
	taskOutputChunkMutex       sync.RWMutex
	taskOutputChunkArgsForCall []struct {
		arg1 int
		arg2 []byte
	}
	TaskStartedStub        func(int)
	taskStartedMutex       sync.RWMutex
	taskStartedArgsForCall []struct {
		arg1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReporter) TaskFinished(arg1 int, arg2 string) {
	fake.taskFinishedMutex.Lock()
	fake.taskFinishedArgsForCall = append(fake.taskFinishedArgsForCall, struct {
		arg1 int
		arg2 string
	}{arg1, arg2})
	stub := fake.TaskFinishedStub
	fake.recordInvocation("TaskFinished", []interface{}{arg1, arg2})
	fake.taskFinishedMutex.Unlock()
	if stub != nil {
		fake.TaskFinishedStub(arg1, arg2)
	}
}

func (fake *FakeReporter) TaskFinishedCallCount() int {
	fake.taskFinishedMutex.RLock()
	defer fake.taskFinishedMutex.RUnlock()
	return len(fake.taskFinishedArgsForCall)
}

func (fake *FakeReporter) TaskFinishedCalls(stub func(int, string)) {
	fake.taskFinishedMutex.Lock()
	defer fake.taskFinishedMutex.Unlock()
	fake.TaskFinishedStub = stub
}

func (fake *FakeReporter) TaskFinishedArgsForCall(i int) (int, string) {
	fake.taskFinishedMutex.RLock()
	defer fake.taskFinishedMutex.RUnlock()
	argsForCall := fake.taskFinishedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeReporter) TaskOutputChunk(arg1 int, arg2 []byte) {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.taskOutputChunkMutex.Lock()
	fake.taskOutputChunkArgsForCall = append(fake.taskOutputChunkArgsForCall, struct {
		arg1 int
		arg2 []byte
	}{arg1, arg2Copy})
	stub := fake.TaskOutputChunkStub
	fake.recordInvocation("TaskOutputChunk", []interface{}{arg1, arg2Copy})
	fake.taskOutputChunkMutex.Unlock()
	if stub != nil {
		fake.TaskOutputChunkStub(arg1, arg2)
	}
}

func (fake *FakeReporter) TaskOutputChunkCallCount() int {
	fake.taskOutputChunkMutex.RLock()
	defer fake.taskOutputChunkMutex.RUnlock()
	return len(fake.taskOutputChunkArgsForCall)
}

func (fake *FakeReporter) TaskOutputChunkCalls(stub func(int, []byte)) {
	fake.taskOutputChunkMutex.Lock()
	defer fake.taskOutputChunkMutex.Unlock()
	fake.TaskOutputChunkStub = stub
}

func (fake *FakeReporter) TaskOutputChunkArgsForCall(i int) (int, []byte) {
	fake.taskOutputChunkMutex.RLock()
	defer fake.taskOutputChunkMutex.RUnlock()
	argsForCall := fake.taskOutputChunkArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeReporter) TaskStarted(arg1 int) {
	fake.taskStartedMutex.Lock()
	fake.taskStartedArgsForCall = append(fake.taskStartedArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.TaskStartedStub
	fake.recordInvocation("TaskStarted", []interface{}{arg1})
	fake.taskStartedMutex.Unlock()
	if stub != nil {
		fake.TaskStartedStub(arg1)
	}
}

func (fake *FakeReporter) TaskStartedCallCount() int {
	fake.taskStartedMutex.RLock()
	defer fake.taskStartedMutex.RUnlock()
	return len(fake.taskStartedArgsForCall)
}

func (fake *FakeReporter) TaskStartedCalls(stub func(int)) {
	fake.taskStartedMutex.Lock()
	defer fake.taskStartedMutex.Unlock()
	fake.TaskStartedStub = stub
}

func (fake *FakeReporter) TaskStartedArgsForCall(i int) int {
	fake.taskStartedMutex.RLock()
	defer fake.taskStartedMutex.RUnlock()
	argsForCall := fake.taskStartedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeReporter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.taskFinishedMutex.RLock()
	defer fake.taskFinishedMutex.RUnlock()
	fake.taskOutputChunkMutex.RLock()
	defer fake.taskOutputChunkMutex.RUnlock()
	fake.taskStartedMutex.RLock()
	defer fake.taskStartedMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReporter) recordInvocation(key string, args []interface{}) {
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

var _ task.Reporter = new(FakeReporter)
