package user

import (
	"fmt"
	"pranav/contact"
	"strings"

	"github.com/google/uuid"
)

type User struct {
	userid      string
	Contactsarr []contact.Contact
	Firstname   string
	Lastname    string
}

func CreateNewUser(firstname, lastname string) *User {
	test := uuid.New().String()
	var newUser User
	newUser.Firstname = firstname
	newUser.Lastname = lastname
	newUser.userid = strings.Replace(test, "-", "", -5)

	return &newUser
}

func (u *User) AddContact(firstname, lastname string) {

	i := 0
	for i = 0; i < len(u.Contactsarr); i++ {
		if u.Contactsarr[i].Firstname == firstname && u.Contactsarr[i].Lastname == lastname {
			fmt.Println("Contact Already exist")
			return
		}
	}
	var newcontact = contact.CreateNewContact(firstname, lastname)
	u.Contactsarr = append(u.Contactsarr, *newcontact)

}

func (u *User) AddDetail(firstname, lastname, Type, val string) {
	i := 0
	for i = 0; i < len(u.Contactsarr); i++ {
		if u.Contactsarr[i].Firstname == firstname && u.Contactsarr[i].Lastname == lastname {
			u.Contactsarr[i].Adddetail(Type, val)
			return
		}
	}

	fmt.Println("Contact does not exist")
}
