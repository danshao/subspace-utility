package model

//Generated using JSON 2 Go struct tool: https://mholt.github.io/json-to-go/

type Session struct {
	// Additional data
	Hub         string `json:"hub"`

	// SessionList and //SessionGet shared column
	SessionName string `json:"sessionName"`
	VLanId      string `json:"vlanId"`

	// SessionList
	Location        string `json:"location"`
	SourceHostName  string `json:"sourceHostName"`
	TCPConnections  string `json:"tcpconnections"`
	TransferBytes   string `json:"transferBytes"`
	TransferPackets string `json:"transferPackets"`

	// SessionGet
	BridgeRouterMode                      string `json:"bridgeRouterMode"`
	ClientBuildReported                   string `json:"clientBuildReported"`
	ClientHostName                        string `json:"clientHostName"`
	ClientHostNameReported                string `json:"clientHostNameReported"`
	ClientIPAddress                       string `json:"clientIPAddress"`
	ClientIPAddressReported               string `json:"clientIPAddressReported"`
	ClientOSNameReported                  string `json:"clientOSNameReported"`
	ClientOSProductIDReported             string `json:"clientOSProductIDReported"`
	ClientOSVersionReported               string `json:"clientOSVersionReported"`
	ClientPortReported                    string `json:"clientPortReported"`
	ClientProductNameReported             string `json:"clientProductNameReported"`
	ClientVersionReported                 string `json:"clientVersionReported"`
	ConnectionName                        string `json:"connectionName"`
	ConnectionStartedAt                   string `json:"connectionStartedAt"`
	CurrentSessionHasBeenEstablishedSince string `json:"currentSessionHasBeenEstablishedSince"`
	Encryption                            string `json:"encryption"`
	FirstSessionHasBeenEstablishedSince   string `json:"firstSessionHasBeenEstablishedSince"`
	HalfDuplexTCPConnectionMode           string `json:"halfDuplexTCPConnectionMode"`
	IncomingBroadcastPackets              string `json:"incomingBroadcastPackets"`
	IncomingBroadcastTotalSize            string `json:"incomingBroadcastTotalSize"`
	IncomingDataSize                      string `json:"incomingDataSize"`
	IncomingUnicastPackets                string `json:"incomingUnicastPackets"`
	IncomingUnicastTotalSize              string `json:"incomingUnicastTotalSize"`
	MaximumNumberOfTCPConnections         string `json:"maximumNumberOfTCPConnections"`
	MonitoringMode                        string `json:"monitoringMode"`
	NumberOfTCPConnections                string `json:"numberOfTCPConnections"`
	OutgoingBroadcastPackets              string `json:"outgoingBroadcastPackets"`
	OutgoingBroadcastTotalSize            string `json:"outgoingBroadcastTotalSize"`
	OutgoingDataSize                      string `json:"outgoingDataSize"`
	OutgoingUnicastPackets                string `json:"outgoingUnicastPackets"`
	OutgoingUnicastTotalSize              string `json:"outgoingUnicastTotalSize"`
	PhysicalUnderlayProtocol              string `json:"physicalUnderlayProtocol"`
	ServerBuild                           string `json:"serverBuild"`
	ServerHostNameReported                string `json:"serverHostNameReported"`
	ServerIPAddressReported               string `json:"serverIPAddressReported"`
	ServerPortReported                    string `json:"serverPortReported"`
	ServerProductName                     string `json:"serverProductName"`
	ServerVersion                         string `json:"serverVersion"`
	SessionKey160Bit                      string `json:"sessionKey160Bit"`
	UDPAccelerationIsActive               string `json:"udpaccelerationIsActive"`
	UDPAccelerationIsSupported            string `json:"udpaccelerationIsSupported"`
	UseOfCompression                      string `json:"useOfCompression"`
	UserNameAuthentication                string `json:"userName"` //Change to userName
	UserNameDatabase                      string `json:"userNameDatabase"`
	VoIPQoSFunction                       string `json:"voIPQoSFunction"`
}
