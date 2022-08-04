package event

import (
	"time"
)

type HldEvent struct {
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



type HldPresentation struct {
	Id                uint
	EventId           uint
	Title             uint
	PresentationType  PresentationType
	SessionStatus     SessionStatus
	Capacity          uint
	ParticipantNumber uint
	Start             time.Time
	End               time.Time
	Archived_at       time.Time
	Created_at        time.Time
	Updated_at        time.Time
}

type SessionStatus string

const (
	ClosedToJoin    SessionStatus = "closed_to_join"
	OpenToJoin      SessionStatus = "open_to_join"
	ArchivedSession SessionStatus = "archived"
)

type PresentationType string

const (
	Workshop PresentationType = "workshop"
	Talk     PresentationType = "talk"
)