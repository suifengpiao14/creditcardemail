package recognizer

import (
	"errors"
	"regexp"
	"../email"
)

//PuFa 招行信用卡识别器
type PuFa struct {
	Email  *email.Email
	Record *Record
}

//Run 识别
func (instance *PuFa) Run() (err error) {
	if instance.Record.Type == RecordTypeBill {
		err = instance.SetRepaymentDate()
		err = instance.SetAmount()
		err = instance.SetMinReplaymentAmount()
		return
	}
	return
}
//GetRecord 获取结果
func (instance *PuFa) GetRecord()*Record {
	return instance.Record
}

//SetRepaymentDate 识别还款日
func (instance *PuFa) SetRepaymentDate() (err error) {
	text := instance.Email.Text
	reg := regexp.MustCompile(`[0-9]{4}/([0-9]{2})/([0-9]{2})`)
	matches := reg.FindStringSubmatch(text)
	if len(matches) < 3 {
		return errors.New("not found replayment date")
	}
	monthStr := matches[1]
	dayStr := matches[2]
	repaymentTime, err := FormatDate(monthStr, dayStr, instance.Email.Header.Date)
	if err != nil {
		return
	}
	instance.Record.RepaymentDate = repaymentTime
	return
}

//SetAmount 识别金额
func (instance *PuFa) SetAmount() (err error) {
	text := instance.Email.Text
	reg := regexp.MustCompile(`￥([0-9,\.]+)`)
	matches := reg.FindStringSubmatch(text)
	if len(matches) < 2 {
		return errors.New("not found amount")
	}
	amount, err := Amount2Float64(matches[1])
	if err != nil {
		return
	}
	instance.Record.Amount = amount
	return
}

//SetMinReplaymentAmount 识别最低还款
func (instance *PuFa) SetMinReplaymentAmount() (err error) {
	return //没有最低
}
