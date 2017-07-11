package validator

import (
	"regexp"
	"net"
	"fmt"
	"bytes"
	valid "github.com/asaskevich/govalidator"
)

var (
	REGEX_HOST = regexp.MustCompile(`^((([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.)+([a-zA-Z]{2,61})$`)
	REGEX_PRE_SHARED_KEY = regexp.MustCompile(`^[a-zA-Z0-9]{1,125}$`)
	REGEX_PRE_SHARED_KEY_RECOMMEND_BY_SOFTETHER = regexp.MustCompile(`^[a-zA-Z0-9]{1,9}$`)
	REGEX_UUID_V4 = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`)
	REGEX_EMAIL = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	LOWERCASE = "a-z"
	UPPERCASE = "A-Z"
	DIGIT = "0-9"
	SPECIAL_CHARACTER = " !\"#$%&'()*+,-./:;<=>?@\\[\\]^_`{|}~"
)

func IsValidHost(host string) bool {
	return REGEX_HOST.MatchString(host)
}

func IsValidIp(ip string) bool {
	return nil != net.ParseIP(ip)
}

func IsIpV4String(ip string) bool {
	result := net.ParseIP(ip)
	return nil != result && nil != result.To4()
}

func IsIpV6String(ip string) bool {
	result := net.ParseIP(ip)
	return nil != result && nil == result.To4() && nil != result.To16()
}

func IsValidPreSharedKey(preSharedKey string) bool {
	return REGEX_PRE_SHARED_KEY.MatchString(preSharedKey)
}

func IsValidPreSharedKeyRecommandBySoftether(preSharedKey string) bool {
	return REGEX_PRE_SHARED_KEY_RECOMMEND_BY_SOFTETHER.MatchString(preSharedKey)
}

func IsValidUuidV4(uuid string) bool {
	//return REGEX_UUID_V4.MatchString(uuid)
	return "" != uuid
}

func IsValidEmail(email string) bool {
	return valid.IsEmail(email)
}

func IsStringInArray(item string, arr []string) bool {
	for _, s := range arr {
		if s == item {
			return true
		}
	}
	return false
}

// Golang is NOT support before text matching
// See https://github.com/google/re2/wiki/Syntax
func IsValidPassword(password string) bool {
	if AtLeastOne(password, LOWERCASE) && AtLeastOne(password, UPPERCASE) && AtLeastOne(password, DIGIT) &&
					OnlyContains(password, 8, LOWERCASE, UPPERCASE, DIGIT, SPECIAL_CHARACTER) {
		return true
	}
	return false
}

func AtLeastOne(search, allowCharacters string) bool {
	match, _ := regexp.MatchString(fmt.Sprintf("[%s]+", allowCharacters), search)
	return match
}

func OnlyContains(search string, length int, allowPatterns ...string) bool {
	var buffer bytes.Buffer
	for _, pattern := range allowPatterns {
		buffer.WriteString(pattern)
	}
	match, _ := regexp.MatchString(fmt.Sprintf("^[%s]{%d,}$", buffer.String(), length), search)
	return match
}