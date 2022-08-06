package ordering

import (
	"comies/app/core/entities/order"
	"comies/app/core/throw"
	"context"
)

func (w workflow) ListOrders(ctx context.Context, f order.Filter) ([]order.Order, error) {
	list, err := w.orders.List(ctx, f)
	if err != nil {
		return nil, throw.Error(err)
	}

	return list, nil
}
