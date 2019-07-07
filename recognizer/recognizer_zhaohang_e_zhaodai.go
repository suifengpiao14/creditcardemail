package recognizer

import (
	"errors"
	"regexp"

	
	"../email"
)


//EZhaoDai 招行e招贷识别器
type EZhaoDai struct {
	Email *email.Email
	Record *Record
}


//Run 识别
func (instance *EZhaoDai) Run() (err error) {
	if instance.Record.Type == RecordTypeBill {
		err = instance.SetRepaymentDate()
		err = instance.SetAmount()
		err = instance.SetMinReplaymentAmount()
		return
	}
	return
}
//GetRecord 获取结果
func (instance *EZhaoDai) GetRecord()*Record {
	return instance.Record
}


//SetRepaymentDate 识别还款日
func (instance *EZhaoDai) SetRepaymentDate() (err error) {
	text := instance.Email.Text
	reg := regexp.MustCompile(`还款日\s*?([0-9]{2})月([0-9]{2})日`)
	matches := reg.FindStringSubmatch(text)
	if len(matches) == 0 {
		return errors.New("not found replayment date")
	}

	repaymentTime,err:=FormatDate(matches[1],matches[2],instance.Email.Header.Date)
	if err != nil {
		return
	}
	instance.Record.RepaymentDate = repaymentTime
	return
}

//SetAmount 识别金额
func (instance *EZhaoDai) SetAmount() (err error) {
	text := instance.Email.Text
	reg := regexp.MustCompile(`金额￥([0-9,\.]+)`)
	matches := reg.FindStringSubmatch(text)
	if len(matches) == 0 {
		return errors.New("not found amount")
	}
	amountStr := matches[1]
	amount, err := Amount2Float64(amountStr)
	if err != nil {
		return
	}
	instance.Record.Amount = amount
	return
}

//SetMinReplaymentAmount 识别最低还款
func (instance *EZhaoDai) SetMinReplaymentAmount() (err error) {
	text := instance.Email.Text
	reg := regexp.MustCompile(`最低还款￥([0-9,\.]+)`)
	matches := reg.FindStringSubmatch(text)
	if len(matches) == 0 {
		return // 最低还款找不到，不返回错误
	}
	amountStr := matches[1]
	amount, err := Amount2Float64(amountStr)
	if err != nil {
		return
	}
	instance.Record.MinReplaymentAmount = amount
	return
}
