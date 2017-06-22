package utils

import (
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/model"
	"regexp"
)

// UserList, UserGet, SessionList, SessionGet shared keys
const USER_NAME = "User Name"

// UserList, UserGet shared keys
const FULL_NAME = "Full Name"
const DESCRIPTION = "Description"
const EXPIRATION_DATE = "Expiration Date"

// UserList, SessionList shared keys
const TRANSFER_BYTES = "Transfer Bytes"
const TRANSFER_PACKETS = "Transfer Packets"

// UserGet, SessionGet shared keys
const OUTGOING_UNICAST_PACKETS = "Outgoing Unicast Packets"
const OUTGOING_UNICAST_TOTAL_SIZE = "Outgoing Unicast Total Size"
const OUTGOING_BROADCAST_PACKETS = "Outgoing Broadcast Packets"
const OUTGOING_BROADCAST_TOTAL_SIZE = "Outgoing Broadcast Total Size"
const INCOMING_UNICAST_PACKETS = "Incoming Unicast Packets"
const INCOMING_UNICAST_TOTAL_SIZE = "Incoming Unicast Total Size"
const INCOMING_BROADCAST_PACKETS = "Incoming Broadcast Packets"
const INCOMING_BROADCAST_TOTAL_SIZE = "Incoming Broadcast Total Size"

// UserList and UserGet alias keys
const AUTH_TYPE_IN_USER_LIST = "Auth Method"
const AUTH_TYPE_IN_USER_GET = "Auth Type"

const NUMBER_OF_LOGINS_IN_USER_LIST = "Num Logins"
const NUMBER_OF_LOGINS_IN_USER_GET = "Number of Logins"

// UserList keys
const GROUP_NAME = "Group Name"
const LAST_LOGIN = "Last Login"

// UserGet keys
const CREATED_ON = "Created on"
const UPDATED_ON = "Updated on"

// SessionList and SessionGet shared keys
const SESSION_NAME = "Session Name"
const VLAN_ID = "VLAN ID"

// SessionList keys
const LOCATION = "Location"
const SOURCE_HOST_NAME = "Source Host Name"
const TCP_CONNECTIONS = "TCP Connections"

// SessionGet keys
const CLIENT_IP_ADDRESS = "Client IP Address"
const CLIENT_HOST_NAME = "Client Host Name"
const USER_NAME_AUTHENTICATION = "User Name (Authentication)"
const USER_NAME_DATABASE = "User Name (Database)"
const SERVER_PRODUCT_NAME = "Server Product Name"
const SERVER_VERSION = "Server Version"
const SERVER_BUILD = "Server Build"
const CONNECTION_STARTED_AT = "Connection Started at"
const FIRST_SESSION_HAS_BEEN_ESTABLISHED_SINCE = "First Session has been Established since"
const CURRENT_SESSION_HAS_BEEN_ESTABLISHED_SINCE = "Current Session has been Established since"
const HALF_DUPLEX_TCP_CONNECTION_MODE = "Half Duplex TCP Connection Mode"
const VOIP_QOS_FUNCTION = "VoIP / QoS Function"
const NUMBER_OF_TCP_CONNECTIONS = "Number of TCP Connections"
const MAXIMUM_NUMBER_OF_TCP_CONNECTIONS = "Maximum Number of TCP Connections"
const ENCRYPTION = "Encryption"
const USE_OF_COMPRESSION = "Use of Compression"
const PHYSICAL_UNDERLAY_PROTOCOL = "Physical Underlay Protocol"
const UDP_ACCELERATION_IS_SUPPORTED = "UDP Acceleration is Supported"
const UDP_ACCELERATION_IS_ACTIVE = "UDP Acceleration is Active"
const CONNECTION_NAME = "Connection Name"
const SESSION_KEY_160_BIT = "Session Key (160 bit)"
const BRIDGE_ROUTER_MODE = "Bridge / Router Mode"
const MONITORING_MODE = "Monitoring Mode"
const CLIENT_PRODUCT_NAME_REPORTED = "Client Product Name (Reported)"
const CLIENT_VERSION_REPORTED = "Client Version (Reported)"
const CLIENT_BUILD_REPORTED = "Client Build (Reported)"
const CLIENT_OS_NAME_REPORTED = "Client OS Name (Reported)"
const CLIENT_OS_VERSION_REPORTED = "Client OS Version (Reported)"
const CLIENT_OS_PRODUCT_ID_REPORTED = "Client OS Product ID (Reported)"
const CLIENT_HOST_NAME_REPORTED = "Client Host Name (Reported)"
const CLIENT_IP_ADDRESS_REPORTED = "Client IP Address  (Reported)"
const CLIENT_PORT_REPORTED = "Client Port (Reported)"
const SERVER_HOST_NAME_REPORTED = "Server Host Name (Reported)"
const SERVER_IP_ADDRESS_REPORTED = "Server IP Address (Reported)"
const SERVER_PORT_REPORTED = "Server Port (Reported)"
const OUTGOING_DATA_SIZE = "Outgoing Data Size"
const INCOMING_DATA_SIZE = "Incoming Data Size"

