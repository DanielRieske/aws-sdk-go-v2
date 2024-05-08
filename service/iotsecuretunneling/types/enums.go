// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ClientMode string

// Enum values for ClientMode
const (
	ClientModeSource      ClientMode = "SOURCE"
	ClientModeDestination ClientMode = "DESTINATION"
	ClientModeAll         ClientMode = "ALL"
)

// Values returns all known values for ClientMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ClientMode) Values() []ClientMode {
	return []ClientMode{
		"SOURCE",
		"DESTINATION",
		"ALL",
	}
}

type ConnectionStatus string

// Enum values for ConnectionStatus
const (
	ConnectionStatusConnected    ConnectionStatus = "CONNECTED"
	ConnectionStatusDisconnected ConnectionStatus = "DISCONNECTED"
)

// Values returns all known values for ConnectionStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (ConnectionStatus) Values() []ConnectionStatus {
	return []ConnectionStatus{
		"CONNECTED",
		"DISCONNECTED",
	}
}

type TunnelStatus string

// Enum values for TunnelStatus
const (
	TunnelStatusOpen   TunnelStatus = "OPEN"
	TunnelStatusClosed TunnelStatus = "CLOSED"
)

// Values returns all known values for TunnelStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client.
//
// The ordering of this slice is not guaranteed to be stable across updates.
func (TunnelStatus) Values() []TunnelStatus {
	return []TunnelStatus{
		"OPEN",
		"CLOSED",
	}
}
