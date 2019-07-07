package recognizer

//Address 邮箱地址类型
type Address string

//Owner 邮件接收者
type Owner struct {
	Name    string
	Address string
}


//Emailer 邮件对象，地址和用户
type Emailer struct {
	Address string
	Name    string
	Issuer  string
}