type DataError struct {
	problem string
}

var reFindUserId = regexp.MustCompile("([0-9]+)_[0-9]{10,}")

func ParseUserList(hub string, data map[string]string) (brief *model.SoftEtherUserBrief, err *DataError) {
	if _, ok := data[USER_NAME]; !ok {
		return nil, &DataError{"Session must has name."}
	}

	brief = &model.SoftEtherUserBrief{
		Hub:             hub,
		UserName:        data[USER_NAME],
		FullName:        data[FULL_NAME],
		GroupName:       data[GROUP_NAME],
		ExpirationDate:  data[EXPIRATION_DATE],
		Description:     data[DESCRIPTION],
		AuthType:        data[AUTH_TYPE_IN_USER_LIST],
		LastLogin:       parseSoftetherDate(data[LAST_LOGIN]),
		TransferBytes:   parseDecimal(data[TRANSFER_BYTES]),
		TransferPackets: parseDecimal(data[TRANSFER_PACKETS]),
		NumberOfLogins:  parseUInt(data[NUMBER_OF_LOGINS_IN_USER_LIST]),
	}

	return brief, nil
}

func ParseUserGet(hub string, data map[string]string) (detail *model.ProfileSnapshot, err *DataError) {
	if _, ok := data[USER_NAME]; !ok {
		return nil, &DataError{"Session must has name."}
	}

	var userId uint = 0
	userName := data[USER_NAME]
	matches := reFindUserId.FindAllStringSubmatch(userName, -1)
	if 1 == len(matches) && 2 == len(matches[0]) {
		userId = parseUInt(matches[0][1])
	}

	detail = &model.ProfileSnapshot{
		Hub:                        hub,
		UserName:                   data[USER_NAME],
		UserId:                     userId,
		FullName:                   data[FULL_NAME],
		ExpirationDate:             data[EXPIRATION_DATE],
		Description:                data[DESCRIPTION],
		AuthType:                   data[AUTH_TYPE_IN_USER_GET],
		NumberOfLogins:             parseUInt(data[NUMBER_OF_LOGINS_IN_USER_GET]),
		IncomingBroadcastPackets:   parseDecimal(data[INCOMING_BROADCAST_PACKETS]),
		IncomingBroadcastTotalSize: parseDecimal(data[INCOMING_BROADCAST_TOTAL_SIZE]),
		IncomingUnicastPackets:     parseDecimal(data[INCOMING_UNICAST_PACKETS]),
		IncomingUnicastTotalSize:   parseDecimal(data[INCOMING_UNICAST_TOTAL_SIZE]),
		OutgoingBroadcastPackets:   parseDecimal(data[OUTGOING_BROADCAST_PACKETS]),
		OutgoingBroadcastTotalSize: parseDecimal(data[OUTGOING_BROADCAST_TOTAL_SIZE]),
		OutgoingUnicastPackets:     parseDecimal(data[OUTGOING_UNICAST_PACKETS]),
		OutgoingUnicastTotalSize:   parseDecimal(data[OUTGOING_UNICAST_TOTAL_SIZE]),
		CreatedOn:                  parseSoftetherDate(data[CREATED_ON]),
		UpdatedOn:                  parseSoftetherDate(data[UPDATED_ON]),
	}

	return detail, nil
}

func ParseSessionList(hub string, data map[string]string) (brief *model.Session, err *DataError) {
	if _, ok := data[SESSION_NAME]; !ok {
		return nil, &DataError{"Session must has name."}
	}

	brief = &model.Session{
		Hub:                    hub,
		SessionName:            data[SESSION_NAME],
		VLanId:                 data[VLAN_ID],
		Location:               data[LOCATION],
		UserNameAuthentication: data[USER_NAME],
		SourceHostName:         data[SOURCE_HOST_NAME],
		TCPConnections:         data[TCP_CONNECTIONS],
		TransferBytes:          data[TRANSFER_BYTES],
		TransferPackets:        data[TRANSFER_PACKETS],
	}

	return brief, nil
}

