package user



type UserManager interface{
	User()
	Users()
	SignUp()
	EditUser()
	DeleteUser()
}


type userManager struct{
	
}