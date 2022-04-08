// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package crew

import (
	"context"
	"sync"
)

// Ensure, that ActionsMock does implement Actions.
// If this is not the case, regenerate this file with moq.
var _ Actions = &ActionsMock{}

// ActionsMock is a mock implementation of Actions.
//
// 	func TestSomethingThatUsesActions(t *testing.T) {
//
// 		// make and configure a mocked Actions
// 		mockedActions := &ActionsMock{
// 			CreateMemberFunc: func(ctx context.Context, op Member) (Member, error) {
// 				panic("mock out the CreateMember method")
// 			},
// 			GetMemberFunc: func(ctx context.Context, key Key) (Member, error) {
// 				panic("mock out the GetMember method")
// 			},
// 			GetMemberWithNicknamesFunc: func(ctx context.Context, operatorNickname string, storeNickname string) (Member, error) {
// 				panic("mock out the GetMemberWithNicknames method")
// 			},
// 			ListMembersFunc: func(ctx context.Context, operatorFilter Filter) ([]Member, int, error) {
// 				panic("mock out the ListMembers method")
// 			},
// 			RemoveMemberFunc: func(ctx context.Context, key Key) error {
// 				panic("mock out the RemoveMember method")
// 			},
// 			UpdateMemberFunc: func(ctx context.Context, op Member) error {
// 				panic("mock out the UpdateMember method")
// 			},
// 		}
//
// 		// use mockedActions in code that requires Actions
// 		// and then make assertions.
//
// 	}
type ActionsMock struct {
	// CreateMemberFunc mocks the CreateMember method.
	CreateMemberFunc func(ctx context.Context, op Member) (Member, error)

	// GetMemberFunc mocks the GetMember method.
	GetMemberFunc func(ctx context.Context, key Key) (Member, error)

	// GetMemberWithNicknamesFunc mocks the GetMemberWithNicknames method.
	GetMemberWithNicknamesFunc func(ctx context.Context, operatorNickname string, storeNickname string) (Member, error)

	// ListMembersFunc mocks the ListMembers method.
	ListMembersFunc func(ctx context.Context, operatorFilter Filter) ([]Member, int, error)

	// RemoveMemberFunc mocks the RemoveMember method.
	RemoveMemberFunc func(ctx context.Context, key Key) error

	// UpdateMemberFunc mocks the UpdateMember method.
	UpdateMemberFunc func(ctx context.Context, op Member) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateMember holds details about calls to the CreateMember method.
		CreateMember []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Op is the op argument value.
			Op Member
		}
		// GetMember holds details about calls to the GetMember method.
		GetMember []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Key is the key argument value.
			Key Key
		}
		// GetMemberWithNicknames holds details about calls to the GetMemberWithNicknames method.
		GetMemberWithNicknames []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// OperatorNickname is the operatorNickname argument value.
			OperatorNickname string
			// StoreNickname is the storeNickname argument value.
			StoreNickname string
		}
		// ListMembers holds details about calls to the ListMembers method.
		ListMembers []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// OperatorFilter is the operatorFilter argument value.
			OperatorFilter Filter
		}
		// RemoveMember holds details about calls to the RemoveMember method.
		RemoveMember []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Key is the key argument value.
			Key Key
		}
		// UpdateMember holds details about calls to the UpdateMember method.
		UpdateMember []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Op is the op argument value.
			Op Member
		}
	}
	lockCreateMember           sync.RWMutex
	lockGetMember              sync.RWMutex
	lockGetMemberWithNicknames sync.RWMutex
	lockListMembers            sync.RWMutex
	lockRemoveMember           sync.RWMutex
	lockUpdateMember           sync.RWMutex
}

