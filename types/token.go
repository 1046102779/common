package types

// 微信开放平台
type OfficialAccountPlatform struct {
	Appid                 string `json:"appid"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
	EncodingAesKey        string `json:"encoding_aes_key"`
	AppSecret             string `json:"app_secret"`
	ComponentAccessToken  string `json:"component_access_token"`
	PreAuthCode           string `json:"pre_auth_code"`
	Token                 string `json:"token"`
}

// 公众号
type OfficialAccount struct {
	Appid                          string `json:"appid"`
	AuthorizerAccessToken          string `json:"authorizer_access_token"`
	AuthorizerAccessTokenExpiresIn int64  `json:"authorizer_access_token_expires_in"`
	AuthorizerRefreshToken         string `json:"authorizer_refresh_token"`
}

// 授权凭证
type ComponentVerifyTicket struct {
	TimeStamp         string `json:"timestamp"`
	Nonce             string `json:"nonce"`
	MsgSign           string `json:"msg_sign"`
	Bts               []byte `json:"bts"`
	AuthorizationCode string `json:"authorization_code"`
}
