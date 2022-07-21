package event

type EventManager interface{
	Event()
	Events()
	AddEvent()
	ModifyEvent()
	PublishEvent()
	ArchiveEvent()
	DeleteEvent()
	Register()
	SingleSubjectRegister()
	GetEventParticipants()
}