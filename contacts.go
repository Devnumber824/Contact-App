package contact

import (
	"pranav/contactdetails"
	"strings"

	"github.com/google/uuid"
)

type Contact struct {
	contactdetail []contactdetails.ContactDetails
	contactid     string
	Firstname     string
	Lastname      string
}

func (c *Contact) IsContactExist(contactid string) bool {
	return c.contactid == contactid
}

func CreateNewContact(firstname, lastname string) *Contact {
	var newcontact Contact

	newcontact.Firstname = firstname
	newcontact.Lastname = lastname
	test := uuid.New().String()
	newcontact.contactid = strings.Replace(test, "-", "", -7)

	return &newcontact
}

func (c *Contact) Adddetail(Type, val string) {
	var condetail = contactdetails.AddDetails(Type, val)
	c.contactdetail = append(c.contactdetail, *condetail)

}
