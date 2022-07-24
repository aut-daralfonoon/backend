package event

type EventManager interface {
	Event(Id uint) (*EventDto, error)
	Events() ([]*EventDto, error)
	AddEvent(*EventDto) (error)
	ModifyEvent(*EventDto) (error)
	PublishEvent(EventId uint) error
	ArchiveEvent(EventId uint) error
	DeleteEvent(EventId uint) error
	Register()
	SingleSubjectRegister(eventId uint, userEmail string) error
	GetEventParticipants(EventId uint) []string
}
