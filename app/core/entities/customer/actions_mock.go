// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package customer

import (
	"context"
	"gomies/pkg/sdk/types"
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
// 			CreateCustomerFunc: func(ctx context.Context, c Customer) (Customer, error) {
// 				panic("mock out the CreateCustomer method")
// 			},
// 			GetCustomerFunc: func(ctx context.Context, uid types.UID) (Customer, error) {
// 				panic("mock out the GetCustomer method")
// 			},
// 			ListCustomersFunc: func(ctx context.Context, f Filter) ([]Customer, int, error) {
// 				panic("mock out the ListCustomers method")
// 			},
// 			RemoveCustomerFunc: func(ctx context.Context, uid types.UID) error {
// 				panic("mock out the RemoveCustomer method")
// 			},
// 			UpdateCustomerFunc: func(ctx context.Context, c Customer) error {
// 				panic("mock out the UpdateCustomer method")
// 			},
// 		}
//
// 		// use mockedActions in code that requires Actions
// 		// and then make assertions.
//
// 	}
type ActionsMock struct {
	// CreateCustomerFunc mocks the CreateCustomer method.
	CreateCustomerFunc func(ctx context.Context, c Customer) (Customer, error)

	// GetCustomerFunc mocks the GetCustomer method.
	GetCustomerFunc func(ctx context.Context, uid types.UID) (Customer, error)

	// ListCustomersFunc mocks the ListCustomers method.
	ListCustomersFunc func(ctx context.Context, f Filter) ([]Customer, int, error)

	// RemoveCustomerFunc mocks the RemoveCustomer method.
	RemoveCustomerFunc func(ctx context.Context, uid types.UID) error

	// UpdateCustomerFunc mocks the UpdateCustomer method.
	UpdateCustomerFunc func(ctx context.Context, c Customer) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateCustomer holds details about calls to the CreateCustomer method.
		CreateCustomer []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// C is the c argument value.
			C Customer
		}
		// GetCustomer holds details about calls to the GetCustomer method.
		GetCustomer []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UID is the uid argument value.
			UID types.UID
		}
		// ListCustomers holds details about calls to the ListCustomers method.
		ListCustomers []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// F is the f argument value.
			F Filter
		}
		// RemoveCustomer holds details about calls to the RemoveCustomer method.
		RemoveCustomer []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// UID is the uid argument value.
			UID types.UID
		}
		// UpdateCustomer holds details about calls to the UpdateCustomer method.
		UpdateCustomer []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// C is the c argument value.
			C Customer
		}
	}
	lockCreateCustomer sync.RWMutex
	lockGetCustomer    sync.RWMutex
	lockListCustomers  sync.RWMutex
	lockRemoveCustomer sync.RWMutex
	lockUpdateCustomer sync.RWMutex
}

// CreateCustomer calls CreateCustomerFunc.
func (mock *ActionsMock) CreateCustomer(ctx context.Context, c Customer) (Customer, error) {
	if mock.CreateCustomerFunc == nil {
		panic("ActionsMock.CreateCustomerFunc: method is nil but Actions.CreateCustomer was just called")
	}
	callInfo := struct {
		Ctx context.Context
		C   Customer
	}{
		Ctx: ctx,
		C:   c,
	}
	mock.lockCreateCustomer.Lock()
	mock.calls.CreateCustomer = append(mock.calls.CreateCustomer, callInfo)
	mock.lockCreateCustomer.Unlock()
	return mock.CreateCustomerFunc(ctx, c)
}

// CreateCustomerCalls gets all the calls that were made to CreateCustomer.
// Check the length with:
//     len(mockedActions.CreateCustomerCalls())
func (mock *ActionsMock) CreateCustomerCalls() []struct {
	Ctx context.Context
	C   Customer
} {
	var calls []struct {
		Ctx context.Context
		C   Customer
	}
	mock.lockCreateCustomer.RLock()
	calls = mock.calls.CreateCustomer
	mock.lockCreateCustomer.RUnlock()
	return calls
}

// GetCustomer calls GetCustomerFunc.
func (mock *ActionsMock) GetCustomer(ctx context.Context, uid types.UID) (Customer, error) {
	if mock.GetCustomerFunc == nil {
		panic("ActionsMock.GetCustomerFunc: method is nil but Actions.GetCustomer was just called")
	}
	callInfo := struct {
		Ctx context.Context
		UID types.UID
	}{
		Ctx: ctx,
		UID: uid,
	}
	mock.lockGetCustomer.Lock()
	mock.calls.GetCustomer = append(mock.calls.GetCustomer, callInfo)
	mock.lockGetCustomer.Unlock()
	return mock.GetCustomerFunc(ctx, uid)
}

