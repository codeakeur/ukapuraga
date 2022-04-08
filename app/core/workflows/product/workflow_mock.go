// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package product

import (
	"context"
	"gomies/app/core/entities/catalog/product"
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
// 			ApproveSaleFunc: func(ctx context.Context, req product.ApproveSaleRequest) error {
// 				panic("mock out the ApproveSale method")
// 			},
// 			CreateIngredientFunc: func(ctx context.Context, productKey product.Key, ingredient product.Ingredient) (product.Ingredient, error) {
// 				panic("mock out the CreateIngredient method")
// 			},
// 			CreateMovementFunc: func(ctx context.Context, productID types.UID, mov Movement) (types.Quantity, error) {
// 				panic("mock out the CreateMovement method")
// 			},
// 			CreateProductFunc: func(ctx context.Context, prd product.Product) (product.Product, error) {
// 				panic("mock out the CreateProduct method")
// 			},
// 			GetProductFunc: func(ctx context.Context, key product.Key) (product.Product, error) {
// 				panic("mock out the GetProduct method")
// 			},
// 			ListProductsFunc: func(ctx context.Context, productFilter product.Filter) ([]product.Product, int, error) {
// 				panic("mock out the ListProducts method")
// 			},
// 			RemoveIngredientFunc: func(ctx context.Context, productKey product.Key, id types.UID) error {
// 				panic("mock out the RemoveIngredient method")
// 			},
// 			RemoveProductFunc: func(ctx context.Context, key product.Key) error {
// 				panic("mock out the RemoveProduct method")
// 			},
// 			UpdateProductFunc: func(ctx context.Context, prd product.Product) error {
// 				panic("mock out the UpdateProduct method")
// 			},
// 		}
//
// 		// use mockedWorkflow in code that requires Workflow
// 		// and then make assertions.
//
// 	}
type WorkflowMock struct {
	// ApproveSaleFunc mocks the ApproveSale method.
	ApproveSaleFunc func(ctx context.Context, req product.ApproveSaleRequest) error

	// CreateIngredientFunc mocks the CreateIngredient method.
	CreateIngredientFunc func(ctx context.Context, productKey product.Key, ingredient product.Ingredient) (product.Ingredient, error)

	// CreateMovementFunc mocks the CreateMovement method.
	CreateMovementFunc func(ctx context.Context, productID types.UID, mov Movement) (types.Quantity, error)

	// CreateProductFunc mocks the CreateProduct method.
	CreateProductFunc func(ctx context.Context, prd product.Product) (product.Product, error)

	// GetProductFunc mocks the GetProduct method.
	GetProductFunc func(ctx context.Context, key product.Key) (product.Product, error)

	// ListProductsFunc mocks the ListProducts method.
	ListProductsFunc func(ctx context.Context, productFilter product.Filter) ([]product.Product, int, error)

	// RemoveIngredientFunc mocks the RemoveIngredient method.
	RemoveIngredientFunc func(ctx context.Context, productKey product.Key, id types.UID) error

	// RemoveProductFunc mocks the RemoveProduct method.
	RemoveProductFunc func(ctx context.Context, key product.Key) error

	// UpdateProductFunc mocks the UpdateProduct method.
	UpdateProductFunc func(ctx context.Context, prd product.Product) error

	// calls tracks calls to the methods.
	calls struct {
		// ApproveSale holds details about calls to the ApproveSale method.
		ApproveSale []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Req is the req argument value.
			Req product.ApproveSaleRequest
		}
		// CreateIngredient holds details about calls to the CreateIngredient method.
		CreateIngredient []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ProductKey is the productKey argument value.
			ProductKey product.Key
			// Ingredient is the ingredient argument value.
			Ingredient product.Ingredient
		}
		// CreateMovement holds details about calls to the CreateMovement method.
		CreateMovement []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ProductID is the productID argument value.
			ProductID types.UID
			// Mov is the mov argument value.
			Mov Movement
		}
		// CreateProduct holds details about calls to the CreateProduct method.
		CreateProduct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Prd is the prd argument value.
			Prd product.Product
		}
		// GetProduct holds details about calls to the GetProduct method.
		GetProduct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Key is the key argument value.
			Key product.Key
		}
		// ListProducts holds details about calls to the ListProducts method.
		ListProducts []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ProductFilter is the productFilter argument value.
			ProductFilter product.Filter
		}
		// RemoveIngredient holds details about calls to the RemoveIngredient method.
		RemoveIngredient []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ProductKey is the productKey argument value.
			ProductKey product.Key
			// ID is the id argument value.
			ID types.UID
		}
		// RemoveProduct holds details about calls to the RemoveProduct method.
		RemoveProduct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Key is the key argument value.
			Key product.Key
		}
		// UpdateProduct holds details about calls to the UpdateProduct method.
		UpdateProduct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Prd is the prd argument value.
			Prd product.Product
		}
	}
	lockApproveSale      sync.RWMutex
	lockCreateIngredient sync.RWMutex
	lockCreateMovement   sync.RWMutex
	lockCreateProduct    sync.RWMutex
	lockGetProduct       sync.RWMutex
	lockListProducts     sync.RWMutex
	lockRemoveIngredient sync.RWMutex
	lockRemoveProduct    sync.RWMutex
	lockUpdateProduct    sync.RWMutex
}

