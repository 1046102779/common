package rpc

import "github.com/1046102779/common/types"

type OfficialAccountServer struct {
	GetSmsRechargePayJsapiParams  func(*types.SmsRechargeInfo) (*types.WechatJSAPIParamInfo, error)
	GetSaleOrderPayNativeParams   func(*types.ProductPayInfo) (string, error)
	GetSmsRechargePayNativeParams func(*types.SmsRechargeInfo) (string, error)
	BindingCustomerAndWxInfo      func(string, int) error
	BindingUserAndWxInfo          func(string, int) error
	GetOfficialAccountId          func(int) (int, error)
	SendMessage                   func(int, int, int16) error
}
