package event

import "time"

type Event struct {
	Id          uint
	Title       uint
	Description string
	Date        time.Time
	Enable      bool
	Poster      string
}