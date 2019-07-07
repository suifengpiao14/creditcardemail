package recognizer

import (
	"errors"
	"regexp"
	"strconv"
	"time"
	
	"../email"
)

//ZhongXin 招行信用卡识别器
type ZhongXin struct {
	Email *email.Email
	Record    *Record
}

//Run 识别
func (instance *ZhongXin) Run() (err error) {
	if instance.Record.Type == RecordTypeBill {
		err = instance.ParseBill()
		return
	}
	return
}
//GetRecord 获取结果
func (instance *ZhongXin) GetRecord()*Record {
	return instance.Record
}

//ParseBill 识别账单
func (instance *ZhongXin) ParseBill() (err error) {
	text := instance.Email.Text
	reg := regexp.MustCompile(`到期还款日：(\d{4})年(\d{2})月(\d{2})日\s+RMB([0-9,\.]+)\s+RMB([0-9,\.]+)\s+RMB([0-9,\.]+).*(\d{4})\s+交易日`)
	matches := reg.FindStringSubmatch(text)
	if len(matches) < 7 {
		return errors.New("not regexp info")
	}
	yearInt, err := strconv.Atoi(matches[1])
	if err != nil {
		return
	}
	montInt, err := strconv.Atoi(matches[2])
	month := time.Month(montInt)
	dayInt, err := strconv.Atoi(matches[3])
	instance.Record.RepaymentDate = time.Date(yearInt, month, dayInt, 0, 0, 0, 0, time.Local)


	instance.Record.Amount, err = Amount2Float64(matches[4])
	if err != nil {
		return
	}
	instance.Record.MinReplaymentAmount, err = Amount2Float64(matches[5])
	if err != nil {
		return
	}
	instance.Record.Quota, err = Amount2Float64(matches[6])
	instance.Record.CardNo = matches[7]
	if err != nil {
		return
	}
		
	return
}