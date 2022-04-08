package category

import (
	"context"
	"gomies/app/core/entities/catalog/category"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_SaveCategory(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	type (
		args struct {
			category category.Category
		}

		opts struct {
			categories *category.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			want    category.Category
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return category created",
			args: args{
				category: category.Category{
					Code: "PRD",
					Name: "Product",
				},
			},
			want: category.Category{
				Code: "PRD",
				Name: "Product",
			},
			opts: opts{
				categories: &category.ActionsMock{
					CreateCategoryFunc: func(ctx context.Context, prd category.Category) (category.Category, error) {
						return prd, nil
					},
				},
			},
		},
		{
			name: "should return error for invalid code",
			args: args{
				category: category.Category{
					Code: "P",
					Name: "Product",
				},
			},
			wantErr: category.ErrInvalidCode,
			opts: opts{
				categories: &category.ActionsMock{
					CreateCategoryFunc: func(ctx context.Context, prd category.Category) (category.Category, error) {
						return prd, nil
					},
				},
			},
		},
		{
			name: "should return error for invalid name",
			args: args{
				category: category.Category{
					Code: "PRD",
					Name: "-",
				},
			},
			wantErr: category.ErrInvalidName,
			opts: opts{
				categories: &category.ActionsMock{
					CreateCategoryFunc: func(ctx context.Context, prd category.Category) (category.Category, error) {
						return prd, nil
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			wf := NewWorkflow(c.opts.categories)
			ingredient, err := wf.CreateCategory(ctx, c.args.category)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, ingredient)

		})
	}

}