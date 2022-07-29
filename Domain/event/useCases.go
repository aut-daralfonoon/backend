package event

type EventManager interface {
	EventModel(Id uint) (*EventDTO, error)
	Events() ([]*EventDTO, error)
	AddEvent(*EventDTO) error
	ModifyEvent(*EventDTO) error
	PublishEvent(EventId uint) error
	ArchiveEvent(EventId uint) error
	DeleteEvent(EventId uint) error
	Register()
	SingleSubjectRegister(eventId uint, userEmail string) error
	GetEventParticipants(EventId uint) []string
}
