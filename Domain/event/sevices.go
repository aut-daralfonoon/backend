package event

import "fmt"

func GetEvent(Id uint) (*EventDto, error) {
	event, err := ReadEvent(Id)
	if err != nil {
		return nil, fmt.Errorf("could not get event: %w", err)
	}

	return MapEventDto(event)
}

func GetEvents() ([]*EventDto, error) {
	events, err := ReadEvents()
	if err != nil {
		return nil, err
	}

	eventDtos := []*EventDto{}
	for _, e := range events {
		ed, err := MapEventDto(e)
		if err != nil {
			return nil, err
		}
		eventDtos = append(eventDtos, ed)
	}

	return eventDtos, nil
}

func AddEvent(eventDto *EventDto) error {
	event, err := MapEvent(eventDto)
	if err != nil {
		return err
	}

	return nil
}

func validateEvent(event *Event) error{
	return nil
}

func ModifyEvent(*EventDto) error {
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

func SingleSubjectRegister(eventId uint, userEmail string) error {
	return nil
}

func GetEventParticipants(EventId uint) []string {
	return nil
}
