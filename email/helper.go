package email

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/mail"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

//GetFromEmail 获取发件人中的邮件
func GetFromEmail(froms []string) (address []string) {
	exp := regexp.MustCompile(`([\w-]+@[\w-]+(?:\.[\w-]+)+)`)
	address = make([]string, 0)
	for _, from := range froms {
		matches := exp.FindStringSubmatch(from)
		if len(matches) > 0 {
			for i, match := range matches {
				if i > 0 {
					match = strings.ToLower(match)
					address = append(address, string(match))
				}
			}
		}

	}
	return

}


//LogEmail 记录请求的邮件，方便模拟测试
func LogEmail(filename, text string) {
	dir := "log/email"
	filename = fmt.Sprintf("%s/%s", dir, filename)
	data := []byte(text)
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

//MockEmail 日志中读取邮件，方便开放调试
func MockEmail(emailText string) *mail.Message {
	msg, err := mail.ReadMessage(bytes.NewBufferString(emailText))
	if err != nil {
		panic(err)
	}
	return msg
}

//ReadEmailFromFile 从文件中读取文本邮件
func ReadEmailFromFile(filename string) (emailText string) {
	path := fmt.Sprintf("%s/%s", GetRootDir(), filename)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	emailText = string(b)
	return
}

// MonthMap 字母月份改数字
var MonthMap=map[string]int{
	"Jan":1,
	"Feb":2,
	"Mar":3,
	"Apr":4,
	"May":5,
	"Jun":6,
	"Jul":7,
	"Aug":8,
	"Sept":9,
	"Oct":10,
	"Nov":11,
	"Dec":12,
}

//GetRootDir 获取main路径
func GetRootDir() string {
	file, err := exec.LookPath(os.Args[0])
	if strings.Contains(file, "debug") { // 调试的时候入口不是main函数，所以固定写死
		return "D:\\go\\credit_card_email"
	}
	if err != nil {
		panic(err)
	}
	path, err := filepath.Abs(file)
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(path)
	return dir
}
