package event

import (
	"time"
)

type Event struct {
	Id           uint
	Title        uint
	Event_Status Event_Status
	Start        time.Time
	End          time.Time
	Description  string
	Poster       string
	Archived_at  time.Time
	Created_at   time.Time
	Updated_at   time.Time
}

type Event_Status string

const (
	INACTIVE Event_Status = "inactive"
	OPEN     Event_Status = "open"
	ARCHIVED Event_Status = "archived"
)
