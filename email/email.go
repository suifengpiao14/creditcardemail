package email

import (
	"bufio"
	"bytes"
	"errors"
	"net/mail"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Email 邮件主体解析
type Email struct {
	EmailMessage  *mail.Message
	Header        *Header
	ContentHeader *ContentHeader
	Content       string
	Text          string
	Boundary      []byte // 邮件主体分割符
}

//NewEmail 创建邮件内容对象实例
func NewEmail(emailMessage *mail.Message) (email *Email, err error) {
	email = &Email{
		EmailMessage: emailMessage,
		Header: &Header{
			EmailMessageHeader: emailMessage.Header,
		},
		ContentHeader: &ContentHeader{},
	}
	return
}

//Parse 解析邮件
func (instance *Email) Parse() (err error) {

	if err = instance.ParseEmail(); err != nil {
		return
	}

	if err = instance.ParseText(); err != nil {
		return
	}

	return
}

//ParseEmail 解析邮件内容
func (instance *Email) ParseEmail() (err error) {
	bodyContentType := instance.EmailMessage.Header["Content-Type"][0]
	if strings.Contains(bodyContentType, "multipart/") {
		ok := instance.setBoundary(bodyContentType) // 设置初始化bounddary
		if !ok {
			return errors.New("not found first boundary")
		}
		err = instance.multipart()
		return
	}
	return errors.New("not found ParseEmail instance")
}

// multipart 解析multipart/* 类型body
func (instance *Email) multipart() (err error) {
	bfReader := bufio.NewReader(instance.EmailMessage.Body)

	// 读取body部分头部
	for {
		// 一直读到最后一个割符处，并且记录 boundary
		byteArr, err := bfReader.ReadBytes('\n')
		if err != nil {
			break
		}

		if bytes.Contains(byteArr, instance.Boundary) {
			err = instance.readContentHeader(bfReader)
			if instance.ContentHeader.Encoding != "" || err != nil { // 读取到Content-Transfer-Encoding 则认为到了最后一个分割符号处（由于不同银行（深圳农村商业银行）邮件不规范，这种处理方式目前兼容性最强）
				break
			}
		}
	}

	// 读取内容主体
	bodyByteArr := make([][]byte, 0)
	for {
		lineArr, err := bfReader.ReadBytes('\n')                      // 读取内容
		if err != nil || bytes.Contains(lineArr, instance.Boundary) { // 读取出错，或者边界字符，则结束循环
			break
		}
		bodyByteArr = append(bodyByteArr, lineArr)
	}
	var sep []byte
	bodyBytes := bytes.Join(bodyByteArr, sep) // 链接拼成body
	instance.Content = string(bodyBytes)

	return
}

// setBoundary 设置分割符
func (instance *Email) setBoundary(text string) (ok bool) {
	reg := regexp.MustCompile(`boundary="(.+)"`) // 由于不同邮件，格式有些区别，此处采用正则匹配，方便最大化兼容
	matches := reg.FindStringSubmatch(text)
	if len(matches) > 1 {
		instance.Boundary = []byte(matches[1])
		ok = true
	}
	return
}

//readContentHeader 读取body中的header
func (instance *Email) readContentHeader(bf *bufio.Reader) (err error) {
	bodyHeader := instance.ContentHeader
	for {
		// 读取内容头部
		header, err := bf.ReadString('\n')
		if strings.Contains(header, "Content-Transfer-Encoding") {
			err = bodyHeader.setEncoding(header)
			continue
		}
		ok := instance.setBoundary(header) // 测试更新boundary
		if ok {
			continue
		}
		// Content-Type 和charset 可能在一行或者两行，需要放置最后判断，if中不能continue

		if strings.Contains(header, "Content-Type") {
			err = bodyHeader.setContentType(header)
		}
		if strings.Contains(header, "charset") {
			err = bodyHeader.setCharset(header)
		}

		if header == "\n" || err != nil {
			break
		}
	}
	return
}

//ParseText 提取html中text
func (instance *Email) ParseText() (err error) {
	instance.DecodeContent()
	instance.ConvertCharset()
	reader := strings.NewReader(instance.Content)
	doc, err := goquery.NewDocumentFromReader(reader)
	text := doc.Find("body").Text()
	reg := regexp.MustCompile(`\s*(\S+)\s*`)
	text = reg.ReplaceAllString(text, `$1`)
	text = ReplaceUtf8Space(text, " ")
	text = strings.TrimSpace(text)
	instance.Text = text
	return
}

// DecodeContent 解密邮件body部分
func (instance *Email) DecodeContent() (err error) {
	if instance.ContentHeader.Encoding == "quoted-printable" {
		instance.Content, err = DecodeQuotedprintable(instance.Content)
		return
	}
	if instance.ContentHeader.Encoding == "base64" {
		instance.Content, err = DecodeBase64(instance.Content)
		return
	}

	return
}

// ConvertCharset 转换字符编码
func (instance *Email) ConvertCharset() (err error) {
	if strings.ToLower(instance.ContentHeader.Charset) == "gb2312" || strings.ToLower(instance.ContentHeader.Charset) == "gbk" {
		instance.Content, err = GBKToUtf8(instance.Content)
	}
	return nil
}
