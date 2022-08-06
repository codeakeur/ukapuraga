package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/core/throw"
	"comies/app/core/types"
	"context"
)

func (w workflow) SetOrderDeliveryMode(ctx context.Context, id types.ID, deliveryMode order.DeliveryMode) error {
	err := w.orders.SetDeliveryMode(ctx, id, deliveryMode)
	if err != nil {
		return throw.Error(err)
	}

	return nil
}
