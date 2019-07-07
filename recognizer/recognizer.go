package recognizer

import (
	"errors"
	"strings"
	"../email"
)

// Recognizer 账单识别接口
type Recognizer interface {
	//执行解析
	Run() (err error)
	//获取执行结果
	GetRecord()*Record
}

//NewRecognizer 创建账单识别实例
func NewRecognizer(emailInstance *email.Email) (recognizer Recognizer, err error) {
	emailer ,has:= GetEmailer(emailInstance.Header.From,FromUsers) 
	if !has {
		err = errors.New("not found recognizer emailer")
		return 
	}
	subject := emailInstance.Header.Subject
	record := &Record{}
	//设置主题
	record.SetSubject(subject)
	//设置类型
	record.SetType(subject)

	// 设置邮件接收者
	record.SetOwner(emailInstance.Header.Reciver[0])
	record.SetEmailer(emailer)

	if emailer.Issuer == IssuerZhaoHang {
		if strings.Contains(subject, "招贷账单") { //e招贷账单 账单，e字符编码转换后不好匹配
			recognizer = &EZhaoDai{
				Email:  emailInstance,
				Record: record,
			}
			return
		}
		if strings.Contains(subject, "账单") {
			recognizer = &ZhaoHangBill{
				Email:  emailInstance,
				Record: record,
			}
			return
		}
		if strings.Contains(subject, "消费提醒") {
			recognizer = &ZhaoHangDetail{
				Email:  emailInstance,
				Record: record,
			}
			return
		}
		err = errors.New("not found other zhao hang Recognizer")
		return
	}

	if emailer.Issuer == IssuerPuFa {
		if strings.Contains(subject, "账单") {
			recognizer = &PuFa{
				Email:  emailInstance,
				Record: record,
			}
			return
		}
		err = errors.New("not found other pu fa Recognizer")
		return
	}

	if emailer.Issuer == IssuerZhongXin {
		// 中信账单和详情在一封邮件中
		if strings.Contains(subject, "账单") {
			recognizer = &ZhongXin{
				Email:  emailInstance,
				Record: record,
			}
			return
		}
		err = errors.New("not found other zhong xin Recognizer")
		return
	}

	if emailer.Issuer == IssuerSZNongCunShangYe {
		// 深圳农村商业银行信用卡
		if strings.Contains(subject, "账单") {
			recognizer = &SZNongCunBill{
				Email:  emailInstance,
				Record: record,
			}
			return
		}
		err = errors.New("not found other shen zhen nong cun shang ye Recognizer")
		return
	}
	if emailer.Issuer == IssuerPingAn {
		// 平安银行信用卡
		if strings.Contains(subject, "账单") {
			recognizer = &PingAn{
				Email:  emailInstance,
				Record: record,
			}
			return
		}
		err = errors.New("not found other ping an Recognizer")
		return
	}

	err = errors.New("not found Recognizer")

	return
}
