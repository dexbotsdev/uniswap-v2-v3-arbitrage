package data

import "strings"

var Blacklist = [...]string{
	"0xd233d1f6fd11640081abb8db125f722b5dc729dc",
	"0x7cd0378010711cb72a6ca35f09d5da2094061d9c",
}

func IsBlacklisted(address string) bool {
	addressLower := strings.ToLower(address)
	for _, blacklisted := range Blacklist {
		if strings.ToLower(blacklisted) == addressLower {
			return true
		}
	}
	return false
}
