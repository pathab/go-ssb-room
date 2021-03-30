// Code generated by counterfeiter. DO NOT EDIT.
package mockdb

import (
	"context"
	"sync"

	"github.com/ssb-ngi-pointer/go-ssb-room/roomdb"
	refs "go.mindeco.de/ssb-refs"
)

type FakeMembersService struct {
	AddStub        func(context.Context, refs.FeedRef, roomdb.Role) (int64, error)
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 context.Context
		arg2 refs.FeedRef
		arg3 roomdb.Role
	}
	addReturns struct {
		result1 int64
		result2 error
	}
	addReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	GetByFeedStub        func(context.Context, refs.FeedRef) (roomdb.Member, error)
	getByFeedMutex       sync.RWMutex
	getByFeedArgsForCall []struct {
		arg1 context.Context
		arg2 refs.FeedRef
	}
	getByFeedReturns struct {
		result1 roomdb.Member
		result2 error
	}
	getByFeedReturnsOnCall map[int]struct {
		result1 roomdb.Member
		result2 error
	}
	GetByIDStub        func(context.Context, int64) (roomdb.Member, error)
	getByIDMutex       sync.RWMutex
	getByIDArgsForCall []struct {
		arg1 context.Context
		arg2 int64
	}
	getByIDReturns struct {
		result1 roomdb.Member
		result2 error
	}
	getByIDReturnsOnCall map[int]struct {
		result1 roomdb.Member
		result2 error
	}
	ListStub        func(context.Context) ([]roomdb.Member, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 context.Context
	}
	listReturns struct {
		result1 []roomdb.Member
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []roomdb.Member
		result2 error
	}
	RemoveFeedStub        func(context.Context, refs.FeedRef) error
	removeFeedMutex       sync.RWMutex
	removeFeedArgsForCall []struct {
		arg1 context.Context
		arg2 refs.FeedRef
	}
	removeFeedReturns struct {
		result1 error
	}
	removeFeedReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveIDStub        func(context.Context, int64) error
	removeIDMutex       sync.RWMutex
	removeIDArgsForCall []struct {
		arg1 context.Context
		arg2 int64
	}
	removeIDReturns struct {
		result1 error
	}
	removeIDReturnsOnCall map[int]struct {
		result1 error
	}
	SetRoleStub        func(context.Context, int64, roomdb.Role) error
	setRoleMutex       sync.RWMutex
	setRoleArgsForCall []struct {
		arg1 context.Context
		arg2 int64
		arg3 roomdb.Role
	}
	setRoleReturns struct {
		result1 error
	}
	setRoleReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMembersService) Add(arg1 context.Context, arg2 refs.FeedRef, arg3 roomdb.Role) (int64, error) {
	fake.addMutex.Lock()
	ret, specificReturn := fake.addReturnsOnCall[len(fake.addArgsForCall)]
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		arg1 context.Context
		arg2 refs.FeedRef
		arg3 roomdb.Role
	}{arg1, arg2, arg3})
	stub := fake.AddStub
	fakeReturns := fake.addReturns
	fake.recordInvocation("Add", []interface{}{arg1, arg2, arg3})
	fake.addMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMembersService) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *FakeMembersService) AddCalls(stub func(context.Context, refs.FeedRef, roomdb.Role) (int64, error)) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = stub
}

