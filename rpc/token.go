package rpc

import "github.com/1046102779/common/types"

// 微信开放平台公众号托管
type WxRelayServer struct {
	GetOfficialAccountPlatformInfo func() (*types.OfficialAccountPlatform, error)
	StoreOfficialAccountInfo       func(*types.OfficialAccount) (err error)
	GetOfficialAccountInfo         func(string) (*types.OfficialAccount, error)
	RefreshComponentVerifyTicket   func(*types.ComponentVerifyTicket) (string, error)
}
