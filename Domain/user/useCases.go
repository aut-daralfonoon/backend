package user



type userInteractor interface{
	User()
	Users()
	CreateUser()
	UpdateUser()
	DeleteUser()
}