func (fake *FakeMembersService) AddArgsForCall(i int) (context.Context, refs.FeedRef, roomdb.Role) {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	argsForCall := fake.addArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeMembersService) AddReturns(result1 int64, result2 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeMembersService) AddReturnsOnCall(i int, result1 int64, result2 error) {
	fake.addMutex.Lock()
	defer fake.addMutex.Unlock()
	fake.AddStub = nil
	if fake.addReturnsOnCall == nil {
		fake.addReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.addReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *FakeMembersService) GetByFeed(arg1 context.Context, arg2 refs.FeedRef) (roomdb.Member, error) {
	fake.getByFeedMutex.Lock()
	ret, specificReturn := fake.getByFeedReturnsOnCall[len(fake.getByFeedArgsForCall)]
	fake.getByFeedArgsForCall = append(fake.getByFeedArgsForCall, struct {
		arg1 context.Context
		arg2 refs.FeedRef
	}{arg1, arg2})
	stub := fake.GetByFeedStub
	fakeReturns := fake.getByFeedReturns
	fake.recordInvocation("GetByFeed", []interface{}{arg1, arg2})
	fake.getByFeedMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMembersService) GetByFeedCallCount() int {
	fake.getByFeedMutex.RLock()
	defer fake.getByFeedMutex.RUnlock()
	return len(fake.getByFeedArgsForCall)
}

func (fake *FakeMembersService) GetByFeedCalls(stub func(context.Context, refs.FeedRef) (roomdb.Member, error)) {
	fake.getByFeedMutex.Lock()
	defer fake.getByFeedMutex.Unlock()
	fake.GetByFeedStub = stub
}

func (fake *FakeMembersService) GetByFeedArgsForCall(i int) (context.Context, refs.FeedRef) {
	fake.getByFeedMutex.RLock()
	defer fake.getByFeedMutex.RUnlock()
	argsForCall := fake.getByFeedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMembersService) GetByFeedReturns(result1 roomdb.Member, result2 error) {
	fake.getByFeedMutex.Lock()
	defer fake.getByFeedMutex.Unlock()
	fake.GetByFeedStub = nil
	fake.getByFeedReturns = struct {
		result1 roomdb.Member
		result2 error
	}{result1, result2}
}

func (fake *FakeMembersService) GetByFeedReturnsOnCall(i int, result1 roomdb.Member, result2 error) {
	fake.getByFeedMutex.Lock()
	defer fake.getByFeedMutex.Unlock()
	fake.GetByFeedStub = nil
	if fake.getByFeedReturnsOnCall == nil {
		fake.getByFeedReturnsOnCall = make(map[int]struct {
			result1 roomdb.Member
			result2 error
		})
	}
	fake.getByFeedReturnsOnCall[i] = struct {
		result1 roomdb.Member
		result2 error
	}{result1, result2}
}

func (fake *FakeMembersService) GetByID(arg1 context.Context, arg2 int64) (roomdb.Member, error) {
	fake.getByIDMutex.Lock()
	ret, specificReturn := fake.getByIDReturnsOnCall[len(fake.getByIDArgsForCall)]
	fake.getByIDArgsForCall = append(fake.getByIDArgsForCall, struct {
		arg1 context.Context
		arg2 int64
	}{arg1, arg2})
	stub := fake.GetByIDStub
	fakeReturns := fake.getByIDReturns
	fake.recordInvocation("GetByID", []interface{}{arg1, arg2})
	fake.getByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMembersService) GetByIDCallCount() int {
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	return len(fake.getByIDArgsForCall)
}

func (fake *FakeMembersService) GetByIDCalls(stub func(context.Context, int64) (roomdb.Member, error)) {
	fake.getByIDMutex.Lock()
	defer fake.getByIDMutex.Unlock()
	fake.GetByIDStub = stub
}

func (fake *FakeMembersService) GetByIDArgsForCall(i int) (context.Context, int64) {
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	argsForCall := fake.getByIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMembersService) GetByIDReturns(result1 roomdb.Member, result2 error) {
	fake.getByIDMutex.Lock()
	defer fake.getByIDMutex.Unlock()
	fake.GetByIDStub = nil
	fake.getByIDReturns = struct {
		result1 roomdb.Member
		result2 error
	}{result1, result2}
}

func (fake *FakeMembersService) GetByIDReturnsOnCall(i int, result1 roomdb.Member, result2 error) {
	fake.getByIDMutex.Lock()
	defer fake.getByIDMutex.Unlock()
	fake.GetByIDStub = nil
	if fake.getByIDReturnsOnCall == nil {
		fake.getByIDReturnsOnCall = make(map[int]struct {
			result1 roomdb.Member
			result2 error
		})
	}
	fake.getByIDReturnsOnCall[i] = struct {
		result1 roomdb.Member
		result2 error
	}{result1, result2}
}

func (fake *FakeMembersService) List(arg1 context.Context) ([]roomdb.Member, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{arg1})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMembersService) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeMembersService) ListCalls(stub func(context.Context) ([]roomdb.Member, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeMembersService) ListArgsForCall(i int) context.Context {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMembersService) ListReturns(result1 []roomdb.Member, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []roomdb.Member
		result2 error
	}{result1, result2}
}

func (fake *FakeMembersService) ListReturnsOnCall(i int, result1 []roomdb.Member, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []roomdb.Member
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []roomdb.Member
		result2 error
	}{result1, result2}
}

func (fake *FakeMembersService) RemoveFeed(arg1 context.Context, arg2 refs.FeedRef) error {
	fake.removeFeedMutex.Lock()
	ret, specificReturn := fake.removeFeedReturnsOnCall[len(fake.removeFeedArgsForCall)]
	fake.removeFeedArgsForCall = append(fake.removeFeedArgsForCall, struct {
		arg1 context.Context
		arg2 refs.FeedRef
	}{arg1, arg2})
	stub := fake.RemoveFeedStub
	fakeReturns := fake.removeFeedReturns
	fake.recordInvocation("RemoveFeed", []interface{}{arg1, arg2})
	fake.removeFeedMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMembersService) RemoveFeedCallCount() int {
	fake.removeFeedMutex.RLock()
	defer fake.removeFeedMutex.RUnlock()
	return len(fake.removeFeedArgsForCall)
}

func (fake *FakeMembersService) RemoveFeedCalls(stub func(context.Context, refs.FeedRef) error) {
	fake.removeFeedMutex.Lock()
	defer fake.removeFeedMutex.Unlock()
	fake.RemoveFeedStub = stub
}

