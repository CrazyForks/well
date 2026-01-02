package cmd

import (
	"encoding/json"

	"remoon.net/well/wg"
)

type MobileVPN interface {
	Start()
}

func SetMobileVPN(vpn MobileVPN) {
	wg.SetMobileVPN(vpn)
}

func StartWireGuard(fd int) error {
	err := wg.StartWireGuard(wg.DeviceParams{FD: fd, Routed: true})
	return err
}

func StopWireGuard() {
	wg.StopWireGuard()
}

func GetRoutes() string {
	routes := wg.GetRoutes()
	s, _ := json.Marshal(routes)
	return string(s)
}
