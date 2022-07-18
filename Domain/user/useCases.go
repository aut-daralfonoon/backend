package user



type UserManager interface{
	User()
	Users()
	RegisterUser()
	EditUser()
	DeleteUser()
}


type userManager struct{
	
}