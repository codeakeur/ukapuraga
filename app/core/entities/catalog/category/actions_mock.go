// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package category

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
// 			CreateCategoryFunc: func(ctx context.Context, cat Category) (Category, error) {
// 				panic("mock out the CreateCategory method")
// 			},
// 			GetCategoryFunc: func(ctx context.Context, categoryKey Key) (Category, error) {
// 				panic("mock out the GetCategory method")
// 			},
// 			ListCategoriesFunc: func(ctx context.Context, categoryFilter Filter) ([]Category, int, error) {
// 				panic("mock out the ListCategories method")
// 			},
// 			RemoveCategoryFunc: func(ctx context.Context, categoryID Key) error {
// 				panic("mock out the RemoveCategory method")
// 			},
// 			UpdateCategoryFunc: func(ctx context.Context, category Category) error {
// 				panic("mock out the UpdateCategory method")
// 			},
// 		}
//
// 		// use mockedActions in code that requires Actions
// 		// and then make assertions.
//
// 	}
type ActionsMock struct {
	// CreateCategoryFunc mocks the CreateCategory method.
	CreateCategoryFunc func(ctx context.Context, cat Category) (Category, error)

	// GetCategoryFunc mocks the GetCategory method.
	GetCategoryFunc func(ctx context.Context, categoryKey Key) (Category, error)

	// ListCategoriesFunc mocks the ListCategories method.
	ListCategoriesFunc func(ctx context.Context, categoryFilter Filter) ([]Category, int, error)

	// RemoveCategoryFunc mocks the RemoveCategory method.
	RemoveCategoryFunc func(ctx context.Context, categoryID Key) error

	// UpdateCategoryFunc mocks the UpdateCategory method.
	UpdateCategoryFunc func(ctx context.Context, category Category) error

	// calls tracks calls to the methods.
	calls struct {
		// CreateCategory holds details about calls to the CreateCategory method.
		CreateCategory []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Cat is the cat argument value.
			Cat Category
		}
		// GetCategory holds details about calls to the GetCategory method.
		GetCategory []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CategoryKey is the categoryKey argument value.
			CategoryKey Key
		}
		// ListCategories holds details about calls to the ListCategories method.
		ListCategories []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CategoryFilter is the categoryFilter argument value.
			CategoryFilter Filter
		}
		// RemoveCategory holds details about calls to the RemoveCategory method.
		RemoveCategory []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// CategoryID is the categoryID argument value.
			CategoryID Key
		}
		// UpdateCategory holds details about calls to the UpdateCategory method.
		UpdateCategory []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Category is the category argument value.
			Category Category
		}
	}
	lockCreateCategory sync.RWMutex
	lockGetCategory    sync.RWMutex
	lockListCategories sync.RWMutex
	lockRemoveCategory sync.RWMutex
	lockUpdateCategory sync.RWMutex
}

