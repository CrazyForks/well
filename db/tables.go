package db

const (
	TablePeers   = "peers"
	TableICEs    = "ices"
	TableLinkers = "linkers"

	TableHookJS = "_hookjs"
)

const (
	TransportModeNoWS       = "NoWebSocket"
	TransportModeNoWebRTC   = "NoWebRTC"
	TransportModeWSRedirect = "EnableWebSocketRedirect"
)
