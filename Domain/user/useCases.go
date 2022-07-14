package user



type UserInteractor interface{
	User()
	Users()
	CreateUser()
	UpdateUser()
	DeleteUser()
}