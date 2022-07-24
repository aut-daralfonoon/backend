package event

type eventManager struct {
}

func NewEventManger() EventManager {
	return &eventManager{}
}

func (em *eventManager) Event(Id uint) (*EventDto, error) {
	
	return nil, nil
}

func (em *eventManager) Events() ([]*EventDto, error) {
	return nil, nil
}

func (em *eventManager) AddEvent(*EventDto) (error) {

}

func (em *eventManager) ModifyEvent(*EventDto) (error) {

}

func (em *eventManager) PublishEvent(EventId uint) error {

}

func (em *eventManager) ArchiveEvent(EventId uint) error {

}

func (em *eventManager) DeleteEvent(EventId uint) error {

}

func (em *eventManager) Register() {

}

func (em *eventManager) SingleSubjectRegister(eventId uint, userEmail string) error {
	return nil
}

func (em *eventManager) GetEventParticipants(EventId uint) []string {

}
