package types

// SaaS平台短信充值信息
type SmsRechargeInfo struct {
	Money             int64  `json:"money"`
	TradeNo           string `json:"trade_no"`
	Openid            string `json:"openid"`
	Title             string `json:"title"`
	OfficialAccountId int64  `json:"official_account_id"`
}

// 微信公众号JSAPI支付参数
type WechatJSAPIParamInfo struct {
	AppId     string `json:"app_id"`
	TimeStamp string `json:"timestamp"`
	NonceStr  string `json:"nonce_str"`
	Package   string `json:"package"`
	SignType  string `json:"sign_type"`
	PaySign   string `json:"pay_sign"`
}

// 商品支付参数信息
type ProductPayInfo struct {
	Money             int64  `json:"money"`
	TradeNo           string `json:"trade_no"`
	Title             string `json:"title"`
	OfficialAccountId int64  `json:"official_account_id"`
	TransactionId     string `json:"transaction_id"`
}
