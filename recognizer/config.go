package recognizer

const (
	// IssuerZhaoHang 招行机构
	IssuerZhaoHang = "zhaohang"
	//IssuerPuFa 浦发机构
	IssuerPuFa = "pufa"
	//IssuerZhongXin 中信机构
	IssuerZhongXin = "zhongxin"
	//IssuerPingAn 平安机构
	IssuerPingAn = "pingan"
	// IssuerSZNongCunShangYe 深圳农村商业银行
	IssuerSZNongCunShangYe = "sznongcunshangye"
)

//FromUsers 关注一下用户发送的邮件
var FromUsers = []Emailer{
	Emailer{
		Address:  "ccsvc@message.cmbchina.com",
		Name:   "招商银行信用卡",
		Issuer: IssuerZhaoHang,
	},
	Emailer{
		Address:  "estmtservice@eb.spdbccc.com.cn",
		Name:   "浦发银行信用卡",
		Issuer: IssuerPuFa,
	},
	Emailer{
		Address:  "citiccard@bill.citiccard.com",
		Name:   "中信银行信用卡",
		Issuer: IssuerZhongXin,
	},
	Emailer{
		Address:  "creditcard@service.pingan.com",
		Name:   "平安银行信用卡",
		Issuer: IssuerPingAn,
	},
	Emailer{
		Address:  "Bill_card@4001961200.com",
		Name:   "深圳农村商业银行",
		Issuer: IssuerSZNongCunShangYe,
	},
}


