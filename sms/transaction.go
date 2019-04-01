package sms

import (
	"isms/sms/util"
	"fmt"
)

// Send SMS layer.
func Send(contact util.Contact) {
	contact = contact.ValidatedPhone()

	url := util.PopulateUrl(&contact)

	fmt.Println(url)
}