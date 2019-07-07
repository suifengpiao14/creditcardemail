package main

import (
	"fmt"
	"net/mail"
	"os"

	"../conf"

	"../email"
	"../recognizer"
)

func main() {
	Fetch()
}

func Fetch() {
	for _, pop3Account := range conf.Pop3Accounts {
		fetcher := email.NewTlsEmailFetcher(pop3Account.Account, pop3Account.Password, pop3Account.Domain, pop3Account.Port)
		err := fetcher.FetchEmails(Handle)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

//Handle 邮件处理类
func Handle(msg *mail.Message, emailText string) (ok bool, err error) {
	emailBody, err := email.NewEmail(msg)
	if err != nil {
		panic(err)
	}
	// 解析邮件
	emailBody.Header.Parse()
	emailBody.Parse()

	recognizer, err := recognizer.NewRecognizer(emailBody)
	if err != nil|| recognizer ==nil {
		fmt.Println(err.Error())
		return true,nil// 找不到识别器的错误，认为正常
	}
	err = recognizer.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println(recognizer.GetRecord())
	return
}
