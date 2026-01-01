package wg

type MobileVPN interface {
	Start(token string) // 因为 POST /api/ipc/device 需要管理员权限, 所以传递下 token
}

var mvpn MobileVPN

func SetMobileVPN(vpn MobileVPN) {
	mvpn = vpn
}
