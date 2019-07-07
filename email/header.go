package email

import (
	"errors"
	"net/mail"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//Header 邮件中header对象
type Header struct {
	EmailMessageHeader mail.Header
	Subject            string
	Date               time.Time   // 邮件接收时间
	From               []string   //邮件发送者
	Reciver            []string   // 邮件接收者

}

//Parse 解析邮件头部
func (instance *Header) Parse() (err error) {
	// 解析邮件发送者
	instance.ParseFrom()

	// 解析邮件时间
	if err = instance.ParseDate(); err != nil {
		return
	}
	// 解析邮件
	if err = instance.ParseSubject(); err != nil {
		return
	}
	// 解析邮件接收者
	if err = instance.ParseReciver(); err != nil {
		return
	}

	return
}


//ParseDate 解析邮件时间
func (instance *Header) ParseDate() (err error) {
	dateStr := instance.EmailMessageHeader["Date"][0]
	reg := regexp.MustCompile(` (\d+) (\w+) (\d{4}) (\d{2}):(\d{2}):(\d{2})`) // 由于不同邮件，格式有些区别，此处采用正则匹配，方便最大化兼容
	matches := reg.FindStringSubmatch(dateStr)
	if len(matches) < 6 {
		err = errors.New("not found date")
		return
	}

	dayInt, err := strconv.Atoi(matches[1])
	if err != nil {
		return
	}
	monthStr := matches[2]
	monthInt, has := MonthMap[monthStr]
	if !has {
		err = errors.New("not found month")
	}
	month := time.Month(monthInt)
	yearInt, err := strconv.Atoi(matches[3])
	hourInt, err := strconv.Atoi(matches[4])
	minuteInt, err := strconv.Atoi(matches[5])
	secondInt, err := strconv.Atoi(matches[6])

	date := time.Date(yearInt, month, dayInt, hourInt, minuteInt, secondInt, 0, time.Local)
	instance.Date = date
	return
}

//ParseReciver 解析邮件接收者
func (instance *Header) ParseReciver() (err error) {
	recivers := instance.EmailMessageHeader["To"]
	for _, reciver := range recivers {
		reciver = strings.ToLower(reciver)
		string := string(reciver)
		instance.Reciver = append(instance.Reciver, string)
	}
	return
}

//ParseFrom 解析邮件发送者
func (instance *Header) ParseFrom() {
	instance.From = GetFromEmail(instance.EmailMessageHeader["From"])
}

//ParseSubject 解析邮件标题
func (instance *Header) ParseSubject() (err error) {
	subjects := instance.EmailMessageHeader["Subject"]
	subject, err := DecodeEncodedWordSyntax(subjects[0])
	instance.Subject = strings.TrimSpace(subject)
	return
}
