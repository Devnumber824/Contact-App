package admin

import (
	"fmt"
	"pranav/user"
)

type Admin struct {
	Firstname  string
	Lastname   string
	EmployeeID int
}

func Createnewadmin(firstname, lastname string, id int) *Admin {
	var admin1 Admin
	admin1.Firstname = firstname
	admin1.Lastname = lastname
	admin1.EmployeeID = id

	return &admin1
}

func (a *Admin) Createuser(firstname, lastname string) user.User {
	var newuser = user.CreateNewUser(firstname, lastname)

	return *newuser
}

func (a *Admin) Readuser(user user.User) {
	fmt.Println(user)
}

func (a *Admin) Updateuser(user user.User, name string, val int) {
	if val == 1 {
		user.Firstname = name
	} else if val == 2 {
		user.Lastname = name
	} else {
		fmt.Println("Invalid input")
	}
}
