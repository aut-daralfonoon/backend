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

func (em *eventManager) AddEvent() {

}

func (em *eventManager) ModifyEvent() {

}

func (em *eventManager) PublishEvent() {

}

func (em *eventManager) ArchiveEvent() {

}

func (em *eventManager) DeleteEvent() {

}

func (em *eventManager) Register() {

}

func (em *eventManager) SingleSubjectRegister(eventId uint, userEmail string) error {
	return nil
}

func (em *eventManager) GetEventParticipants() {

}
