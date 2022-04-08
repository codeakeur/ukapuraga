// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package crew

import (
	"context"
	"gomies/app/core/entities/iam/crew"
	"gomies/pkg/sdk/session"
	"sync"
)

// Ensure, that WorkflowMock does implement Workflow.
// If this is not the case, regenerate this file with moq.
var _ Workflow = &WorkflowMock{}

// WorkflowMock is a mock implementation of Workflow.
//
// 	func TestSomethingThatUsesWorkflow(t *testing.T) {
//
// 		// make and configure a mocked Workflow
// 		mockedWorkflow := &WorkflowMock{
// 			AuthenticateMemberFunc: func(ctx context.Context, auth AuthRequest) (session.Session, error) {
// 				panic("mock out the AuthenticateMember method")
// 			},
// 			CreateMemberFunc: func(ctx context.Context, op crew.Member) (crew.Member, error) {
// 				panic("mock out the CreateMember method")
// 			},
// 			GetMemberFunc: func(ctx context.Context, key crew.Key) (crew.Member, error) {
// 				panic("mock out the GetMember method")
// 			},
// 			ListMembersFunc: func(ctx context.Context, operatorFilter crew.Filter) ([]crew.Member, int, error) {
// 				panic("mock out the ListMembers method")
// 			},
// 			RemoveMemberFunc: func(ctx context.Context, key crew.Key) error {
// 				panic("mock out the RemoveMember method")
// 			},
// 			UpdateMemberFunc: func(ctx context.Context, op crew.Member) error {
// 				panic("mock out the UpdateMember method")
// 			},
// 		}
//
// 		// use mockedWorkflow in code that requires Workflow
// 		// and then make assertions.
//
// 	}
type WorkflowMock struct {
	// AuthenticateMemberFunc mocks the AuthenticateMember method.
	AuthenticateMemberFunc func(ctx context.Context, auth AuthRequest) (session.Session, error)

	// CreateMemberFunc mocks the CreateMember method.
	CreateMemberFunc func(ctx context.Context, op crew.Member) (crew.Member, error)

	// GetMemberFunc mocks the GetMember method.
	GetMemberFunc func(ctx context.Context, key crew.Key) (crew.Member, error)

	// ListMembersFunc mocks the ListMembers method.
	ListMembersFunc func(ctx context.Context, operatorFilter crew.Filter) ([]crew.Member, int, error)

	// RemoveMemberFunc mocks the RemoveMember method.
	RemoveMemberFunc func(ctx context.Context, key crew.Key) error

	// UpdateMemberFunc mocks the UpdateMember method.
	UpdateMemberFunc func(ctx context.Context, op crew.Member) error

	// calls tracks calls to the methods.
	calls struct {
		// AuthenticateMember holds details about calls to the AuthenticateMember method.
		AuthenticateMember []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Auth is the auth argument value.
			Auth AuthRequest
		}
		// CreateMember holds details about calls to the CreateMember method.
		CreateMember []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Op is the op argument value.
			Op crew.Member
		}
		// GetMember holds details about calls to the GetMember method.
		GetMember []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Key is the key argument value.
			Key crew.Key
		}
		// ListMembers holds details about calls to the ListMembers method.
		ListMembers []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// OperatorFilter is the operatorFilter argument value.
			OperatorFilter crew.Filter
		}
		// RemoveMember holds details about calls to the RemoveMember method.
		RemoveMember []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Key is the key argument value.
			Key crew.Key
		}
		// UpdateMember holds details about calls to the UpdateMember method.
		UpdateMember []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Op is the op argument value.
			Op crew.Member
		}
	}
	lockAuthenticateMember sync.RWMutex
	lockCreateMember       sync.RWMutex
	lockGetMember          sync.RWMutex
	lockListMembers        sync.RWMutex
	lockRemoveMember       sync.RWMutex
	lockUpdateMember       sync.RWMutex
}

