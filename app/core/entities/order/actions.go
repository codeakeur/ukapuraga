package order

import (
	"comies/app/core/types"
	"context"
	"time"
)

//go:generate moq -fmt goimports -out actions_mock.go . Actions:ActionsMock
type Actions interface {
	Create(ctx context.Context, o Order) (Order, error)
	List(ctx context.Context, f Filter) ([]Order, error)
	GetByID(ctx context.Context, id types.ID) (Order, error)
	SetDeliveryMode(ctx context.Context, id types.ID, deliverType DeliveryMode) error
	SetObservation(ctx context.Context, id types.ID, observation string) error
	SetPlacementDate(ctx context.Context, id types.ID, date time.Time) error
	UpdateFlow(ctx context.Context, f FlowUpdate) (FlowUpdate, error)
	ListFlow(ctx context.Context, orderID types.ID) ([]FlowUpdate, error)
	GetStatus(ctx context.Context, orderID types.ID) (Status, error)
	Remove(ctx context.Context, id types.ID) error
}
