package email

import (
	"bytes"
	"fmt"
	"net/mail"
	"github.com/bytbox/go-pop3"
)

type EmailRetriever interface {
	Retr(int) (string, error)
}

type EmailFetcher interface {
	FetchEmails() error
}

type tlsEmailFetcher struct {
	username string
	password string
	popUrl   string
	popPort  int
}

//NewTlsEmailFetcher 带ssl链接对象
func NewTlsEmailFetcher(username, password, url string, port int) *tlsEmailFetcher {
	return &tlsEmailFetcher{
		username: username,
		password: password,
		popUrl:   url,
		popPort:  port,
	}
}

func (f *tlsEmailFetcher) FetchEmails(handle func(*mail.Message,string)(bool,error)) error {
	uri := fmt.Sprintf("%s:%d", f.popUrl, f.popPort)
	client, err := pop3.DialTLS(uri)
	if err != nil {
		return fmt.Errorf("could not dial server: %v", err)
	}
	defer client.Quit()

	err = client.Auth(f.username, f.password)
	if err != nil {
		return fmt.Errorf("could not authenticate: %v", err)
	}

	msgIds, _, err := client.ListAll()
	if err != nil {
		return fmt.Errorf("could not list messages: %v", err)
	}

	return f.harvestMessages(client, msgIds, handle)
}

func (f *tlsEmailFetcher) harvestMessages(retriever EmailRetriever, msgIds []int, handle func(*mail.Message, string)(bool,error)) error {
	for _, id := range msgIds {
		text, err := retriever.Retr(id)
		if err != nil {
			return fmt.Errorf("could not retrieve message (id=%d): %v", id, err)
		}
		msg, err := mail.ReadMessage(bytes.NewBufferString(text))
		if err != nil {
			return fmt.Errorf("could not read message (id=%d): %v", id, err)
		}
		_,err = handle(msg,text)

		if err != nil {
			return err
		}
	}
	return nil
}
