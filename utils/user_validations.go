package utils

import (
	"net"
	"regexp"
	"strings"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func IsEmailValid(e string) (bool, string) {
	if len(e) < 3 && len(e) > 254 {
		return false, "invalid length"
	}
	if !emailRegex.MatchString(e) {
		return false, "invalid regex"
	}
	parts := strings.Split(e, "@")
	mx, err := net.LookupMX(parts[1])
	if err != nil || len(mx) == 0 {
		return false, "domain does not exist"
	}
	return true, ""
}

func IsUserNameValid(e string) (bool, string) {
	if len(e) < 3 {
		return false, "Invalid length"
	}

	return true, ""
}