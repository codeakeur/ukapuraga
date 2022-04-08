// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package customer

import (
	"context"
	"gomies/pkg/sdk/types"
	"sync"
)

// Ensure, that ContactingServiceMock does implement ContactingService.
// If this is not the case, regenerate this file with moq.
var _ ContactingService = &ContactingServiceMock{}

// ContactingServiceMock is a mock implementation of ContactingService.
//
// 	func TestSomethingThatUsesContactingService(t *testing.T) {
//
// 		// make and configure a mocked ContactingService
// 		mockedContactingService := &ContactingServiceMock{
// 			RemoveAllCustomerContactsFunc: func(ctx context.Context, customerUID types.UID) error {
// 				panic("mock out the RemoveAllCustomerContacts method")
// 			},
// 			SaveContactFunc: func(ctx context.Context, addresses []Address, phones []Phone) error {
// 				panic("mock out the SaveContact method")
// 			},
// 		}
//
// 		// use mockedContactingService in code that requires ContactingService
// 		// and then make assertions.
//
// 	}
type ContactingServiceMock struct {
	// RemoveAllCustomerContactsFunc mocks the RemoveAllCustomerContacts method.
	RemoveAllCustomerContactsFunc func(ctx context.Context, customerUID types.UID) error

	// SaveContactFunc mocks the SaveContact method.
	SaveContactFunc func(ctx context.Context, addresses []Address, phones []Phone) error

	// calls tracks calls to the methods.
	calls struct {
		// RemoveAllCustomerContacts holds details about calls to the RemoveAllCustomerContacts method.
		RemoveAllCustomerContacts []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CustomerUID is the customerUID argument value.
			CustomerUID types.UID
		}
		// SaveContact holds details about calls to the SaveContact method.
		SaveContact []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Addresses is the addresses argument value.
			Addresses []Address
			// Phones is the phones argument value.
			Phones []Phone
		}
	}
	lockRemoveAllCustomerContacts sync.RWMutex
	lockSaveContact               sync.RWMutex
}

// RemoveAllCustomerContacts calls RemoveAllCustomerContactsFunc.
func (mock *ContactingServiceMock) RemoveAllCustomerContacts(ctx context.Context, customerUID types.UID) error {
	if mock.RemoveAllCustomerContactsFunc == nil {
		panic("ContactingServiceMock.RemoveAllCustomerContactsFunc: method is nil but ContactingService.RemoveAllCustomerContacts was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		CustomerUID types.UID
	}{
		Ctx:         ctx,
		CustomerUID: customerUID,
	}
	mock.lockRemoveAllCustomerContacts.Lock()
	mock.calls.RemoveAllCustomerContacts = append(mock.calls.RemoveAllCustomerContacts, callInfo)
	mock.lockRemoveAllCustomerContacts.Unlock()
	return mock.RemoveAllCustomerContactsFunc(ctx, customerUID)
}

// RemoveAllCustomerContactsCalls gets all the calls that were made to RemoveAllCustomerContacts.
// Check the length with:
//     len(mockedContactingService.RemoveAllCustomerContactsCalls())
func (mock *ContactingServiceMock) RemoveAllCustomerContactsCalls() []struct {
	Ctx         context.Context
	CustomerUID types.UID
} {
	var calls []struct {
		Ctx         context.Context
		CustomerUID types.UID
	}
	mock.lockRemoveAllCustomerContacts.RLock()
	calls = mock.calls.RemoveAllCustomerContacts
	mock.lockRemoveAllCustomerContacts.RUnlock()
	return calls
}

// SaveContact calls SaveContactFunc.
func (mock *ContactingServiceMock) SaveContact(ctx context.Context, addresses []Address, phones []Phone) error {
	if mock.SaveContactFunc == nil {
		panic("ContactingServiceMock.SaveContactFunc: method is nil but ContactingService.SaveContact was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		Addresses []Address
		Phones    []Phone
	}{
		Ctx:       ctx,
		Addresses: addresses,
		Phones:    phones,
	}
	mock.lockSaveContact.Lock()
	mock.calls.SaveContact = append(mock.calls.SaveContact, callInfo)
	mock.lockSaveContact.Unlock()
	return mock.SaveContactFunc(ctx, addresses, phones)
}

// SaveContactCalls gets all the calls that were made to SaveContact.
// Check the length with:
//     len(mockedContactingService.SaveContactCalls())
func (mock *ContactingServiceMock) SaveContactCalls() []struct {
	Ctx       context.Context
	Addresses []Address
	Phones    []Phone
} {
	var calls []struct {
		Ctx       context.Context
		Addresses []Address
		Phones    []Phone
	}
	mock.lockSaveContact.RLock()
	calls = mock.calls.SaveContact
	mock.lockSaveContact.RUnlock()
	return calls
}