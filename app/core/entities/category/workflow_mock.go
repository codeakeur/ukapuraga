// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package category

import (
	"context"
	"gomies/app/core/types/id"
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
// 			CreateFunc: func(contextMoqParam context.Context, category Category) (Category, error) {
// 				panic("mock out the Create method")
// 			},
// 			GetFunc: func(contextMoqParam context.Context, external id.External) (Category, error) {
// 				panic("mock out the Get method")
// 			},
// 			ListFunc: func(contextMoqParam context.Context, filter Filter) ([]Category, error) {
// 				panic("mock out the List method")
// 			},
// 			RemoveFunc: func(contextMoqParam context.Context, external id.External) error {
// 				panic("mock out the Remove method")
// 			},
// 			UpdateFunc: func(contextMoqParam context.Context, category Category) error {
// 				panic("mock out the Update method")
// 			},
// 		}
//
// 		// use mockedWorkflow in code that requires Workflow
// 		// and then make assertions.
//
// 	}
type WorkflowMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(contextMoqParam context.Context, category Category) (Category, error)

	// GetFunc mocks the Get method.
	GetFunc func(contextMoqParam context.Context, external id.External) (Category, error)

	// ListFunc mocks the List method.
	ListFunc func(contextMoqParam context.Context, filter Filter) ([]Category, error)

	// RemoveFunc mocks the Remove method.
	RemoveFunc func(contextMoqParam context.Context, external id.External) error

	// UpdateFunc mocks the Update method.
	UpdateFunc func(contextMoqParam context.Context, category Category) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Category is the category argument value.
			Category Category
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// External is the external argument value.
			External id.External
		}
		// List holds details about calls to the List method.
		List []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Filter is the filter argument value.
			Filter Filter
		}
		// Remove holds details about calls to the Remove method.
		Remove []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// External is the external argument value.
			External id.External
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// ContextMoqParam is the contextMoqParam argument value.
			ContextMoqParam context.Context
			// Category is the category argument value.
			Category Category
		}
	}
	lockCreate sync.RWMutex
	lockGet    sync.RWMutex
	lockList   sync.RWMutex
	lockRemove sync.RWMutex
	lockUpdate sync.RWMutex
}

// Create calls CreateFunc.
func (mock *WorkflowMock) Create(contextMoqParam context.Context, category Category) (Category, error) {
	if mock.CreateFunc == nil {
		panic("WorkflowMock.CreateFunc: method is nil but Workflow.Create was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Category        Category
	}{
		ContextMoqParam: contextMoqParam,
		Category:        category,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(contextMoqParam, category)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//     len(mockedWorkflow.CreateCalls())
func (mock *WorkflowMock) CreateCalls() []struct {
	ContextMoqParam context.Context
	Category        Category
} {
	var calls []struct {
		ContextMoqParam context.Context
		Category        Category
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *WorkflowMock) Get(contextMoqParam context.Context, external id.External) (Category, error) {
	if mock.GetFunc == nil {
		panic("WorkflowMock.GetFunc: method is nil but Workflow.Get was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		External        id.External
	}{
		ContextMoqParam: contextMoqParam,
		External:        external,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(contextMoqParam, external)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedWorkflow.GetCalls())
func (mock *WorkflowMock) GetCalls() []struct {
	ContextMoqParam context.Context
	External        id.External
} {
	var calls []struct {
		ContextMoqParam context.Context
		External        id.External
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// List calls ListFunc.
func (mock *WorkflowMock) List(contextMoqParam context.Context, filter Filter) ([]Category, error) {
	if mock.ListFunc == nil {
		panic("WorkflowMock.ListFunc: method is nil but Workflow.List was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Filter          Filter
	}{
		ContextMoqParam: contextMoqParam,
		Filter:          filter,
	}
	mock.lockList.Lock()
	mock.calls.List = append(mock.calls.List, callInfo)
	mock.lockList.Unlock()
	return mock.ListFunc(contextMoqParam, filter)
}

// ListCalls gets all the calls that were made to List.
// Check the length with:
//     len(mockedWorkflow.ListCalls())
func (mock *WorkflowMock) ListCalls() []struct {
	ContextMoqParam context.Context
	Filter          Filter
} {
	var calls []struct {
		ContextMoqParam context.Context
		Filter          Filter
	}
	mock.lockList.RLock()
	calls = mock.calls.List
	mock.lockList.RUnlock()
	return calls
}

// Remove calls RemoveFunc.
func (mock *WorkflowMock) Remove(contextMoqParam context.Context, external id.External) error {
	if mock.RemoveFunc == nil {
		panic("WorkflowMock.RemoveFunc: method is nil but Workflow.Remove was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		External        id.External
	}{
		ContextMoqParam: contextMoqParam,
		External:        external,
	}
	mock.lockRemove.Lock()
	mock.calls.Remove = append(mock.calls.Remove, callInfo)
	mock.lockRemove.Unlock()
	return mock.RemoveFunc(contextMoqParam, external)
}

// RemoveCalls gets all the calls that were made to Remove.
// Check the length with:
//     len(mockedWorkflow.RemoveCalls())
func (mock *WorkflowMock) RemoveCalls() []struct {
	ContextMoqParam context.Context
	External        id.External
} {
	var calls []struct {
		ContextMoqParam context.Context
		External        id.External
	}
	mock.lockRemove.RLock()
	calls = mock.calls.Remove
	mock.lockRemove.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *WorkflowMock) Update(contextMoqParam context.Context, category Category) error {
	if mock.UpdateFunc == nil {
		panic("WorkflowMock.UpdateFunc: method is nil but Workflow.Update was just called")
	}
	callInfo := struct {
		ContextMoqParam context.Context
		Category        Category
	}{
		ContextMoqParam: contextMoqParam,
		Category:        category,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(contextMoqParam, category)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//     len(mockedWorkflow.UpdateCalls())
func (mock *WorkflowMock) UpdateCalls() []struct {
	ContextMoqParam context.Context
	Category        Category
} {
	var calls []struct {
		ContextMoqParam context.Context
		Category        Category
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}
