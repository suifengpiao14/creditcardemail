package main

import (
	"testing"

	"./email"
)

func TestHandle(t *testing.T) {
	//filename:="log/email/email_35.log"
	//filename:="log/email/email_36.log"
	//filename:="log/email/email_37.log"
	filename:="log/email/email.log"
	emailText:=email.ReadEmailFromFile(filename)
	msg := email.MockEmail(emailText)
	Handle(msg,emailText)
}
