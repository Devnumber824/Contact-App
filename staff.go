package staff

import (
	"fmt"
	"pranav/user"
)

type Staff struct {
	Firstname  string
	Lastname   string
	Employeeid int
}

func CreateNewStaff(firstname, lastname string, id int) *Staff {
	var newstaff Staff
	newstaff.Firstname = firstname
	newstaff.Lastname = lastname
	newstaff.Employeeid = id
	return &newstaff
}

func (s *Staff) CreateContact(user *user.User, firstname, lastname string) {
	user.AddContact(firstname, lastname)
}

func (s *Staff) CreateContactDetils(user *user.User, firstname, lastname string, Type, val string) {
	user.AddDetail(firstname, lastname, Type, val)
}

func (s *Staff) ReadContact(user *user.User) {
	fmt.Println(user)
}

func UpdateContact(user *user.User, firstname, lastname, name string, val int) {
	i := 0
	for i = 0; i < len(user.Contactsarr); i++ {
		if user.Contactsarr[i].Firstname == firstname && user.Contactsarr[i].Lastname == lastname {
			if val == 1 {
				user.Contactsarr[i].Firstname = name
			} else {
				user.Contactsarr[i].Lastname = name
			}

		}
	}
}
