package types

import "time"

type TimeNow interface {
	Now() time.Time
}
