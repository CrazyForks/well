package wg

type MobileVPN interface {
	Start()
}

var mvpn MobileVPN

func SetMobileVPN(vpn MobileVPN) {
	mvpn = vpn
}

var ListenAddr string

func StartWireGuard(params DeviceParams) (err error) {
	devLocker.Lock()
	defer devLocker.Unlock()
	return startWireGuard(params)
}

func StopWireGuard() {
	devLocker.Lock()
	defer devLocker.Unlock()
	if dev := wgBind.Device.Swap(nil); dev != nil {
		dev.Close()
	}
}
