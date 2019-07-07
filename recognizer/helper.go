package recognizer

import (
	"strings"
	"strconv"
	"time"
)

//GetEmailer 获取邮件发送者(需要传递设定好的Emailer对象数组)
func GetEmailer(recivers []string, relations []Emailer) (emailer Emailer, has bool) {
	emailer, has = RelationalUser(recivers, relations)
	return
}

// RelationalUser 发送邮件者中是否含有关注的用户
func RelationalUser(froms []string, relations []Emailer) (emailer Emailer, has bool) {
	for _, relation := range relations {
		for _, from := range froms {
			if from == relation.Address {
				emailer = relation
				has = true
				return
			}
		}
	}
	return
}



//Amount2Float64 金额转为float64类型
func Amount2Float64(amountStr string) (amount float64, err error) {
	str := strings.ReplaceAll(amountStr, ",", "") // 删除分割符号
	amount, err = strconv.ParseFloat(str, 64)
	return
}

// FormatDate 字符串中提取日期后格式化
func FormatDate(monthStr, dayStr string, emailDate time.Time) (formatTime time.Time, err error) {
	monthInt, err := strconv.Atoi(monthStr)
	if err != nil {
		return
	}
	dayInt, err := strconv.Atoi(dayStr)
	if err != nil {
		return
	}
	month := time.Month(monthInt)
	currentTime := time.Now()
	year := emailDate.Year()
	emialMonth := emailDate.Month()
	if month < emialMonth {
		year++
	}
	formatTime = time.Date(year, month, dayInt, 0, 0, 0, 0, currentTime.Location())
	return
}

