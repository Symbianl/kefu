package models

type Order struct{
	OrderNo     string `orm:"pk"`
	BoardId     string `orm:"null"`
	ChanelId   int64 `orm:"0"`
	OpenId      string `orm:"null"`
	ProductId   int64  `orm:"0"`
	UserFee    int64 `orm:"0"` //用户实付
	Fee         int64  `orm:"0"` //店家实收
	Coins       int64  `orm:"0"`
	Discount  int64  	      //账户支付余额记录本次交易免费额度
	ShowFee     float64
	ShowCoins   float64
	Date        int64 `orm:"0"`
	ShowDate    string
	PrepayId    string `orm:"null"`
	State       int64  `orm:"0"`
	IsOnline    int64  `orm:"1"` //1 true, 0 false(offline), -1 privatepay used
	PayKey      int64  `orm:"null"`
	CltOrderNo  string
	IsOdd       bool   //用于显示
	PayLink     string //构造的带预支付码的直接支付链接
	OwnerOpenId string

	//微信订单相关信息
	WxBankType     string `orm:"null"`
	WxFeeType      string `orm:"null"`
	WxTransationId string `orm:"null"`
	WxIsSubscribe  string `orm:"null"`
	WxCashFee      int64  `orm:"null"`
	WxTotalFee     int64  `orm:"null"`
	WxTradeType    string `orm:"null"`
	Transfered     int64
	TransferedId   string
	BossOpenid     string
	BossTrueName string
	//支付宝订单相关信息
	AliBuyerId     string
	AliBuyerEmail  string
	AliSellerId    string
	//辅助信息
	//scan pay sn
	ScanPayCode    string
}