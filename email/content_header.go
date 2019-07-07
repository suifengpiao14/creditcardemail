package email

import (
	"errors"
	"regexp"
	"strings"
)

//ContentHeader 邮件body中header对象
type ContentHeader struct {
	ContentType string
	Charset     string
	Encoding    string
}



func (instance *ContentHeader) setEncoding(header string) (err error) {
	reg := regexp.MustCompile(`: *([\w-]+)`)
	matches := reg.FindStringSubmatch(header)
	if len(matches) < 2 {
		err = errors.New("not found Content-Transfer-Encoding")
		return
	}
	instance.Encoding = strings.TrimSpace(matches[1])
	return
}
func (instance *ContentHeader) setContentType(header string) (err error) {
	reg := regexp.MustCompile(`: *([\w/]+)`)
	matches := reg.FindStringSubmatch(header)
	if len(matches) < 2 {
		err = errors.New("not found Content-Type")
		return
	}
	instance.ContentType = strings.TrimSpace(matches[1])
	return
}

func (instance *ContentHeader) setCharset(header string) (err error) {
	reg := regexp.MustCompile(`=["]?([\w-]+)`)
	matches := reg.FindStringSubmatch(header)
	if len(matches) < 2 {
		err = errors.New("not found charset")
		return
	}
	instance.Charset = strings.TrimSpace(matches[1])
	return
}


