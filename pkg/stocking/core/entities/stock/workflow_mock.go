// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package stock

import (
	"context"
	"gomies/pkg/sdk/types"
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
// 			ClosePeriodFunc: func(ctx context.Context, filter Filter) error {
// 				panic("mock out the ClosePeriod method")
// 			},
// 			ComputeFunc: func(ctx context.Context, filter Filter) (types.Quantity, error) {
// 				panic("mock out the Compute method")
// 			},
// 			ComputeSomeFunc: func(ctx context.Context, filter Filter, resourcesIDs ...types.UID) ([]types.Quantity, error) {
// 				panic("mock out the ComputeSome method")
// 			},
// 			ListMovementsFunc: func(ctx context.Context, filter Filter) ([]Movement, error) {
// 				panic("mock out the ListMovements method")
// 			},
// 			RemoveMovementFunc: func(ctx context.Context, resourceID types.UID, movementID types.UID) error {
// 				panic("mock out the RemoveMovement method")
// 			},
// 			SaveMovementsFunc: func(ctx context.Context, config Config, resourceID types.UID, movements ...Movement) (AdditionResult, error) {
// 				panic("mock out the SaveMovements method")
// 			},
// 		}
//
// 		// use mockedWorkflow in code that requires Workflow
// 		// and then make assertions.
//
// 	}
type WorkflowMock struct {
	// ClosePeriodFunc mocks the ClosePeriod method.
	ClosePeriodFunc func(ctx context.Context, filter Filter) error

	// ComputeFunc mocks the Compute method.
	ComputeFunc func(ctx context.Context, filter Filter) (types.Quantity, error)

	// ComputeSomeFunc mocks the ComputeSome method.
	ComputeSomeFunc func(ctx context.Context, filter Filter, resourcesIDs ...types.UID) ([]types.Quantity, error)

	// ListMovementsFunc mocks the ListMovements method.
	ListMovementsFunc func(ctx context.Context, filter Filter) ([]Movement, error)

	// RemoveMovementFunc mocks the RemoveMovement method.
	RemoveMovementFunc func(ctx context.Context, resourceID types.UID, movementID types.UID) error

	// SaveMovementsFunc mocks the SaveMovements method.
	SaveMovementsFunc func(ctx context.Context, config Config, resourceID types.UID, movements ...Movement) (AdditionResult, error)

	// calls tracks calls to the methods.
	calls struct {
		// ClosePeriod holds details about calls to the ClosePeriod method.
		ClosePeriod []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Filter is the filter argument value.
			Filter Filter
		}
		// Compute holds details about calls to the Compute method.
		Compute []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Filter is the filter argument value.
			Filter Filter
		}
		// ComputeSome holds details about calls to the ComputeSome method.
		ComputeSome []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Filter is the filter argument value.
			Filter Filter
			// ResourcesIDs is the resourcesIDs argument value.
			ResourcesIDs []types.UID
		}
		// ListMovements holds details about calls to the ListMovements method.
		ListMovements []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Filter is the filter argument value.
			Filter Filter
		}
		// RemoveMovement holds details about calls to the RemoveMovement method.
		RemoveMovement []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ResourceID is the resourceID argument value.
			ResourceID types.UID
			// MovementID is the movementID argument value.
			MovementID types.UID
		}
		// SaveMovements holds details about calls to the SaveMovements method.
		SaveMovements []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Config is the config argument value.
			Config Config
			// ResourceID is the resourceID argument value.
			ResourceID types.UID
			// Movements is the movements argument value.
			Movements []Movement
		}
	}
	lockClosePeriod    sync.RWMutex
	lockCompute        sync.RWMutex
	lockComputeSome    sync.RWMutex
	lockListMovements  sync.RWMutex
	lockRemoveMovement sync.RWMutex
	lockSaveMovements  sync.RWMutex
}

