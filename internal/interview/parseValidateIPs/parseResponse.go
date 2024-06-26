package parseValidateIPs

import (
	"encoding/json"
	"go-sharpen-blade/internal/interview/util/validation"
	"slices"
	"strings"
)

/*
Example response:

	{
	    "timestamp": 1710974865,
	    "message": "Here are the IPs: 123.12.32.1 and this also 10.100.12.3 256.32.23.1"
	}
*/
type response struct {
	Timestamp int64
	Message   string
}

func extractValidIPs(responseBody string) []string {
	r := response{}
	err := json.Unmarshal([]byte(responseBody), &r)
	if err != nil {
		return []string{}
	}

	colonPos := strings.Index(r.Message, ":")
	var ips string
	if colonPos == -1 {
		ips = r.Message
	} else {
		ips = r.Message[colonPos+1:]
	}

	return slices.DeleteFunc(strings.Fields(ips), func(ip string) bool {
		return !validation.IsValidIPv4(ip)
	})
}
