package product

import (
	"context"
	"gomies/pkg/catalog/core/entities/product"
	"gomies/pkg/sdk/fault"
)

func (w workflow) ListProducts(ctx context.Context, filter product.Filter) ([]product.Product, error) {
	const operation = "Workflows.Product.ListProducts"

	list, err := w.products.List(ctx, filter)
	if err != nil {
		return []product.Product{}, fault.Wrap(err, operation)
	}
	return list, err
}
