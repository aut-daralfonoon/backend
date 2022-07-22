package event

type EventManager interface {
	Event(Id uint) (*EventDto, error)
	Events() ([]*EventDto, error)
	AddEvent()
	ModifyEvent()
	PublishEvent()
	ArchiveEvent()
	DeleteEvent()
	Register() 
	SingleSubjectRegister(eventId uint, userEmail string) error
	GetEventParticipants()
}
