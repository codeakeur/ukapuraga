package products

import (
	"comies/app/core/menu"
	"comies/app/data/conn"
	"context"
	"errors"

	"github.com/jackc/pgconn"
)

func Update(ctx context.Context, prd menu.Product) error {
	const script = `
		update 
			products
		set
			code = $1,
			name = $2,
			type = $3,
			cost_price = $4,
			sale_price = $5,
			sale_unit = $6,
			minimum_sale = $7,
			minimum_quantity = $8,
			maximum_quantity = $9,
			location = $10
		where 
			id = $11
	`

	cmd, err := conn.ExecFromContext(ctx, script,
		prd.Code,
		prd.Name,
		prd.Type,
		prd.CostPrice,
		prd.SalePrice,
		prd.SaleUnit,
		prd.MinimumSale,
		prd.MinimumQuantity,
		prd.MaximumQuantity,
		prd.Location,
		prd.ID,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {

			if pgErr.Code == conn.DuplicateError && pgErr.ConstraintName == conn.ProductCodeUK {
				return menu.ErrCodeAlreadyExists
			}
		}
		return err
	}

	if cmd.RowsAffected() != 1 {
		return menu.ErrNotFound
	}

	return nil
}
