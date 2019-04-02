package sms

import (
	"isms/sms/util"
)

// Send direct SMS layer.
func SendDirectSMS(contact util.Contact) {
	contact.ValidatedPhone()
	url := contact.DirectUrl()
	contact.Send(url.String())
}

// Check credit balance layer.
func CheckBalance(contact util.Contact) {
	url := contact.CheckBalanceUrl()
	contact.Send(url.String())
}