package stock

import (
	"gomies/pkg/sdk/types"
	"time"
)

type (
	Filter struct {
		ResourceID  types.UID
		InitialDate time.Time
		FinalDate   time.Time
		types.Filter
	}
)

func (f Filter) Validate() error {
	if f.ResourceID.Empty() {
		return ErrMissingResourceID
	}

	if f.FinalDate.Before(f.InitialDate) {
		return ErrInvalidPeriod
	}

	return nil
}
