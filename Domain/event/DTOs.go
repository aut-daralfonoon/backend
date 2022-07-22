package event

import "time"

type EventDto struct {
	Id          uint         `json:"id"`
	Title       uint         `json:"title"`
	Description string       `json:"description"`
	Date        time.Time    `json:"date"`
	Enable      bool         `json:"enable"`
	Poster      string       `json:"poster"`
	Presenter   PresenterDto `json:"presenter"`
	Sponsor     SponsorDto   `json:"sponsor"`
}

type PresentationDto struct {
}

type PresenterDto struct {
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type SponsorDto struct {
	Name		string	`json:"name"`
	Description	string	`json:"description"`
	Logo		string	`json:"logo"`
}