// AuthenticateMember calls AuthenticateMemberFunc.
func (mock *WorkflowMock) AuthenticateMember(ctx context.Context, auth AuthRequest) (session.Session, error) {
	if mock.AuthenticateMemberFunc == nil {
		panic("WorkflowMock.AuthenticateMemberFunc: method is nil but Workflow.AuthenticateMember was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Auth AuthRequest
	}{
		Ctx:  ctx,
		Auth: auth,
	}
	mock.lockAuthenticateMember.Lock()
	mock.calls.AuthenticateMember = append(mock.calls.AuthenticateMember, callInfo)
	mock.lockAuthenticateMember.Unlock()
	return mock.AuthenticateMemberFunc(ctx, auth)
}

// AuthenticateMemberCalls gets all the calls that were made to AuthenticateMember.
// Check the length with:
//     len(mockedWorkflow.AuthenticateMemberCalls())
func (mock *WorkflowMock) AuthenticateMemberCalls() []struct {
	Ctx  context.Context
	Auth AuthRequest
} {
	var calls []struct {
		Ctx  context.Context
		Auth AuthRequest
	}
	mock.lockAuthenticateMember.RLock()
	calls = mock.calls.AuthenticateMember
	mock.lockAuthenticateMember.RUnlock()
	return calls
}

// CreateMember calls CreateMemberFunc.
func (mock *WorkflowMock) CreateMember(ctx context.Context, op crew.Member) (crew.Member, error) {
	if mock.CreateMemberFunc == nil {
		panic("WorkflowMock.CreateMemberFunc: method is nil but Workflow.CreateMember was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Op  crew.Member
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
//     len(mockedWorkflow.CreateMemberCalls())
func (mock *WorkflowMock) CreateMemberCalls() []struct {
	Ctx context.Context
	Op  crew.Member
} {
	var calls []struct {
		Ctx context.Context
		Op  crew.Member
	}
	mock.lockCreateMember.RLock()
	calls = mock.calls.CreateMember
	mock.lockCreateMember.RUnlock()
	return calls
}

// GetMember calls GetMemberFunc.
func (mock *WorkflowMock) GetMember(ctx context.Context, key crew.Key) (crew.Member, error) {
	if mock.GetMemberFunc == nil {
		panic("WorkflowMock.GetMemberFunc: method is nil but Workflow.GetMember was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Key crew.Key
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
//     len(mockedWorkflow.GetMemberCalls())
func (mock *WorkflowMock) GetMemberCalls() []struct {
	Ctx context.Context
	Key crew.Key
} {
	var calls []struct {
		Ctx context.Context
		Key crew.Key
	}
	mock.lockGetMember.RLock()
	calls = mock.calls.GetMember
	mock.lockGetMember.RUnlock()
	return calls
}

// ListMembers calls ListMembersFunc.
func (mock *WorkflowMock) ListMembers(ctx context.Context, operatorFilter crew.Filter) ([]crew.Member, int, error) {
	if mock.ListMembersFunc == nil {
		panic("WorkflowMock.ListMembersFunc: method is nil but Workflow.ListMembers was just called")
	}
	callInfo := struct {
		Ctx            context.Context
		OperatorFilter crew.Filter
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
//     len(mockedWorkflow.ListMembersCalls())
func (mock *WorkflowMock) ListMembersCalls() []struct {
	Ctx            context.Context
	OperatorFilter crew.Filter
} {
	var calls []struct {
		Ctx            context.Context
		OperatorFilter crew.Filter
	}
	mock.lockListMembers.RLock()
	calls = mock.calls.ListMembers
	mock.lockListMembers.RUnlock()
	return calls
}

// RemoveMember calls RemoveMemberFunc.
func (mock *WorkflowMock) RemoveMember(ctx context.Context, key crew.Key) error {
	if mock.RemoveMemberFunc == nil {
		panic("WorkflowMock.RemoveMemberFunc: method is nil but Workflow.RemoveMember was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Key crew.Key
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
//     len(mockedWorkflow.RemoveMemberCalls())
func (mock *WorkflowMock) RemoveMemberCalls() []struct {
	Ctx context.Context
	Key crew.Key
} {
	var calls []struct {
		Ctx context.Context
		Key crew.Key
	}
	mock.lockRemoveMember.RLock()
	calls = mock.calls.RemoveMember
	mock.lockRemoveMember.RUnlock()
	return calls
}

// UpdateMember calls UpdateMemberFunc.
func (mock *WorkflowMock) UpdateMember(ctx context.Context, op crew.Member) error {
	if mock.UpdateMemberFunc == nil {
		panic("WorkflowMock.UpdateMemberFunc: method is nil but Workflow.UpdateMember was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Op  crew.Member
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
//     len(mockedWorkflow.UpdateMemberCalls())
func (mock *WorkflowMock) UpdateMemberCalls() []struct {
	Ctx context.Context
	Op  crew.Member
} {
	var calls []struct {
		Ctx context.Context
		Op  crew.Member
	}
	mock.lockUpdateMember.RLock()
	calls = mock.calls.UpdateMember
	mock.lockUpdateMember.RUnlock()
	return calls
}