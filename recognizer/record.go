package recognizer

import (
	"time"
	"strings"
	"errors"
)

//Record 信用卡消费记录模型
type Record struct {
	Id                  int       `json:"id" xorm:"int(1) not null pk autoincr comment('ID')"`
	Owner              string    `json:"owner" xorm:"varchar(100) not null comment('卡所有者')"`
	OwnerEmail          string    `json:"owner_email" xorm:"varchar(100) not null comment('卡所有者邮箱')"`
	Issuer              string    `json:"issuer" xorm:"varchar(100) not null comment('发卡机构')"`
	Product              string    `json:"product" xorm:"varchar(100) not null comment('产品')"`
	Subject             string    `json:"subject" xorm:"varchar(100) not null comment('邮件主题')"`
	Type                string    `json:"type" xorm:"varchar(100) not null comment('信息类型，月账单-month_bill，消费明细 consumption_details')"`
	ConsumptionDate     time.Time `json:"consumption_date" xorm:"datetime  comment('消费日')"`
	RepaymentDate       time.Time `json:"repayment_date" xorm:"datetime  comment('还款日')"`
	BillingDate         time.Time `json:"billing_date" xorm:"datetime  comment('账单日')"`
	MinReplaymentAmount float64   `json:"min_replayment_amount" xorm:"float(15,2)  comment('最低还款')"`
	Amount              float64   `json:"amount" xorm:"float(15,2) not null comment('金额')"`
	Quota              float64   `json:"quota" xorm:"float(15,2) not null comment('额度')"`
	UsedQuota              float64   `json:"used_quota" xorm:"float(15,2) not null comment('已用额度')"`
	RemainingQuota              float64   `json:"remaining_quota" xorm:"float(15,2) not null comment('剩余额度')"`
	CardNo                string    `json:"card_no" xorm:"varchar(100)  comment('卡号后四位')"`
	Merchant            string    `json:"merchant" xorm:"varchar(100) comment('刷卡商户')"`
	Unique              string    `json:"unique" xorm:"varchar(100) not null comment('唯一键')"`
	CreatedTime         time.Time `json:"created_time" xorm:"datetime not null created  comment('记录采集时间')"`
}

const (
	// RecordTypeBill 月账单类型常量
	RecordTypeBill = "month_bill"
	// RecordTypeDetails 消费记录常量
	RecordTypeDetails = "consumption_details"
)


//SetSubject 识别主题
func (instance *Record)SetSubject(subject string) (err error) {
	instance.Subject = subject
	return
}

//SetEmailer 识别邮件接收者
func (instance *Record)SetEmailer(emailer Emailer) (err error) {
	instance.Issuer = emailer.Issuer
	return
}
//SetOwner 识别邮件接收者
func (instance *Record)SetOwner(address string) (err error) {
	instance.Owner = address
	instance.OwnerEmail = address
	return
}

//SetType 识别资料类型
func (instance *Record)SetType(subject string) (err error) {
	if strings.Contains(subject, "账单") {
		instance.Type = RecordTypeBill
		return nil
	}
	if strings.Contains(subject, "消费") {
		instance.Type = RecordTypeDetails
		return nil
	}
	err = errors.New("not found type")
	return
}
