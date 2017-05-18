package utils

import (
	"gitlab.ecoworkinc.com/Subspace/subspace-utility/subspace/model"
	"regexp"
)

// UserList and UserGet shared keys
const USER_NAME = "User Name"
const FULL_NAME = "Full Name"
const DESCRIPTION = "Description"
const EXPIRATION_DATE = "Expiration Date"

// UserList and UserGet alias keys
const AUTH_TYPE_IN_USER_LIST = "Auth Method"
const AUTH_TYPE_IN_USER_GET = "Auth Type"

const NUMBER_OF_LOGINS_IN_USER_LIST = "Num Logins"
const NUMBER_OF_LOGINS_IN_USER_GET = "Number of Logins"

// UserList keys
const GROUP_NAME = "Group Name"
const LAST_LOGIN = "Last Login"
const TRANSFER_BYTES = "Transfer Bytes"
const TRANSFER_PACKETS = "Transfer Packets"

// UserGet keys
const OUTGOING_UNICAST_PACKETS = "Outgoing Unicast Packets"
const OUTGOING_UNICAST_TOTAL_SIZE = "Outgoing Unicast Total Size"
const OUTGOING_BROADCAST_PACKETS = "Outgoing Broadcast Packets"
const OUTGOING_BROADCAST_TOTAL_SIZE = "Outgoing Broadcast Total Size"
const INCOMING_UNICAST_PACKETS = "Incoming Unicast Packets"
const INCOMING_UNICAST_TOTAL_SIZE = "Incoming Unicast Total Size"
const INCOMING_BROADCAST_PACKETS = "Incoming Broadcast Packets"
const INCOMING_BROADCAST_TOTAL_SIZE = "Incoming Broadcast Total Size"
const CREATED_ON = "Created on"
const UPDATED_ON = "Updated on"

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