// GetCustomerCalls gets all the calls that were made to GetCustomer.
// Check the length with:
//     len(mockedActions.GetCustomerCalls())
func (mock *ActionsMock) GetCustomerCalls() []struct {
	Ctx context.Context
	UID types.UID
} {
	var calls []struct {
		Ctx context.Context
		UID types.UID
	}
	mock.lockGetCustomer.RLock()
	calls = mock.calls.GetCustomer
	mock.lockGetCustomer.RUnlock()
	return calls
}

// ListCustomers calls ListCustomersFunc.
func (mock *ActionsMock) ListCustomers(ctx context.Context, f Filter) ([]Customer, int, error) {
	if mock.ListCustomersFunc == nil {
		panic("ActionsMock.ListCustomersFunc: method is nil but Actions.ListCustomers was just called")
	}
	callInfo := struct {
		Ctx context.Context
		F   Filter
	}{
		Ctx: ctx,
		F:   f,
	}
	mock.lockListCustomers.Lock()
	mock.calls.ListCustomers = append(mock.calls.ListCustomers, callInfo)
	mock.lockListCustomers.Unlock()
	return mock.ListCustomersFunc(ctx, f)
}

// ListCustomersCalls gets all the calls that were made to ListCustomers.
// Check the length with:
//     len(mockedActions.ListCustomersCalls())
func (mock *ActionsMock) ListCustomersCalls() []struct {
	Ctx context.Context
	F   Filter
} {
	var calls []struct {
		Ctx context.Context
		F   Filter
	}
	mock.lockListCustomers.RLock()
	calls = mock.calls.ListCustomers
	mock.lockListCustomers.RUnlock()
	return calls
}

// RemoveCustomer calls RemoveCustomerFunc.
func (mock *ActionsMock) RemoveCustomer(ctx context.Context, uid types.UID) error {
	if mock.RemoveCustomerFunc == nil {
		panic("ActionsMock.RemoveCustomerFunc: method is nil but Actions.RemoveCustomer was just called")
	}
	callInfo := struct {
		Ctx context.Context
		UID types.UID
	}{
		Ctx: ctx,
		UID: uid,
	}
	mock.lockRemoveCustomer.Lock()
	mock.calls.RemoveCustomer = append(mock.calls.RemoveCustomer, callInfo)
	mock.lockRemoveCustomer.Unlock()
	return mock.RemoveCustomerFunc(ctx, uid)
}

// RemoveCustomerCalls gets all the calls that were made to RemoveCustomer.
// Check the length with:
//     len(mockedActions.RemoveCustomerCalls())
func (mock *ActionsMock) RemoveCustomerCalls() []struct {
	Ctx context.Context
	UID types.UID
} {
	var calls []struct {
		Ctx context.Context
		UID types.UID
	}
	mock.lockRemoveCustomer.RLock()
	calls = mock.calls.RemoveCustomer
	mock.lockRemoveCustomer.RUnlock()
	return calls
}

// UpdateCustomer calls UpdateCustomerFunc.
func (mock *ActionsMock) UpdateCustomer(ctx context.Context, c Customer) error {
	if mock.UpdateCustomerFunc == nil {
		panic("ActionsMock.UpdateCustomerFunc: method is nil but Actions.UpdateCustomer was just called")
	}
	callInfo := struct {
		Ctx context.Context
		C   Customer
	}{
		Ctx: ctx,
		C:   c,
	}
	mock.lockUpdateCustomer.Lock()
	mock.calls.UpdateCustomer = append(mock.calls.UpdateCustomer, callInfo)
	mock.lockUpdateCustomer.Unlock()
	return mock.UpdateCustomerFunc(ctx, c)
}

// UpdateCustomerCalls gets all the calls that were made to UpdateCustomer.
// Check the length with:
//     len(mockedActions.UpdateCustomerCalls())
func (mock *ActionsMock) UpdateCustomerCalls() []struct {
	Ctx context.Context
	C   Customer
} {
	var calls []struct {
		Ctx context.Context
		C   Customer
	}
	mock.lockUpdateCustomer.RLock()
	calls = mock.calls.UpdateCustomer
	mock.lockUpdateCustomer.RUnlock()
	return calls
}
