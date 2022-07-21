package event

type EventManager interface {
	Event(Id uint) (*Event, error)
	Events() ([]*Event, error)
	AddEvent()
	ModifyEvent()
	PublishEvent()
	ArchiveEvent()
	DeleteEvent()
	Register() 
	SingleSubjectRegister(eventId uint, userEmail string) error
	GetEventParticipants()
}
