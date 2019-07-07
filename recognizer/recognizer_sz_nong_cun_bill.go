package recognizer

import (
	"errors"
	"regexp"
	
	"../email"
)

//SZNongCunBill 深圳农村商业银行信用卡识别器
type SZNongCunBill struct {
	Email *email.Email
	Record    *Record
}

//Run 识别
func (instance *SZNongCunBill) Run() (err error) {
	if instance.Record.Type == RecordTypeBill {
		err = instance.ParseBill()
		return
	}
	return
}

//GetRecord 获取结果
func (instance *SZNongCunBill) GetRecord()*Record {
	return instance.Record
}

//ParseBill 识别账单
func (instance *SZNongCunBill) ParseBill() (err error) {
	text := instance.Email.Text
	reg := regexp.MustCompile(`本期应还金额:￥([0-9,\.]+)最低还款额:￥([0-9,\.]+)到期还款日:(\d+)月(\d+)日`)
	matches := reg.FindStringSubmatch(text)
	if len(matches) < 4 {
		return errors.New("SZNongCunBill not regexp info")
	}
	instance.Record.Amount, err = Amount2Float64(matches[1])
	if err != nil {
		return
	}
	instance.Record.MinReplaymentAmount, err = Amount2Float64(matches[2])
	if err != nil {
		return
	}
	instance.Record.RepaymentDate ,err = FormatDate(matches[3], matches[4], instance.Email.Header.Date)
		
	return
}