package email

import (
	"encoding/base64"
	"errors"
	"io/ioutil"
	"mime/quotedprintable"
	"regexp"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

// DecodeEncodedWordSyntax 解码邮件编码格式
func DecodeEncodedWordSyntax(str string) (decoded string, err error) {
	exp := regexp.MustCompile(`=\?{1}([\w-]+)\?{1}([BQbq])\?{1}(.+)\?{1}=`)
	matches := exp.FindStringSubmatch(str)
	if len(matches) < 1 {
		return "", errors.New("DecodeEncodedWordSyntax not match")
	}
	charset := strings.ToLower(matches[1])
	encoding := strings.ToUpper(matches[2])
	encodedText := matches[3]
	encodedStr := string(encodedText)
	if encoding == "B" {
		decoded, err = DecodeBase64(encodedStr)
	} else if encoding == "Q" {
		decoded, err = DecodeQuotedprintable(encodedStr)
	}
	if charset == "gb2312" || charset == "gbk" {
		decoded, err = GBKToUtf8(decoded)
	}

	return
}

// DecodeQuotedprintable 解码邮件中Quoted-printable编码
func DecodeQuotedprintable(str string) (decoded string, err error) {
	decodedByte, err := ioutil.ReadAll(quotedprintable.NewReader(strings.NewReader(str)))
	decoded = string(decodedByte)
	return
}

// DecodeBase64 解码邮件中base64编码
func DecodeBase64(str string) (decoded string, err error) {
	decodedByte, err := base64.StdEncoding.DecodeString(str)
	decoded = string(decodedByte)
	return
}

//GBKToUtf8 Gb2312 转 Utf8
func GBKToUtf8(str string) (output string, err error) {
	reader := transform.NewReader(strings.NewReader(str), simplifiedchinese.GBK.NewDecoder())
	outputBytes, err := ioutil.ReadAll(reader)
	output = string(outputBytes)
	return
}

//ReplaceUtf8Space 替换utf8中特殊空格
func ReplaceUtf8Space(str ,old  string) (output string) {
	utf8Space := string([]byte{0xc2, 0xa0})
	output=strings.ReplaceAll(str,utf8Space,old)
	return 
}
