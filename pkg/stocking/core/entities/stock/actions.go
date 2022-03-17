package stock

import (
	"context"
	"gomies/pkg/sdk/types"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock

type Actions interface {
	// ComputeStock computes the quantity stored of a specific resource
	//
	// Possible errors:
	//   - ErrMustHaveTargetID: if the TargetID was not set
	ComputeStock(ctx context.Context, target types.External, stockFilter Filter) (Actual, error)

	// GetMovement retrieves a stock movement with their details
	//
	// Possible errors:
	//   - fault.ErrNotFound: if the movement does not exist
	GetMovement(ctx context.Context, target types.External, movementID types.External) (Movement, error)

	// ListMovements retrieves stock movements for a specific resource
	// Possible errors:
	//   - ErrMustHaveTargetID: if the TargetID was not set
	ListMovements(ctx context.Context, target types.External, stockFilter Filter) ([]Movement, error)

	// AddToStock adds a stock movement for a specific resource
	AddToStock(ctx context.Context, target types.External, mov Movement) (Movement, error)

	// RemoveFromStock removes a stock movement from a resource stock
	RemoveFromStock(ctx context.Context, target types.External, movementID types.External) error
}
