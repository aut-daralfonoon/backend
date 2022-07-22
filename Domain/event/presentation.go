package event

import "time"

type Presentation struct {
	Id          uint
	EventId     uint
	Title       uint
	Start       time.Time
	End         time.Time
	Archived_at time.Time
	Created_at  time.Time
	Updated_at  time.Time
}
