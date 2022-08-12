package event

import (
	"fmt"

	"github.com/AUT-Daralfonoon/back-end/infra/datastore"
	"gorm.io/gorm"
)

type EventManager struct {
	db *gorm.DB
}

func InitManager() *EventManager {
	db := datastore.GetorCreateDBConn()
	return &EventManager{
		db: db,
	}
}

func (manager EventManager) Event(Id uint) (*EventDTO, error) {
	event, err := ReadEvent(Id)
	if err != nil {
		return nil, fmt.Errorf("could not get event: %w", err)
	}

	return MapEventDTO(event, nil, nil)
}

func (manager *EventManager) Events() ([]*EventDTO, error) {
	events, err := ReadEvents()
	if err != nil {
		return nil, err
	}

	eventDtos := []*EventDTO{}
	for _, e := range events {
		prs, err := PresentationsByEvent(e.Id)
		if err != nil {
			return nil, err
		}
		presenters, sponsors, err := getPresenters(prs[0].ID)
		if err != nil {
			return nil, err
		}

		ed, err := MapEventDTO(e, presenters[0], sponsors[0])
		if err != nil {
			return nil, err
		}
		eventDtos = append(eventDtos, ed)
	}

	return eventDtos, nil
}

func (manager *EventManager) AddEvent(eventDto *EventDTO) error {
	event, err := MapEvent(eventDto)
	if err != nil {
		return err
	}

	if err := validateEvent(event); err != nil {
		return fmt.Errorf("invalid event: %w", err)
	}

	if err := CreateEvent(event); err != nil {
		return fmt.Errorf("could not save: %w", err)
	}

	return nil
}

func validateEvent(event *HldEvent) error {
	return nil
}

func (manager *EventManager) ModifyEvent(*EventDTO) error {
	return nil
}

func (manager *EventManager) PublishEvent(EventId uint) error {
	return nil
}

func (manager *EventManager) ArchiveEvent(EventId uint) error {
	return nil
}

func (manager *EventManager) DeleteEvent(EventId uint) error {
	return nil
}

func (manager *EventManager) Register() {

}

func (manager *EventManager) SingleSubjectRegister(regDTO *RegisterDTO) error {
	return nil
}

func (manager *EventManager) GetEventParticipants(EventId uint) []string {
	return nil
}

func Presentation(ID uint) (*PresentationDTO, error) {
	return nil, nil
}

func Presentations() ([]*PresentationDTO, error) {
	return nil, nil
}

func PresentationsByEvent(eventID uint) ([]*PresentationDTO, error) {
	return nil, nil
}

func getPresenters(prID uint) ([]*PresenterDTO, []*SponsorDTO, error) {
	return nil, nil, nil
}