// CreateMember calls CreateMemberFunc.
func (mock *ActionsMock) CreateMember(ctx context.Context, op Member) (Member, error) {
	if mock.CreateMemberFunc == nil {
		panic("ActionsMock.CreateMemberFunc: method is nil but Actions.CreateMember was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Op  Member
	}{
		Ctx: ctx,
		Op:  op,
	}
	mock.lockCreateMember.Lock()
	mock.calls.CreateMember = append(mock.calls.CreateMember, callInfo)
	mock.lockCreateMember.Unlock()
	return mock.CreateMemberFunc(ctx, op)
}

// CreateMemberCalls gets all the calls that were made to CreateMember.
// Check the length with:
//     len(mockedActions.CreateMemberCalls())
func (mock *ActionsMock) CreateMemberCalls() []struct {
	Ctx context.Context
	Op  Member
} {
	var calls []struct {
		Ctx context.Context
		Op  Member
	}
	mock.lockCreateMember.RLock()
	calls = mock.calls.CreateMember
	mock.lockCreateMember.RUnlock()
	return calls
}

// GetMember calls GetMemberFunc.
func (mock *ActionsMock) GetMember(ctx context.Context, key Key) (Member, error) {
	if mock.GetMemberFunc == nil {
		panic("ActionsMock.GetMemberFunc: method is nil but Actions.GetMember was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Key Key
	}{
		Ctx: ctx,
		Key: key,
	}
	mock.lockGetMember.Lock()
	mock.calls.GetMember = append(mock.calls.GetMember, callInfo)
	mock.lockGetMember.Unlock()
	return mock.GetMemberFunc(ctx, key)
}

// GetMemberCalls gets all the calls that were made to GetMember.
// Check the length with:
//     len(mockedActions.GetMemberCalls())
func (mock *ActionsMock) GetMemberCalls() []struct {
	Ctx context.Context
	Key Key
} {
	var calls []struct {
		Ctx context.Context
		Key Key
	}
	mock.lockGetMember.RLock()
	calls = mock.calls.GetMember
	mock.lockGetMember.RUnlock()
	return calls
}

// GetMemberWithNicknames calls GetMemberWithNicknamesFunc.
func (mock *ActionsMock) GetMemberWithNicknames(ctx context.Context, operatorNickname string, storeNickname string) (Member, error) {
	if mock.GetMemberWithNicknamesFunc == nil {
		panic("ActionsMock.GetMemberWithNicknamesFunc: method is nil but Actions.GetMemberWithNicknames was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		OperatorNickname string
		StoreNickname    string
	}{
		Ctx:              ctx,
		OperatorNickname: operatorNickname,
		StoreNickname:    storeNickname,
	}
	mock.lockGetMemberWithNicknames.Lock()
	mock.calls.GetMemberWithNicknames = append(mock.calls.GetMemberWithNicknames, callInfo)
	mock.lockGetMemberWithNicknames.Unlock()
	return mock.GetMemberWithNicknamesFunc(ctx, operatorNickname, storeNickname)
}

// GetMemberWithNicknamesCalls gets all the calls that were made to GetMemberWithNicknames.
// Check the length with:
//     len(mockedActions.GetMemberWithNicknamesCalls())
func (mock *ActionsMock) GetMemberWithNicknamesCalls() []struct {
	Ctx              context.Context
	OperatorNickname string
	StoreNickname    string
} {
	var calls []struct {
		Ctx              context.Context
		OperatorNickname string
		StoreNickname    string
	}
	mock.lockGetMemberWithNicknames.RLock()
	calls = mock.calls.GetMemberWithNicknames
	mock.lockGetMemberWithNicknames.RUnlock()
	return calls
}

// ListMembers calls ListMembersFunc.
func (mock *ActionsMock) ListMembers(ctx context.Context, operatorFilter Filter) ([]Member, int, error) {
	if mock.ListMembersFunc == nil {
		panic("ActionsMock.ListMembersFunc: method is nil but Actions.ListMembers was just called")
	}
	callInfo := struct {
		Ctx            context.Context
		OperatorFilter Filter
	}{
		Ctx:            ctx,
		OperatorFilter: operatorFilter,
	}
	mock.lockListMembers.Lock()
	mock.calls.ListMembers = append(mock.calls.ListMembers, callInfo)
	mock.lockListMembers.Unlock()
	return mock.ListMembersFunc(ctx, operatorFilter)
}

// ListMembersCalls gets all the calls that were made to ListMembers.
// Check the length with:
//     len(mockedActions.ListMembersCalls())
func (mock *ActionsMock) ListMembersCalls() []struct {
	Ctx            context.Context
	OperatorFilter Filter
} {
	var calls []struct {
		Ctx            context.Context
		OperatorFilter Filter
	}
	mock.lockListMembers.RLock()
	calls = mock.calls.ListMembers
	mock.lockListMembers.RUnlock()
	return calls
}

// RemoveMember calls RemoveMemberFunc.
func (mock *ActionsMock) RemoveMember(ctx context.Context, key Key) error {
	if mock.RemoveMemberFunc == nil {
		panic("ActionsMock.RemoveMemberFunc: method is nil but Actions.RemoveMember was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Key Key
	}{
		Ctx: ctx,
		Key: key,
	}
	mock.lockRemoveMember.Lock()
	mock.calls.RemoveMember = append(mock.calls.RemoveMember, callInfo)
	mock.lockRemoveMember.Unlock()
	return mock.RemoveMemberFunc(ctx, key)
}

// RemoveMemberCalls gets all the calls that were made to RemoveMember.
// Check the length with:
//     len(mockedActions.RemoveMemberCalls())
func (mock *ActionsMock) RemoveMemberCalls() []struct {
	Ctx context.Context
	Key Key
} {
	var calls []struct {
		Ctx context.Context
		Key Key
	}
	mock.lockRemoveMember.RLock()
	calls = mock.calls.RemoveMember
	mock.lockRemoveMember.RUnlock()
	return calls
}

// UpdateMember calls UpdateMemberFunc.
func (mock *ActionsMock) UpdateMember(ctx context.Context, op Member) error {
	if mock.UpdateMemberFunc == nil {
		panic("ActionsMock.UpdateMemberFunc: method is nil but Actions.UpdateMember was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Op  Member
	}{
		Ctx: ctx,
		Op:  op,
	}
	mock.lockUpdateMember.Lock()
	mock.calls.UpdateMember = append(mock.calls.UpdateMember, callInfo)
	mock.lockUpdateMember.Unlock()
	return mock.UpdateMemberFunc(ctx, op)
}

// UpdateMemberCalls gets all the calls that were made to UpdateMember.
// Check the length with:
//     len(mockedActions.UpdateMemberCalls())
func (mock *ActionsMock) UpdateMemberCalls() []struct {
	Ctx context.Context
	Op  Member
} {
	var calls []struct {
		Ctx context.Context
		Op  Member
	}
	mock.lockUpdateMember.RLock()
	calls = mock.calls.UpdateMember
	mock.lockUpdateMember.RUnlock()
	return calls
}