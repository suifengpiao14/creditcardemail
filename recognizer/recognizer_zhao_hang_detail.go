package recognizer

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	
	"../email"
)

//ZhaoHangDetail 招行信用卡消费记录识别器
type ZhaoHangDetail struct {
	Email *email.Email
	Record    *Record
}

//Run 识别
func (instance *ZhaoHangDetail) Run() (err error) {
	if instance.Record.Type == RecordTypeDetails {
		err = instance.Parse()
		return
	}
	return
}
//GetRecord 获取结果
func (instance *ZhaoHangDetail) GetRecord()*Record {
	return instance.Record
}

//Parse 识别还款日
func (instance *ZhaoHangDetail) Parse() (err error) {
	text := instance.Email.Text
	reg := regexp.MustCompile(`￥([\d,\.]+)￥([\d,\.]+)\p{Han}+(\d{4})(\d{4})(\d{2})(\d{2})(\d{2}):(\d{2}):(\d{2})\p{Han}+ (\p{Han}+)([\d,\.]+)`)
	matches := reg.FindStringSubmatch(text)
	if len(matches) < 11 {
		return errors.New("not regexp info")
	}
	instance.Record.UsedQuota, err = Amount2Float64(matches[1])
	if err != nil {
		return
	}
	instance.Record.RemainingQuota, err = Amount2Float64(matches[2])
	if err != nil {
		return
	}
	instance.Record.CardNo = matches[3]
	yearInt, err := strconv.Atoi(matches[4])
	if err != nil {
		return
	}
	montInt, err := strconv.Atoi(matches[5])
	month := time.Month(montInt)
	dayInt, err := strconv.Atoi(matches[6])
	hourInt, err := strconv.Atoi(matches[7])
	minInt, err := strconv.Atoi(matches[8])
	secondInt, err := strconv.Atoi(matches[9])
	instance.Record.Merchant = matches[10]
	instance.Record.Amount, err = Amount2Float64(matches[11])

	instance.Record.ConsumptionDate = time.Date(yearInt, month, dayInt, hourInt, minInt, secondInt, 0, time.Local)

	return
}
