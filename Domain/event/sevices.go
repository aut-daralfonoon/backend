package event

import "fmt"

type eventManager struct {
}

func NewEventManger() EventManager {
	return &eventManager{}
}

func (em *eventManager) Event(Id uint) (*EventDto, error) {
	event := &Event{Id: Id}
	err := em.GetEvent(event)
	if err != nil {
		return nil, fmt.Errorf("could not get event: %w", err)
	}

	return MapEventDto(event)
}

func (em *eventManager) Events() ([]*EventDto, error) {
	var events []*Event
	err := em.GetEvents(events)
	if err != nil {
		return nil, err
	}
	eventDtos := []*EventDto{}
	for _, e := range events {
		ed ,err := MapEventDto(e)
		if err != nil {
			return nil, err
		}
		eventDtos = append(eventDtos, ed )
	}
	return eventDtos, nil
}

func (em *eventManager) AddEvent(eventDto *EventDto) error {
	event,err := MapEvent(eventDto)
	if err != nil {
		return err
	}
	

	return nil
}

func (em *eventManager) ModifyEvent(*EventDto) error {
	return nil
}

func (em *eventManager) PublishEvent(EventId uint) error {
	return nil
}

func (em *eventManager) ArchiveEvent(EventId uint) error {
	return nil
}

func (em *eventManager) DeleteEvent(EventId uint) error {
	return nil
}

func (em *eventManager) Register() {

}

func (em *eventManager) SingleSubjectRegister(eventId uint, userEmail string) error {
	return nil
}

func (em *eventManager) GetEventParticipants(EventId uint) []string {
	return nil
}
