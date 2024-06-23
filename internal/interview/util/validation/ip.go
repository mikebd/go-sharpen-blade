package validation

import (
	"strconv"
	"strings"
)

// IsValidIPv4 returns true if the given string is a valid IPv4 address.
// This is the interview version of: net.ParseIP(ip) != nil
func IsValidIPv4(ip string) bool {
	segments := strings.Split(ip, ".")

	if len(segments) != 4 {
		return false
	}

	for _, segment := range segments {
		if !IsValidIPv4Segment(segment) {
			return false
		}
	}

	return true
}

// IsValidIPv4Segment returns true if the given string is a valid IPv4 ip.
// Valid segments are: 0-255.  Multiple leading zeros are invalid.
func IsValidIPv4Segment(segment string) bool {
	if len(segment) == 0 || len(segment) > 3 {
		return false
	}

	if segment[0] == '0' && len(segment) > 1 {
		return false
	}

	n, err := strconv.Atoi(segment)
	if err != nil {
		return false
	}

	return n >= 0 && n <= 255
}
