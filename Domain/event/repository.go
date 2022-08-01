package event

func CreateEvent(event *EventModel) error {
	return nil
}

func ReadEvent(Id uint) (*EventModel, error) {
	var event *EventModel
	return event, nil
}

func ReadEvents() ([]*EventModel, error) {
	var events []*EventModel
	return events, nil
}