func (fake *FakeMembersService) RemoveFeedArgsForCall(i int) (context.Context, refs.FeedRef) {
	fake.removeFeedMutex.RLock()
	defer fake.removeFeedMutex.RUnlock()
	argsForCall := fake.removeFeedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMembersService) RemoveFeedReturns(result1 error) {
	fake.removeFeedMutex.Lock()
	defer fake.removeFeedMutex.Unlock()
	fake.RemoveFeedStub = nil
	fake.removeFeedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMembersService) RemoveFeedReturnsOnCall(i int, result1 error) {
	fake.removeFeedMutex.Lock()
	defer fake.removeFeedMutex.Unlock()
	fake.RemoveFeedStub = nil
	if fake.removeFeedReturnsOnCall == nil {
		fake.removeFeedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeFeedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMembersService) RemoveID(arg1 context.Context, arg2 int64) error {
	fake.removeIDMutex.Lock()
	ret, specificReturn := fake.removeIDReturnsOnCall[len(fake.removeIDArgsForCall)]
	fake.removeIDArgsForCall = append(fake.removeIDArgsForCall, struct {
		arg1 context.Context
		arg2 int64
	}{arg1, arg2})
	stub := fake.RemoveIDStub
	fakeReturns := fake.removeIDReturns
	fake.recordInvocation("RemoveID", []interface{}{arg1, arg2})
	fake.removeIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMembersService) RemoveIDCallCount() int {
	fake.removeIDMutex.RLock()
	defer fake.removeIDMutex.RUnlock()
	return len(fake.removeIDArgsForCall)
}

func (fake *FakeMembersService) RemoveIDCalls(stub func(context.Context, int64) error) {
	fake.removeIDMutex.Lock()
	defer fake.removeIDMutex.Unlock()
	fake.RemoveIDStub = stub
}

func (fake *FakeMembersService) RemoveIDArgsForCall(i int) (context.Context, int64) {
	fake.removeIDMutex.RLock()
	defer fake.removeIDMutex.RUnlock()
	argsForCall := fake.removeIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMembersService) RemoveIDReturns(result1 error) {
	fake.removeIDMutex.Lock()
	defer fake.removeIDMutex.Unlock()
	fake.RemoveIDStub = nil
	fake.removeIDReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMembersService) RemoveIDReturnsOnCall(i int, result1 error) {
	fake.removeIDMutex.Lock()
	defer fake.removeIDMutex.Unlock()
	fake.RemoveIDStub = nil
	if fake.removeIDReturnsOnCall == nil {
		fake.removeIDReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeIDReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMembersService) SetRole(arg1 context.Context, arg2 int64, arg3 roomdb.Role) error {
	fake.setRoleMutex.Lock()
	ret, specificReturn := fake.setRoleReturnsOnCall[len(fake.setRoleArgsForCall)]
	fake.setRoleArgsForCall = append(fake.setRoleArgsForCall, struct {
		arg1 context.Context
		arg2 int64
		arg3 roomdb.Role
	}{arg1, arg2, arg3})
	stub := fake.SetRoleStub
	fakeReturns := fake.setRoleReturns
	fake.recordInvocation("SetRole", []interface{}{arg1, arg2, arg3})
	fake.setRoleMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMembersService) SetRoleCallCount() int {
	fake.setRoleMutex.RLock()
	defer fake.setRoleMutex.RUnlock()
	return len(fake.setRoleArgsForCall)
}

func (fake *FakeMembersService) SetRoleCalls(stub func(context.Context, int64, roomdb.Role) error) {
	fake.setRoleMutex.Lock()
	defer fake.setRoleMutex.Unlock()
	fake.SetRoleStub = stub
}

func (fake *FakeMembersService) SetRoleArgsForCall(i int) (context.Context, int64, roomdb.Role) {
	fake.setRoleMutex.RLock()
	defer fake.setRoleMutex.RUnlock()
	argsForCall := fake.setRoleArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeMembersService) SetRoleReturns(result1 error) {
	fake.setRoleMutex.Lock()
	defer fake.setRoleMutex.Unlock()
	fake.SetRoleStub = nil
	fake.setRoleReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMembersService) SetRoleReturnsOnCall(i int, result1 error) {
	fake.setRoleMutex.Lock()
	defer fake.setRoleMutex.Unlock()
	fake.SetRoleStub = nil
	if fake.setRoleReturnsOnCall == nil {
		fake.setRoleReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setRoleReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMembersService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.getByFeedMutex.RLock()
	defer fake.getByFeedMutex.RUnlock()
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.removeFeedMutex.RLock()
	defer fake.removeFeedMutex.RUnlock()
	fake.removeIDMutex.RLock()
	defer fake.removeIDMutex.RUnlock()
	fake.setRoleMutex.RLock()
	defer fake.setRoleMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMembersService) recordInvocation(key string, args []interface{}) {
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

var _ roomdb.MembersService = new(FakeMembersService)
