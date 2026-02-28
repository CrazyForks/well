package db

const (
	TablePeers   = "peers"
	TableICEs    = "ices"
	TableLinkers = "linkers"

	TableHookJS  = "_hookjs"
	TableExports = "exports"
)

const (
	TransportModeNoWS       = "NoWebSocket"
	TransportModeNoWebRTC   = "NoWebRTC"
	TransportModeWSRedirect = "EnableWebSocketRedirect"
)