// ApproveSale calls ApproveSaleFunc.
func (mock *WorkflowMock) ApproveSale(ctx context.Context, req product.ApproveSaleRequest) error {
	if mock.ApproveSaleFunc == nil {
		panic("WorkflowMock.ApproveSaleFunc: method is nil but Workflow.ApproveSale was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Req product.ApproveSaleRequest
	}{
		Ctx: ctx,
		Req: req,
	}
	mock.lockApproveSale.Lock()
	mock.calls.ApproveSale = append(mock.calls.ApproveSale, callInfo)
	mock.lockApproveSale.Unlock()
	return mock.ApproveSaleFunc(ctx, req)
}

// ApproveSaleCalls gets all the calls that were made to ApproveSale.
// Check the length with:
//     len(mockedWorkflow.ApproveSaleCalls())
func (mock *WorkflowMock) ApproveSaleCalls() []struct {
	Ctx context.Context
	Req product.ApproveSaleRequest
} {
	var calls []struct {
		Ctx context.Context
		Req product.ApproveSaleRequest
	}
	mock.lockApproveSale.RLock()
	calls = mock.calls.ApproveSale
	mock.lockApproveSale.RUnlock()
	return calls
}

// CreateIngredient calls CreateIngredientFunc.
func (mock *WorkflowMock) CreateIngredient(ctx context.Context, productKey product.Key, ingredient product.Ingredient) (product.Ingredient, error) {
	if mock.CreateIngredientFunc == nil {
		panic("WorkflowMock.CreateIngredientFunc: method is nil but Workflow.CreateIngredient was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		ProductKey product.Key
		Ingredient product.Ingredient
	}{
		Ctx:        ctx,
		ProductKey: productKey,
		Ingredient: ingredient,
	}
	mock.lockCreateIngredient.Lock()
	mock.calls.CreateIngredient = append(mock.calls.CreateIngredient, callInfo)
	mock.lockCreateIngredient.Unlock()
	return mock.CreateIngredientFunc(ctx, productKey, ingredient)
}

// CreateIngredientCalls gets all the calls that were made to CreateIngredient.
// Check the length with:
//     len(mockedWorkflow.CreateIngredientCalls())
func (mock *WorkflowMock) CreateIngredientCalls() []struct {
	Ctx        context.Context
	ProductKey product.Key
	Ingredient product.Ingredient
} {
	var calls []struct {
		Ctx        context.Context
		ProductKey product.Key
		Ingredient product.Ingredient
	}
	mock.lockCreateIngredient.RLock()
	calls = mock.calls.CreateIngredient
	mock.lockCreateIngredient.RUnlock()
	return calls
}

// CreateMovement calls CreateMovementFunc.
func (mock *WorkflowMock) CreateMovement(ctx context.Context, productID types.UID, mov Movement) (types.Quantity, error) {
	if mock.CreateMovementFunc == nil {
		panic("WorkflowMock.CreateMovementFunc: method is nil but Workflow.CreateMovement was just called")
	}
	callInfo := struct {
		Ctx       context.Context
		ProductID types.UID
		Mov       Movement
	}{
		Ctx:       ctx,
		ProductID: productID,
		Mov:       mov,
	}
	mock.lockCreateMovement.Lock()
	mock.calls.CreateMovement = append(mock.calls.CreateMovement, callInfo)
	mock.lockCreateMovement.Unlock()
	return mock.CreateMovementFunc(ctx, productID, mov)
}

// CreateMovementCalls gets all the calls that were made to CreateMovement.
// Check the length with:
//     len(mockedWorkflow.CreateMovementCalls())
func (mock *WorkflowMock) CreateMovementCalls() []struct {
	Ctx       context.Context
	ProductID types.UID
	Mov       Movement
} {
	var calls []struct {
		Ctx       context.Context
		ProductID types.UID
		Mov       Movement
	}
	mock.lockCreateMovement.RLock()
	calls = mock.calls.CreateMovement
	mock.lockCreateMovement.RUnlock()
	return calls
}

// CreateProduct calls CreateProductFunc.
func (mock *WorkflowMock) CreateProduct(ctx context.Context, prd product.Product) (product.Product, error) {
	if mock.CreateProductFunc == nil {
		panic("WorkflowMock.CreateProductFunc: method is nil but Workflow.CreateProduct was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Prd product.Product
	}{
		Ctx: ctx,
		Prd: prd,
	}
	mock.lockCreateProduct.Lock()
	mock.calls.CreateProduct = append(mock.calls.CreateProduct, callInfo)
	mock.lockCreateProduct.Unlock()
	return mock.CreateProductFunc(ctx, prd)
}

// CreateProductCalls gets all the calls that were made to CreateProduct.
// Check the length with:
//     len(mockedWorkflow.CreateProductCalls())
func (mock *WorkflowMock) CreateProductCalls() []struct {
	Ctx context.Context
	Prd product.Product
} {
	var calls []struct {
		Ctx context.Context
		Prd product.Product
	}
	mock.lockCreateProduct.RLock()
	calls = mock.calls.CreateProduct
	mock.lockCreateProduct.RUnlock()
	return calls
}

// GetProduct calls GetProductFunc.
func (mock *WorkflowMock) GetProduct(ctx context.Context, key product.Key) (product.Product, error) {
	if mock.GetProductFunc == nil {
		panic("WorkflowMock.GetProductFunc: method is nil but Workflow.GetProduct was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Key product.Key
	}{
		Ctx: ctx,
		Key: key,
	}
	mock.lockGetProduct.Lock()
	mock.calls.GetProduct = append(mock.calls.GetProduct, callInfo)
	mock.lockGetProduct.Unlock()
	return mock.GetProductFunc(ctx, key)
}

// GetProductCalls gets all the calls that were made to GetProduct.
// Check the length with:
//     len(mockedWorkflow.GetProductCalls())
func (mock *WorkflowMock) GetProductCalls() []struct {
	Ctx context.Context
	Key product.Key
} {
	var calls []struct {
		Ctx context.Context
		Key product.Key
	}
	mock.lockGetProduct.RLock()
	calls = mock.calls.GetProduct
	mock.lockGetProduct.RUnlock()
	return calls
}

// ListProducts calls ListProductsFunc.
func (mock *WorkflowMock) ListProducts(ctx context.Context, productFilter product.Filter) ([]product.Product, int, error) {
	if mock.ListProductsFunc == nil {
		panic("WorkflowMock.ListProductsFunc: method is nil but Workflow.ListProducts was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		ProductFilter product.Filter
	}{
		Ctx:           ctx,
		ProductFilter: productFilter,
	}
	mock.lockListProducts.Lock()
	mock.calls.ListProducts = append(mock.calls.ListProducts, callInfo)
	mock.lockListProducts.Unlock()
	return mock.ListProductsFunc(ctx, productFilter)
}

// ListProductsCalls gets all the calls that were made to ListProducts.
// Check the length with:
//     len(mockedWorkflow.ListProductsCalls())
func (mock *WorkflowMock) ListProductsCalls() []struct {
	Ctx           context.Context
	ProductFilter product.Filter
} {
	var calls []struct {
		Ctx           context.Context
		ProductFilter product.Filter
	}
	mock.lockListProducts.RLock()
	calls = mock.calls.ListProducts
	mock.lockListProducts.RUnlock()
	return calls
}

// RemoveIngredient calls RemoveIngredientFunc.
func (mock *WorkflowMock) RemoveIngredient(ctx context.Context, productKey product.Key, id types.UID) error {
	if mock.RemoveIngredientFunc == nil {
		panic("WorkflowMock.RemoveIngredientFunc: method is nil but Workflow.RemoveIngredient was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		ProductKey product.Key
		ID         types.UID
	}{
		Ctx:        ctx,
		ProductKey: productKey,
		ID:         id,
	}
	mock.lockRemoveIngredient.Lock()
	mock.calls.RemoveIngredient = append(mock.calls.RemoveIngredient, callInfo)
	mock.lockRemoveIngredient.Unlock()
	return mock.RemoveIngredientFunc(ctx, productKey, id)
}

// RemoveIngredientCalls gets all the calls that were made to RemoveIngredient.
// Check the length with:
//     len(mockedWorkflow.RemoveIngredientCalls())
func (mock *WorkflowMock) RemoveIngredientCalls() []struct {
	Ctx        context.Context
	ProductKey product.Key
	ID         types.UID
} {
	var calls []struct {
		Ctx        context.Context
		ProductKey product.Key
		ID         types.UID
	}
	mock.lockRemoveIngredient.RLock()
	calls = mock.calls.RemoveIngredient
	mock.lockRemoveIngredient.RUnlock()
	return calls
}

// RemoveProduct calls RemoveProductFunc.
func (mock *WorkflowMock) RemoveProduct(ctx context.Context, key product.Key) error {
	if mock.RemoveProductFunc == nil {
		panic("WorkflowMock.RemoveProductFunc: method is nil but Workflow.RemoveProduct was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Key product.Key
	}{
		Ctx: ctx,
		Key: key,
	}
	mock.lockRemoveProduct.Lock()
	mock.calls.RemoveProduct = append(mock.calls.RemoveProduct, callInfo)
	mock.lockRemoveProduct.Unlock()
	return mock.RemoveProductFunc(ctx, key)
}

// RemoveProductCalls gets all the calls that were made to RemoveProduct.
// Check the length with:
//     len(mockedWorkflow.RemoveProductCalls())
func (mock *WorkflowMock) RemoveProductCalls() []struct {
	Ctx context.Context
	Key product.Key
} {
	var calls []struct {
		Ctx context.Context
		Key product.Key
	}
	mock.lockRemoveProduct.RLock()
	calls = mock.calls.RemoveProduct
	mock.lockRemoveProduct.RUnlock()
	return calls
}

// UpdateProduct calls UpdateProductFunc.
func (mock *WorkflowMock) UpdateProduct(ctx context.Context, prd product.Product) error {
	if mock.UpdateProductFunc == nil {
		panic("WorkflowMock.UpdateProductFunc: method is nil but Workflow.UpdateProduct was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Prd product.Product
	}{
		Ctx: ctx,
		Prd: prd,
	}
	mock.lockUpdateProduct.Lock()
	mock.calls.UpdateProduct = append(mock.calls.UpdateProduct, callInfo)
	mock.lockUpdateProduct.Unlock()
	return mock.UpdateProductFunc(ctx, prd)
}

// UpdateProductCalls gets all the calls that were made to UpdateProduct.
// Check the length with:
//     len(mockedWorkflow.UpdateProductCalls())
func (mock *WorkflowMock) UpdateProductCalls() []struct {
	Ctx context.Context
	Prd product.Product
} {
	var calls []struct {
		Ctx context.Context
		Prd product.Product
	}
	mock.lockUpdateProduct.RLock()
	calls = mock.calls.UpdateProduct
	mock.lockUpdateProduct.RUnlock()
	return calls
}
