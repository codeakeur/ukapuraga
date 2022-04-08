// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package product

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
// 			CreateProductFunc: func(ctx context.Context, prd Product) (Product, error) {
// 				panic("mock out the CreateProduct method")
// 			},
// 			GetProductSaleInfoFunc: func(ctx context.Context, key Key) (Sale, error) {
// 				panic("mock out the GetProductSaleInfo method")
// 			},
// 			GetProductStockInfoFunc: func(ctx context.Context, key Key) (Stock, error) {
// 				panic("mock out the GetProductStockInfo method")
// 			},
// 			GetProductsFunc: func(ctx context.Context, key Key) (Product, error) {
// 				panic("mock out the GetProducts method")
// 			},
// 			ListIngredientsFunc: func(ctx context.Context, productKey Key) ([]Ingredient, error) {
// 				panic("mock out the ListIngredients method")
// 			},
// 			ListProductsFunc: func(ctx context.Context, productFilter Filter) ([]Product, int, error) {
// 				panic("mock out the ListProducts method")
// 			},
// 			RemoveAllIngredientsFunc: func(ctx context.Context, productKey Key) error {
// 				panic("mock out the RemoveAllIngredients method")
// 			},
// 			RemoveIngredientFunc: func(ctx context.Context, productKey Key, ingredientID types.UID) error {
// 				panic("mock out the RemoveIngredient method")
// 			},
// 			RemoveProductFunc: func(ctx context.Context, key Key) error {
// 				panic("mock out the RemoveProduct method")
// 			},
// 			SaveIngredientsFunc: func(ctx context.Context, productKey Key, ingredients ...Ingredient) ([]Ingredient, error) {
// 				panic("mock out the SaveIngredients method")
// 			},
// 			UpdateProductFunc: func(ctx context.Context, prd Product) error {
// 				panic("mock out the UpdateProduct method")
// 			},
// 		}
//
// 		// use mockedActions in code that requires Actions
// 		// and then make assertions.
//
// 	}
type ActionsMock struct {
	// CreateProductFunc mocks the CreateProduct method.
	CreateProductFunc func(ctx context.Context, prd Product) (Product, error)

	// GetProductSaleInfoFunc mocks the GetProductSaleInfo method.
	GetProductSaleInfoFunc func(ctx context.Context, key Key) (Sale, error)

	// GetProductStockInfoFunc mocks the GetProductStockInfo method.
	GetProductStockInfoFunc func(ctx context.Context, key Key) (Stock, error)

	// GetProductsFunc mocks the GetProducts method.
	GetProductsFunc func(ctx context.Context, key Key) (Product, error)

	// ListIngredientsFunc mocks the ListIngredients method.
	ListIngredientsFunc func(ctx context.Context, productKey Key) ([]Ingredient, error)

	// ListProductsFunc mocks the ListProducts method.
	ListProductsFunc func(ctx context.Context, productFilter Filter) ([]Product, int, error)

	// RemoveAllIngredientsFunc mocks the RemoveAllIngredients method.
	RemoveAllIngredientsFunc func(ctx context.Context, productKey Key) error

	// RemoveIngredientFunc mocks the RemoveIngredient method.
	RemoveIngredientFunc func(ctx context.Context, productKey Key, ingredientID types.UID) error

	// RemoveProductFunc mocks the RemoveProduct method.
	RemoveProductFunc func(ctx context.Context, key Key) error

	// SaveIngredientsFunc mocks the SaveIngredients method.
	SaveIngredientsFunc func(ctx context.Context, productKey Key, ingredients ...Ingredient) ([]Ingredient, error)

	// UpdateProductFunc mocks the UpdateProduct method.
	UpdateProductFunc func(ctx context.Context, prd Product) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateProduct holds details about calls to the CreateProduct method.
		CreateProduct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Prd is the prd argument value.
			Prd Product
		}
		// GetProductSaleInfo holds details about calls to the GetProductSaleInfo method.
		GetProductSaleInfo []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Key is the key argument value.
			Key Key
		}
		// GetProductStockInfo holds details about calls to the GetProductStockInfo method.
		GetProductStockInfo []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Key is the key argument value.
			Key Key
		}
		// GetProducts holds details about calls to the GetProducts method.
		GetProducts []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Key is the key argument value.
			Key Key
		}
		// ListIngredients holds details about calls to the ListIngredients method.
		ListIngredients []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ProductKey is the productKey argument value.
			ProductKey Key
		}
		// ListProducts holds details about calls to the ListProducts method.
		ListProducts []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ProductFilter is the productFilter argument value.
			ProductFilter Filter
		}
		// RemoveAllIngredients holds details about calls to the RemoveAllIngredients method.
		RemoveAllIngredients []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ProductKey is the productKey argument value.
			ProductKey Key
		}
		// RemoveIngredient holds details about calls to the RemoveIngredient method.
		RemoveIngredient []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ProductKey is the productKey argument value.
			ProductKey Key
			// IngredientID is the ingredientID argument value.
			IngredientID types.UID
		}
		// RemoveProduct holds details about calls to the RemoveProduct method.
		RemoveProduct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Key is the key argument value.
			Key Key
		}
		// SaveIngredients holds details about calls to the SaveIngredients method.
		SaveIngredients []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// ProductKey is the productKey argument value.
			ProductKey Key
			// Ingredients is the ingredients argument value.
			Ingredients []Ingredient
		}
		// UpdateProduct holds details about calls to the UpdateProduct method.
		UpdateProduct []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Prd is the prd argument value.
			Prd Product
		}
	}
	lockCreateProduct        sync.RWMutex
	lockGetProductSaleInfo   sync.RWMutex
	lockGetProductStockInfo  sync.RWMutex
	lockGetProducts          sync.RWMutex
	lockListIngredients      sync.RWMutex
	lockListProducts         sync.RWMutex
	lockRemoveAllIngredients sync.RWMutex
	lockRemoveIngredient     sync.RWMutex
	lockRemoveProduct        sync.RWMutex
	lockSaveIngredients      sync.RWMutex
	lockUpdateProduct        sync.RWMutex
}

