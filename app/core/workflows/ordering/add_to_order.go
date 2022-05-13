package ordering

import (
	"context"
	"gomies/app/core/entities/order"
	"gomies/app/sdk/fault"
	"sync"
)

func (w workflow) AddToOrder(ctx context.Context, i order.Item) (ItemAdditionResult, error) {
	const operation = "Workflow.Ordering.AddToOrder"

	if i.OrderID.Empty() {
		return ItemAdditionResult{}, fault.Wrap(fault.ErrMissingUID, operation)
	}

	i, err := w.orders.CreateItem(ctx, i)
	if err != nil {
		return ItemAdditionResult{}, fault.Wrap(err, operation)
	}

	contentNumber := len(i.Products)
	reservations := make(chan Reservation, contentNumber)
	failures := make(chan error, contentNumber)

	wg := sync.WaitGroup{}
	wg.Add(contentNumber)
	for _, product := range i.Products {
		const operation = "Workflow.Ordering.AddToOrder.ContentRoutine"
		product := product

		go func() {
			defer wg.Done()

			product, err := w.orders.CreateContent(ctx, product)
			if err != nil {
				failures <- fault.Wrap(err, operation)
			}

			res, err := w.products.ReserveResources(ctx, i.ID, Reservation{
				ID:        i.ID,
				ProductID: product.ProductID,
				Quantity:  product.Quantity,
				Ignore:    product.Details.IgnoreIngredients,
				Replace:   product.Details.ReplaceIngredients,
			})
			if err != nil {
				failures <- fault.Wrap(err, operation)
			}

			reservations <- res
		}()
	}

	wg.Wait()
	close(failures)
	close(reservations)

	if err := <-failures; err != nil {
		return ItemAdditionResult{}, fault.Wrap(err, operation)
	}

	var result ItemAdditionResult
	for r := range reservations {
		if len(r.Failures) > 0 {
			result.Failed = append(result.Failed, r)
			continue
		}

		result.Succeeded = append(result.Succeeded, r)
	}

	return result, nil
}