package cmd

import "remoon.net/well/wg"

type MobileVPN = wg.MobileVPN

func SetMobileVPN(vpn MobileVPN) {
	wg.SetMobileVPN(vpn)
}
