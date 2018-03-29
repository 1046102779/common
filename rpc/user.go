package rpc

type UserServer struct {
	GetWechatBindingUserInfo func(int, int16) (string, error)
}
