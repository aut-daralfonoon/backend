package event

import "fmt"

func Event(Id uint) (*EventDTO, error) {
	event, err := ReadEvent(Id)
	if err != nil {
		return nil, fmt.Errorf("could not get event: %w", err)
	}

	return MapEventDTO(event,nil, nil)
}

func Events() ([]*EventDTO, error) {
	events, err := ReadEvents()
	if err != nil {
		return nil, err
	}

	eventDtos := []*EventDTO{}
	for _, e := range events {
		prs ,err  := PresentationsByEvent(e.Id)
		if err != nil {
			return nil, err
		}
		presenters, sponsors ,err := getPresenters(prs[0].ID)
		if err != nil {
			return nil, err
		}

		ed, err := MapEventDTO(e,presenters[0], sponsors[0] )
		if err != nil {
			return nil, err
		}
		eventDtos = append(eventDtos, ed)
	}

	return eventDtos, nil
}

func AddEvent(eventDto *EventDTO) error {
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

func validateEvent(event *EventModel) error {
	return nil
}

func ModifyEvent(*EventDTO) error {
	return nil
}

func PublishEvent(EventId uint) error {
	return nil
}

func ArchiveEvent(EventId uint) error {
	return nil
}

func DeleteEvent(EventId uint) error {
	return nil
}

func Register() {

}

func SingleSubjectRegister(regDTO *RegisterDTO) error {
	return nil
}

func GetEventParticipants(EventId uint) []string {
	return nil
}
