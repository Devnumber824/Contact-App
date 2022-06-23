package contactdetails

import (
	"strings"

	"github.com/google/uuid"
)

type ContactDetails struct {
	Type string
	id   string
	data string
}

func AddDetails(Type, value string) *ContactDetails {
	var newcontactDetails ContactDetails
	test := uuid.New().String()
	newcontactDetails.id = strings.Replace(test, "-", "", -5)
	newcontactDetails.Type = Type
	newcontactDetails.data = value

	return &newcontactDetails
}
