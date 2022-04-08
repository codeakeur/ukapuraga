package crew

import (
	"context"
	"gomies/app/core/entities/iam/crew"
	"gomies/pkg/sdk/types"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWorkflow_CreateMember(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	type (
		args struct {
			member crew.Member
		}

		opts struct {
			crew *crew.ActionsMock
		}

		test struct {
			name    string
			args    args
			opts    opts
			want    crew.Member
			wantErr error
		}
	)

	cases := []test{
		{
			name: "should return member created",
			args: args{
				member: crew.Member{
					Name:           "",
					FullName:       "",
					Nickname:       "",
					Reference:      "",
					PictureURL:     "",
					PasswordChange: time.Time{},
					LastSeen:       time.Time{},
					Password:       "",
					Permissions:    "",
					StoreID:        0,
					Store:          types.Store{},
				},
			},
			want: crew.Member{
				Name:           "",
				FullName:       "",
				Nickname:       "",
				Reference:      "",
				PictureURL:     "",
				PasswordChange: time.Time{},
				LastSeen:       time.Time{},
				Password:       "",
				Permissions:    "",
				StoreID:        0,
				Store:          types.Store{},
			},
			opts: opts{
				crew: &crew.ActionsMock{
					CreateMemberFunc: func(ctx context.Context, op crew.Member) (crew.Member, error) {
						return op, nil
					},
				},
			},
		},
	}

	for _, c := range cases {
		c := c

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			wf := workflow{
				crew: c.opts.crew,
			}
			ingredient, err := wf.CreateMember(ctx, c.args.member)

			assert.ErrorIs(t, err, c.wantErr)
			assert.Equal(t, c.want, ingredient)

		})
	}

}