// ClosePeriod calls ClosePeriodFunc.
func (mock *WorkflowMock) ClosePeriod(ctx context.Context, filter Filter) error {
	if mock.ClosePeriodFunc == nil {
		panic("WorkflowMock.ClosePeriodFunc: method is nil but Workflow.ClosePeriod was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Filter Filter
	}{
		Ctx:    ctx,
		Filter: filter,
	}
	mock.lockClosePeriod.Lock()
	mock.calls.ClosePeriod = append(mock.calls.ClosePeriod, callInfo)
	mock.lockClosePeriod.Unlock()
	return mock.ClosePeriodFunc(ctx, filter)
}

// ClosePeriodCalls gets all the calls that were made to ClosePeriod.
// Check the length with:
//     len(mockedWorkflow.ClosePeriodCalls())
func (mock *WorkflowMock) ClosePeriodCalls() []struct {
	Ctx    context.Context
	Filter Filter
} {
	var calls []struct {
		Ctx    context.Context
		Filter Filter
	}
	mock.lockClosePeriod.RLock()
	calls = mock.calls.ClosePeriod
	mock.lockClosePeriod.RUnlock()
	return calls
}

// Compute calls ComputeFunc.
func (mock *WorkflowMock) Compute(ctx context.Context, filter Filter) (types.Quantity, error) {
	if mock.ComputeFunc == nil {
		panic("WorkflowMock.ComputeFunc: method is nil but Workflow.Compute was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Filter Filter
	}{
		Ctx:    ctx,
		Filter: filter,
	}
	mock.lockCompute.Lock()
	mock.calls.Compute = append(mock.calls.Compute, callInfo)
	mock.lockCompute.Unlock()
	return mock.ComputeFunc(ctx, filter)
}

// ComputeCalls gets all the calls that were made to Compute.
// Check the length with:
//     len(mockedWorkflow.ComputeCalls())
func (mock *WorkflowMock) ComputeCalls() []struct {
	Ctx    context.Context
	Filter Filter
} {
	var calls []struct {
		Ctx    context.Context
		Filter Filter
	}
	mock.lockCompute.RLock()
	calls = mock.calls.Compute
	mock.lockCompute.RUnlock()
	return calls
}

// ComputeSome calls ComputeSomeFunc.
func (mock *WorkflowMock) ComputeSome(ctx context.Context, filter Filter, resourcesIDs ...types.UID) ([]types.Quantity, error) {
	if mock.ComputeSomeFunc == nil {
		panic("WorkflowMock.ComputeSomeFunc: method is nil but Workflow.ComputeSome was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		Filter       Filter
		ResourcesIDs []types.UID
	}{
		Ctx:          ctx,
		Filter:       filter,
		ResourcesIDs: resourcesIDs,
	}
	mock.lockComputeSome.Lock()
	mock.calls.ComputeSome = append(mock.calls.ComputeSome, callInfo)
	mock.lockComputeSome.Unlock()
	return mock.ComputeSomeFunc(ctx, filter, resourcesIDs...)
}

// ComputeSomeCalls gets all the calls that were made to ComputeSome.
// Check the length with:
//     len(mockedWorkflow.ComputeSomeCalls())
func (mock *WorkflowMock) ComputeSomeCalls() []struct {
	Ctx          context.Context
	Filter       Filter
	ResourcesIDs []types.UID
} {
	var calls []struct {
		Ctx          context.Context
		Filter       Filter
		ResourcesIDs []types.UID
	}
	mock.lockComputeSome.RLock()
	calls = mock.calls.ComputeSome
	mock.lockComputeSome.RUnlock()
	return calls
}

// ListMovements calls ListMovementsFunc.
func (mock *WorkflowMock) ListMovements(ctx context.Context, filter Filter) ([]Movement, error) {
	if mock.ListMovementsFunc == nil {
		panic("WorkflowMock.ListMovementsFunc: method is nil but Workflow.ListMovements was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		Filter Filter
	}{
		Ctx:    ctx,
		Filter: filter,
	}
	mock.lockListMovements.Lock()
	mock.calls.ListMovements = append(mock.calls.ListMovements, callInfo)
	mock.lockListMovements.Unlock()
	return mock.ListMovementsFunc(ctx, filter)
}

// ListMovementsCalls gets all the calls that were made to ListMovements.
// Check the length with:
//     len(mockedWorkflow.ListMovementsCalls())
func (mock *WorkflowMock) ListMovementsCalls() []struct {
	Ctx    context.Context
	Filter Filter
} {
	var calls []struct {
		Ctx    context.Context
		Filter Filter
	}
	mock.lockListMovements.RLock()
	calls = mock.calls.ListMovements
	mock.lockListMovements.RUnlock()
	return calls
}

// RemoveMovement calls RemoveMovementFunc.
func (mock *WorkflowMock) RemoveMovement(ctx context.Context, resourceID types.UID, movementID types.UID) error {
	if mock.RemoveMovementFunc == nil {
		panic("WorkflowMock.RemoveMovementFunc: method is nil but Workflow.RemoveMovement was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		ResourceID types.UID
		MovementID types.UID
	}{
		Ctx:        ctx,
		ResourceID: resourceID,
		MovementID: movementID,
	}
	mock.lockRemoveMovement.Lock()
	mock.calls.RemoveMovement = append(mock.calls.RemoveMovement, callInfo)
	mock.lockRemoveMovement.Unlock()
	return mock.RemoveMovementFunc(ctx, resourceID, movementID)
}

// RemoveMovementCalls gets all the calls that were made to RemoveMovement.
// Check the length with:
//     len(mockedWorkflow.RemoveMovementCalls())
func (mock *WorkflowMock) RemoveMovementCalls() []struct {
	Ctx        context.Context
	ResourceID types.UID
	MovementID types.UID
} {
	var calls []struct {
		Ctx        context.Context
		ResourceID types.UID
		MovementID types.UID
	}
	mock.lockRemoveMovement.RLock()
	calls = mock.calls.RemoveMovement
	mock.lockRemoveMovement.RUnlock()
	return calls
}

// SaveMovements calls SaveMovementsFunc.
func (mock *WorkflowMock) SaveMovements(ctx context.Context, config Config, resourceID types.UID, movements ...Movement) (AdditionResult, error) {
	if mock.SaveMovementsFunc == nil {
		panic("WorkflowMock.SaveMovementsFunc: method is nil but Workflow.SaveMovements was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		Config     Config
		ResourceID types.UID
		Movements  []Movement
	}{
		Ctx:        ctx,
		Config:     config,
		ResourceID: resourceID,
		Movements:  movements,
	}
	mock.lockSaveMovements.Lock()
	mock.calls.SaveMovements = append(mock.calls.SaveMovements, callInfo)
	mock.lockSaveMovements.Unlock()
	return mock.SaveMovementsFunc(ctx, config, resourceID, movements...)
}

// SaveMovementsCalls gets all the calls that were made to SaveMovements.
// Check the length with:
//     len(mockedWorkflow.SaveMovementsCalls())
func (mock *WorkflowMock) SaveMovementsCalls() []struct {
	Ctx        context.Context
	Config     Config
	ResourceID types.UID
	Movements  []Movement
} {
	var calls []struct {
		Ctx        context.Context
		Config     Config
		ResourceID types.UID
		Movements  []Movement
	}
	mock.lockSaveMovements.RLock()
	calls = mock.calls.SaveMovements
	mock.lockSaveMovements.RUnlock()
	return calls
}
