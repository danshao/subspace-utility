package utils

import (
	"time"
	"math/big"
	"strconv"
)

const SOFTETHER_DATE_FORMAT = "2006-01-02 15:04:05"
func parseSoftetherDate(value string) (*time.Time) {
	if t, err := time.Parse(SOFTETHER_DATE_FORMAT, value); nil == err {
		return &t
	}
	return nil
}

func parseBigInt(value string) (*big.Int) {
	integer := new(big.Int)
	if _, ok := integer.SetString(value, 10); ok {
		return integer
	}
	return big.NewInt(0)
}

func parseDecimal(value string) (uint64) {
	if integer, err := strconv.ParseUint(value, 10, 64); nil == err {
		return integer
	}
	return 0
}

func parseUInt(value string) (uint) {
	if res, err := strconv.ParseUint(value, 10, 32); nil == err {
		return uint(res)
	}
	return 0
}