// CreateCategory calls CreateCategoryFunc.
func (mock *ActionsMock) CreateCategory(ctx context.Context, cat Category) (Category, error) {
	if mock.CreateCategoryFunc == nil {
		panic("ActionsMock.CreateCategoryFunc: method is nil but Actions.CreateCategory was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Cat Category
	}{
		Ctx: ctx,
		Cat: cat,
	}
	mock.lockCreateCategory.Lock()
	mock.calls.CreateCategory = append(mock.calls.CreateCategory, callInfo)
	mock.lockCreateCategory.Unlock()
	return mock.CreateCategoryFunc(ctx, cat)
}

// CreateCategoryCalls gets all the calls that were made to CreateCategory.
// Check the length with:
//     len(mockedActions.CreateCategoryCalls())
func (mock *ActionsMock) CreateCategoryCalls() []struct {
	Ctx context.Context
	Cat Category
} {
	var calls []struct {
		Ctx context.Context
		Cat Category
	}
	mock.lockCreateCategory.RLock()
	calls = mock.calls.CreateCategory
	mock.lockCreateCategory.RUnlock()
	return calls
}

// GetCategory calls GetCategoryFunc.
func (mock *ActionsMock) GetCategory(ctx context.Context, categoryKey Key) (Category, error) {
	if mock.GetCategoryFunc == nil {
		panic("ActionsMock.GetCategoryFunc: method is nil but Actions.GetCategory was just called")
	}
	callInfo := struct {
		Ctx         context.Context
		CategoryKey Key
	}{
		Ctx:         ctx,
		CategoryKey: categoryKey,
	}
	mock.lockGetCategory.Lock()
	mock.calls.GetCategory = append(mock.calls.GetCategory, callInfo)
	mock.lockGetCategory.Unlock()
	return mock.GetCategoryFunc(ctx, categoryKey)
}

// GetCategoryCalls gets all the calls that were made to GetCategory.
// Check the length with:
//     len(mockedActions.GetCategoryCalls())
func (mock *ActionsMock) GetCategoryCalls() []struct {
	Ctx         context.Context
	CategoryKey Key
} {
	var calls []struct {
		Ctx         context.Context
		CategoryKey Key
	}
	mock.lockGetCategory.RLock()
	calls = mock.calls.GetCategory
	mock.lockGetCategory.RUnlock()
	return calls
}

// ListCategories calls ListCategoriesFunc.
func (mock *ActionsMock) ListCategories(ctx context.Context, categoryFilter Filter) ([]Category, int, error) {
	if mock.ListCategoriesFunc == nil {
		panic("ActionsMock.ListCategoriesFunc: method is nil but Actions.ListCategories was just called")
	}
	callInfo := struct {
		Ctx            context.Context
		CategoryFilter Filter
	}{
		Ctx:            ctx,
		CategoryFilter: categoryFilter,
	}
	mock.lockListCategories.Lock()
	mock.calls.ListCategories = append(mock.calls.ListCategories, callInfo)
	mock.lockListCategories.Unlock()
	return mock.ListCategoriesFunc(ctx, categoryFilter)
}

// ListCategoriesCalls gets all the calls that were made to ListCategories.
// Check the length with:
//     len(mockedActions.ListCategoriesCalls())
func (mock *ActionsMock) ListCategoriesCalls() []struct {
	Ctx            context.Context
	CategoryFilter Filter
} {
	var calls []struct {
		Ctx            context.Context
		CategoryFilter Filter
	}
	mock.lockListCategories.RLock()
	calls = mock.calls.ListCategories
	mock.lockListCategories.RUnlock()
	return calls
}

// RemoveCategory calls RemoveCategoryFunc.
func (mock *ActionsMock) RemoveCategory(ctx context.Context, categoryID Key) error {
	if mock.RemoveCategoryFunc == nil {
		panic("ActionsMock.RemoveCategoryFunc: method is nil but Actions.RemoveCategory was just called")
	}
	callInfo := struct {
		Ctx        context.Context
		CategoryID Key
	}{
		Ctx:        ctx,
		CategoryID: categoryID,
	}
	mock.lockRemoveCategory.Lock()
	mock.calls.RemoveCategory = append(mock.calls.RemoveCategory, callInfo)
	mock.lockRemoveCategory.Unlock()
	return mock.RemoveCategoryFunc(ctx, categoryID)
}

// RemoveCategoryCalls gets all the calls that were made to RemoveCategory.
// Check the length with:
//     len(mockedActions.RemoveCategoryCalls())
func (mock *ActionsMock) RemoveCategoryCalls() []struct {
	Ctx        context.Context
	CategoryID Key
} {
	var calls []struct {
		Ctx        context.Context
		CategoryID Key
	}
	mock.lockRemoveCategory.RLock()
	calls = mock.calls.RemoveCategory
	mock.lockRemoveCategory.RUnlock()
	return calls
}

// UpdateCategory calls UpdateCategoryFunc.
func (mock *ActionsMock) UpdateCategory(ctx context.Context, category Category) error {
	if mock.UpdateCategoryFunc == nil {
		panic("ActionsMock.UpdateCategoryFunc: method is nil but Actions.UpdateCategory was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Category Category
	}{
		Ctx:      ctx,
		Category: category,
	}
	mock.lockUpdateCategory.Lock()
	mock.calls.UpdateCategory = append(mock.calls.UpdateCategory, callInfo)
	mock.lockUpdateCategory.Unlock()
	return mock.UpdateCategoryFunc(ctx, category)
}

// UpdateCategoryCalls gets all the calls that were made to UpdateCategory.
// Check the length with:
//     len(mockedActions.UpdateCategoryCalls())
func (mock *ActionsMock) UpdateCategoryCalls() []struct {
	Ctx      context.Context
	Category Category
} {
	var calls []struct {
		Ctx      context.Context
		Category Category
	}
	mock.lockUpdateCategory.RLock()
	calls = mock.calls.UpdateCategory
	mock.lockUpdateCategory.RUnlock()
	return calls
}