func ParseSessionGet(hub string, data map[string]string) (detail *model.Session, err *DataError) {
	if _, ok := data[SESSION_NAME]; !ok {
		return nil, &DataError{"Session must has name."}
	}

	detail = &model.Session{
		Hub:                                   hub,
		BridgeRouterMode:                      data[BRIDGE_ROUTER_MODE],
		ClientBuildReported:                   data[CLIENT_BUILD_REPORTED],
		ClientHostName:                        data[CLIENT_HOST_NAME],
		ClientHostNameReported:                data[CLIENT_HOST_NAME_REPORTED],
		ClientIPAddress:                       data[CLIENT_IP_ADDRESS],
		ClientIPAddressReported:               data[CLIENT_IP_ADDRESS_REPORTED],
		ClientOSNameReported:                  data[CLIENT_OS_NAME_REPORTED],
		ClientOSProductIDReported:             data[CLIENT_OS_PRODUCT_ID_REPORTED],
		ClientOSVersionReported:               data[CLIENT_OS_VERSION_REPORTED],
		ClientPortReported:                    data[CLIENT_PORT_REPORTED],
		ClientProductNameReported:             data[CLIENT_PRODUCT_NAME_REPORTED],
		ClientVersionReported:                 data[CLIENT_VERSION_REPORTED],
		ConnectionName:                        data[CONNECTION_NAME],
		ConnectionStartedAt:                   data[CONNECTION_STARTED_AT],
		CurrentSessionHasBeenEstablishedSince: data[CURRENT_SESSION_HAS_BEEN_ESTABLISHED_SINCE],
		Encryption:                            data[ENCRYPTION],
		FirstSessionHasBeenEstablishedSince:   data[FIRST_SESSION_HAS_BEEN_ESTABLISHED_SINCE],
		HalfDuplexTCPConnectionMode:           data[HALF_DUPLEX_TCP_CONNECTION_MODE],
		IncomingBroadcastPackets:              data[INCOMING_BROADCAST_PACKETS],
		IncomingBroadcastTotalSize:            data[INCOMING_BROADCAST_TOTAL_SIZE],
		IncomingDataSize:                      data[INCOMING_DATA_SIZE],
		IncomingUnicastPackets:                data[INCOMING_UNICAST_PACKETS],
		IncomingUnicastTotalSize:              data[INCOMING_UNICAST_TOTAL_SIZE],
		MaximumNumberOfTCPConnections:         data[MAXIMUM_NUMBER_OF_TCP_CONNECTIONS],
		MonitoringMode:                        data[MONITORING_MODE],
		NumberOfTCPConnections:                data[NUMBER_OF_TCP_CONNECTIONS],
		OutgoingBroadcastPackets:              data[OUTGOING_BROADCAST_PACKETS],
		OutgoingBroadcastTotalSize:            data[OUTGOING_BROADCAST_TOTAL_SIZE],
		OutgoingDataSize:                      data[OUTGOING_DATA_SIZE],
		OutgoingUnicastPackets:                data[OUTGOING_UNICAST_PACKETS],
		OutgoingUnicastTotalSize:              data[OUTGOING_UNICAST_TOTAL_SIZE],
		PhysicalUnderlayProtocol:              data[PHYSICAL_UNDERLAY_PROTOCOL],
		ServerBuild:                           data[SERVER_BUILD],
		ServerHostNameReported:                data[SERVER_HOST_NAME_REPORTED],
		ServerIPAddressReported:               data[SERVER_IP_ADDRESS_REPORTED],
		ServerPortReported:                    data[SERVER_PORT_REPORTED],
		ServerProductName:                     data[SERVER_PRODUCT_NAME],
		ServerVersion:                         data[SERVER_VERSION],
		SessionKey160Bit:                      data[SESSION_KEY_160_BIT],
		SessionName:                           data[SESSION_NAME],
		UDPAccelerationIsActive:               data[UDP_ACCELERATION_IS_ACTIVE],
		UDPAccelerationIsSupported:            data[UDP_ACCELERATION_IS_SUPPORTED],
		UseOfCompression:                      data[USE_OF_COMPRESSION],
		UserNameAuthentication:                data[USER_NAME_AUTHENTICATION],
		UserNameDatabase:                      data[USER_NAME_DATABASE],
		VLanId:                                data[VLAN_ID],
		VoIPQoSFunction:                       data[VOIP_QOS_FUNCTION],
	}

	return detail, nil
}