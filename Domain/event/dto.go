package event

import "time"

type EventDTO struct {
	Id          uint         `json:"id"`
	Title       uint         `json:"title"`
	Description string       `json:"description"`
	Date        time.Time    `json:"date"`
	Enable      bool         `json:"enable"`
	Poster      string       `json:"poster"`
	Presenter   PresenterDTO `json:"presenter"`
	Sponsor     SponsorDTO   `json:"sponsor"`
}

type PresentationDTO struct {
	ID uint
}

type PresenterDTO struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type SponsorDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
}

type RegisterDTO struct {
	EventID       uint   `json:"event_id"`
	UserID        uint   `json:"user_id"`
	Presentations []uint `json:"presentations"`
}