// CreateProduct calls CreateProductFunc.
func (mock *ActionsMock) CreateProduct(ctx context.Context, prd Product) (Product, error) {
	if mock.CreateProductFunc == nil {
		panic("ActionsMock.CreateProductFunc: method is nil but Actions.CreateProduct was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Prd Product
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
//     len(mockedActions.CreateProductCalls())
func (mock *ActionsMock) CreateProductCalls() []struct {
	Ctx context.Context
	Prd Product
} {
	var calls []struct {
		Ctx context.Context
		Prd Product
	}
	mock.lockCreateProduct.RLock()
	calls = mock.calls.CreateProduct
	mock.lockCreateProduct.RUnlock()
	return calls
}

// GetProductSaleInfo calls GetProductSaleInfoFunc.
func (mock *ActionsMock) GetProductSaleInfo(ctx context.Context, key Key) (Sale, error) {
	if mock.GetProductSaleInfoFunc == nil {
		panic("ActionsMock.GetProductSaleInfoFunc: method is nil but Actions.GetProductSaleInfo was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Key Key
	}{
		Ctx: ctx,
		Key: key,
	}
	mock.lockGetProductSaleInfo.Lock()
	mock.calls.GetProductSaleInfo = append(mock.calls.GetProductSaleInfo, callInfo)
	mock.lockGetProductSaleInfo.Unlock()
	return mock.GetProductSaleInfoFunc(ctx, key)
}

// GetProductSaleInfoCalls gets all the calls that were made to GetProductSaleInfo.
// Check the length with:
//     len(mockedActions.GetProductSaleInfoCalls())
func (mock *ActionsMock) GetProductSaleInfoCalls() []struct {
	Ctx context.Context
	Key Key
} {
	var calls []struct {
		Ctx context.Context
		Key Key
	}
	mock.lockGetProductSaleInfo.RLock()
	calls = mock.calls.GetProductSaleInfo
	mock.lockGetProductSaleInfo.RUnlock()
	return calls
}

// GetProductStockInfo calls GetProductStockInfoFunc.
func (mock *ActionsMock) GetProductStockInfo(ctx context.Context, key Key) (Stock, error) {
	if mock.GetProductStockInfoFunc == nil {
		panic("ActionsMock.GetProductStockInfoFunc: method is nil but Actions.GetProductStockInfo was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Key Key
	}{
		Ctx: ctx,
		Key: key,
	}
	mock.lockGetProductStockInfo.Lock()
	mock.calls.GetProductStockInfo = append(mock.calls.GetProductStockInfo, callInfo)
	mock.lockGetProductStockInfo.Unlock()
	return mock.GetProductStockInfoFunc(ctx, key)
}

// GetProductStockInfoCalls gets all the calls that were made to GetProductStockInfo.
// Check the length with:
//     len(mockedActions.GetProductStockInfoCalls())
func (mock *ActionsMock) GetProductStockInfoCalls() []struct {
	Ctx context.Context
	Key Key
} {
	var calls []struct {
		Ctx context.Context
		Key Key
	}
	mock.lockGetProductStockInfo.RLock()
	calls = mock.calls.GetProductStockInfo
	mock.lockGetProductStockInfo.RUnlock()
	return calls
}

// GetProducts calls GetProductsFunc.
func (mock *ActionsMock) GetProducts(ctx context.Context, key Key) (Product, error) {
	if mock.GetProductsFunc == nil {
		panic("ActionsMock.GetProductsFunc: method is nil but Actions.GetProducts was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Key Key
	}{
		Ctx: ctx,
		Key: key,
	}
	mock.lockGetProducts.Lock()
	mock.calls.GetProducts = append(mock.calls.GetProducts, callInfo)
	mock.lockGetProducts.Unlock()
	return mock.GetProductsFunc(ctx, key)
}

// GetProductsCalls gets all the calls that were made to GetProducts.
// Check the length with:
//     len(mockedActions.GetProductsCalls())
func (mock *ActionsMock) GetProductsCalls() []struct {
	Ctx context.Context
	Key Key
} {
	var calls []struct {
		Ctx context.Context
		Key Key
	}
	mock.lockGetProducts.RLock()
	calls = mock.calls.GetProducts
	mock.lockGetProducts.RUnlock()
	return calls
}

// ListIngredients calls ListIngredientsFunc.
func (mock *ActionsMock) ListIngredients(ctx context.Context, productKey Key) ([]Ingredient, error) {
	if mock.ListIngredientsFunc == nil {
		panic("ActionsMock.ListIngredientsFunc: method is nil but Actions.ListIngredients was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		ProductKey Key
	}{
		Ctx:        ctx,
		ProductKey: productKey,
	}
	mock.lockListIngredients.Lock()
	mock.calls.ListIngredients = append(mock.calls.ListIngredients, callInfo)
	mock.lockListIngredients.Unlock()
	return mock.ListIngredientsFunc(ctx, productKey)
}

// ListIngredientsCalls gets all the calls that were made to ListIngredients.
// Check the length with:
//     len(mockedActions.ListIngredientsCalls())
func (mock *ActionsMock) ListIngredientsCalls() []struct {
	Ctx        context.Context
	ProductKey Key
} {
	var calls []struct {
		Ctx        context.Context
		ProductKey Key
	}
	mock.lockListIngredients.RLock()
	calls = mock.calls.ListIngredients
	mock.lockListIngredients.RUnlock()
	return calls
}

// ListProducts calls ListProductsFunc.
func (mock *ActionsMock) ListProducts(ctx context.Context, productFilter Filter) ([]Product, int, error) {
	if mock.ListProductsFunc == nil {
		panic("ActionsMock.ListProductsFunc: method is nil but Actions.ListProducts was just called")
	}
	callInfo := struct {
		Ctx           context.Context
		ProductFilter Filter
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
//     len(mockedActions.ListProductsCalls())
func (mock *ActionsMock) ListProductsCalls() []struct {
	Ctx           context.Context
	ProductFilter Filter
} {
	var calls []struct {
		Ctx           context.Context
		ProductFilter Filter
	}
	mock.lockListProducts.RLock()
	calls = mock.calls.ListProducts
	mock.lockListProducts.RUnlock()
	return calls
}

// RemoveAllIngredients calls RemoveAllIngredientsFunc.
func (mock *ActionsMock) RemoveAllIngredients(ctx context.Context, productKey Key) error {
	if mock.RemoveAllIngredientsFunc == nil {
		panic("ActionsMock.RemoveAllIngredientsFunc: method is nil but Actions.RemoveAllIngredients was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		ProductKey Key
	}{
		Ctx:        ctx,
		ProductKey: productKey,
	}
	mock.lockRemoveAllIngredients.Lock()
	mock.calls.RemoveAllIngredients = append(mock.calls.RemoveAllIngredients, callInfo)
	mock.lockRemoveAllIngredients.Unlock()
	return mock.RemoveAllIngredientsFunc(ctx, productKey)
}

// RemoveAllIngredientsCalls gets all the calls that were made to RemoveAllIngredients.
// Check the length with:
//     len(mockedActions.RemoveAllIngredientsCalls())
func (mock *ActionsMock) RemoveAllIngredientsCalls() []struct {
	Ctx        context.Context
	ProductKey Key
} {
	var calls []struct {
		Ctx        context.Context
		ProductKey Key
	}
	mock.lockRemoveAllIngredients.RLock()
	calls = mock.calls.RemoveAllIngredients
	mock.lockRemoveAllIngredients.RUnlock()
	return calls
}

// RemoveIngredient calls RemoveIngredientFunc.
func (mock *ActionsMock) RemoveIngredient(ctx context.Context, productKey Key, ingredientID types.UID) error {
	if mock.RemoveIngredientFunc == nil {
		panic("ActionsMock.RemoveIngredientFunc: method is nil but Actions.RemoveIngredient was just called")
	}
	callInfo := struct {
		Ctx          context.Context
		ProductKey   Key
		IngredientID types.UID
	}{
		Ctx:          ctx,
		ProductKey:   productKey,
		IngredientID: ingredientID,
	}
	mock.lockRemoveIngredient.Lock()
	mock.calls.RemoveIngredient = append(mock.calls.RemoveIngredient, callInfo)
	mock.lockRemoveIngredient.Unlock()
	return mock.RemoveIngredientFunc(ctx, productKey, ingredientID)
}

// RemoveIngredientCalls gets all the calls that were made to RemoveIngredient.
// Check the length with:
//     len(mockedActions.RemoveIngredientCalls())
func (mock *ActionsMock) RemoveIngredientCalls() []struct {
	Ctx          context.Context
	ProductKey   Key
	IngredientID types.UID
} {
	var calls []struct {
		Ctx          context.Context
		ProductKey   Key
		IngredientID types.UID
	}
	mock.lockRemoveIngredient.RLock()
	calls = mock.calls.RemoveIngredient
	mock.lockRemoveIngredient.RUnlock()
	return calls
}

// RemoveProduct calls RemoveProductFunc.
func (mock *ActionsMock) RemoveProduct(ctx context.Context, key Key) error {
	if mock.RemoveProductFunc == nil {
		panic("ActionsMock.RemoveProductFunc: method is nil but Actions.RemoveProduct was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Key Key
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
//     len(mockedActions.RemoveProductCalls())
func (mock *ActionsMock) RemoveProductCalls() []struct {
	Ctx context.Context
	Key Key
} {
	var calls []struct {
		Ctx context.Context
		Key Key
	}
	mock.lockRemoveProduct.RLock()
	calls = mock.calls.RemoveProduct
	mock.lockRemoveProduct.RUnlock()
	return calls
}

// SaveIngredients calls SaveIngredientsFunc.
func (mock *ActionsMock) SaveIngredients(ctx context.Context, productKey Key, ingredients ...Ingredient) ([]Ingredient, error) {
	if mock.SaveIngredientsFunc == nil {
		panic("ActionsMock.SaveIngredientsFunc: method is nil but Actions.SaveIngredients was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		ProductKey  Key
		Ingredients []Ingredient
	}{
		Ctx:         ctx,
		ProductKey:  productKey,
		Ingredients: ingredients,
	}
	mock.lockSaveIngredients.Lock()
	mock.calls.SaveIngredients = append(mock.calls.SaveIngredients, callInfo)
	mock.lockSaveIngredients.Unlock()
	return mock.SaveIngredientsFunc(ctx, productKey, ingredients...)
}

// SaveIngredientsCalls gets all the calls that were made to SaveIngredients.
// Check the length with:
//     len(mockedActions.SaveIngredientsCalls())
func (mock *ActionsMock) SaveIngredientsCalls() []struct {
	Ctx         context.Context
	ProductKey  Key
	Ingredients []Ingredient
} {
	var calls []struct {
		Ctx         context.Context
		ProductKey  Key
		Ingredients []Ingredient
	}
	mock.lockSaveIngredients.RLock()
	calls = mock.calls.SaveIngredients
	mock.lockSaveIngredients.RUnlock()
	return calls
}

// UpdateProduct calls UpdateProductFunc.
func (mock *ActionsMock) UpdateProduct(ctx context.Context, prd Product) error {
	if mock.UpdateProductFunc == nil {
		panic("ActionsMock.UpdateProductFunc: method is nil but Actions.UpdateProduct was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Prd Product
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
//     len(mockedActions.UpdateProductCalls())
func (mock *ActionsMock) UpdateProductCalls() []struct {
	Ctx context.Context
	Prd Product
} {
	var calls []struct {
		Ctx context.Context
		Prd Product
	}
	mock.lockUpdateProduct.RLock()
	calls = mock.calls.UpdateProduct
	mock.lockUpdateProduct.RUnlock()
	return calls
}