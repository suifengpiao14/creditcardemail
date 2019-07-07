package email

import (
	"fmt"
	"testing"
)

func TestParseBody(t *testing.T) {
	filename := "log/email/email.log"

	emailText := ReadEmailFromFile(filename)
	msg := MockEmail(emailText)
	emailBody, err := NewEmail(msg)
	if err != nil {
		panic(err)
	}
	// 解析邮件
	emailBody.Header.Parse()
	emailBody.Parse()
	fmt.Println(emailBody)
}
