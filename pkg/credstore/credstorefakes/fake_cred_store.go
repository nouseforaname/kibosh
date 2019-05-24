// Code generated by counterfeiter. DO NOT EDIT.
package credstorefakes

import (
	"sync"

	"github.com/cf-platform-eng/kibosh/pkg/credstore"
)

type FakeCredStore struct {
	DeleteStub        func(string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(string) (interface{}, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
	}
	getReturns struct {
		result1 interface{}
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	PutStub        func(string, interface{}) (interface{}, error)
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	putReturns struct {
		result1 interface{}
		result2 error
	}
	putReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCredStore) Delete(arg1 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakeCredStore) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeCredStore) DeleteCalls(stub func(string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeCredStore) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCredStore) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCredStore) DeleteReturnsOnCall(i int, result1 error) {
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

func (fake *FakeCredStore) Get(arg1 string) (interface{}, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Get", []interface{}{arg1})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCredStore) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeCredStore) GetCalls(stub func(string) (interface{}, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeCredStore) GetArgsForCall(i int) string {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCredStore) GetReturns(result1 interface{}, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeCredStore) GetReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeCredStore) Put(arg1 string, arg2 interface{}) (interface{}, error) {
	fake.putMutex.Lock()
	ret, specificReturn := fake.putReturnsOnCall[len(fake.putArgsForCall)]
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	fake.recordInvocation("Put", []interface{}{arg1, arg2})
	fake.putMutex.Unlock()
	if fake.PutStub != nil {
		return fake.PutStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.putReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCredStore) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeCredStore) PutCalls(stub func(string, interface{}) (interface{}, error)) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = stub
}

func (fake *FakeCredStore) PutArgsForCall(i int) (string, interface{}) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	argsForCall := fake.putArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeCredStore) PutReturns(result1 interface{}, result2 error) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeCredStore) PutReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = nil
	if fake.putReturnsOnCall == nil {
		fake.putReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.putReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeCredStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCredStore) recordInvocation(key string, args []interface{}) {
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

var _ credstore.CredStore = new(FakeCredStore)
