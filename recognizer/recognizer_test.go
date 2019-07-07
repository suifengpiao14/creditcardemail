package recognizer

import (
	"fmt"
	"testing"

	"../email"
)

func TestParseBody(t *testing.T) {
	filename := "log/email/email.log"

	emailText := email.ReadEmailFromFile(filename)
	msg := email.MockEmail(emailText)
	emailBody, err := email.NewEmail(msg)
	if err != nil {
		panic(err)
	}
	// 解析邮件
	emailBody.Header.Parse()
	emailBody.Parse()

	recognizer, err := NewRecognizer(emailBody)
	if err != nil {
		panic(err)
	}
	err = recognizer.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println(emailBody.Text)
	fmt.Println(recognizer.GetRecord())
}
