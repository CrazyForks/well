package wg

type MobileVPN interface {
	Start()
}

var mvpn MobileVPN

func SetMobileVPN(vpn MobileVPN) {
	mvpn = vpn
}
