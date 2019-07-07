package recognizer

import (
	"errors"
	"regexp"
	"strconv"
	"time"

	"../email"
)

//PingAn 平安银行信用卡识别器
type PingAn struct {
	Email  *email.Email
	Record *Record
}

//Run 识别
func (instance *PingAn) Run() (err error) {
	if instance.Record.Type == RecordTypeBill {
		err = instance.ParseBill()
		return
	}
	return
}

//GetRecord 获取结果
func (instance *PingAn) GetRecord()*Record {
	return instance.Record
}

//ParseBill 识别账单
func (instance *PingAn) ParseBill() (err error) {
	text := instance.Email.Text
	reg := regexp.MustCompile(`本期还款日(\d{4})-(\d{2})-(\d{2})信用额度¥([\d,\.]+).+本期应还金额本期最低应还金额¥([\d,\.]+)\$[\d,\.]+¥([\d,\.]+)\$`)
	matches := reg.FindStringSubmatch(text)
	if len(matches) < 6 {
		return errors.New("PingAn not regexp info")
	}
	yearInt, err := strconv.Atoi(matches[1])
	if err != nil {
		return
	}
	montInt, err := strconv.Atoi(matches[2])
	month := time.Month(montInt)
	dayInt, err := strconv.Atoi(matches[3])
	instance.Record.RepaymentDate = time.Date(yearInt, month, dayInt, 0, 0, 0, 0, time.Local)

	instance.Record.Quota, err = Amount2Float64(matches[4])
	if err != nil {
		return
	}

	instance.Record.Amount, err = Amount2Float64(matches[5])
	if err != nil {
		return
	}
	instance.Record.MinReplaymentAmount, err = Amount2Float64(matches[6])
	if err != nil {
		return
	}

	return
}
