package stock

import (
	"context"
	"gomies/pkg/sdk/fault"
	"gomies/pkg/sdk/types"
	"gomies/pkg/stocking/core/entities/stock"
)

func (w workflow) RemoveMovement(ctx context.Context, resourceID types.UID, movementID types.UID) error {
	const operation = "Workflows.Stock.RemoveMovement"

	if resourceID.Empty() || movementID.Empty() {
		return fault.Wrap(stock.ErrMissingResourceID, operation)
	}

	err := w.stocks.RemoveMovement(ctx, resourceID, movementID)
	if err != nil {
		return fault.Wrap(err, operation)
	}

	return nil
}
