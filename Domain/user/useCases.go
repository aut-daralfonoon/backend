package user



type UserManager interface{
	User()
	Users()
	CreateUser()
	UpdateUser()
	DeleteUser()
}


type userInteractor struct{
	
}