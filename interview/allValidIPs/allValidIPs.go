package allValidIPs

// Given an input string of integers, return a slice of all valid
// IPv4 addresses that can be created by inserting "." in the input string

import (
	"go-sharpen-blade/interview/util/validation"
	"strings"
)

type segmentLengths [4]int

func (sl segmentLengths) build(s string) (string, bool) {
	if !sl.valid(s) {
		return "", false
	}

	segments := make([]string, 0, 4)
	start := 0
	for _, l := range sl {
		segment := s[start : start+l]
		if !validation.IsValidIPv4Segment(segment) {
			return "", false
		}
		segments = append(segments, segment)
		start += l
	}

	return strings.Join(segments, "."), true
}

// valid returns true if the sum of the segment lengths consume the whole string s
// and the individual segment lengths are valid for an IPv4
func (sl segmentLengths) valid(s string) bool {
	sum := 0

	for _, l := range sl {
		if l < 1 || l > 4 {
			return false
		}

		sum += l
	}

	return sum == len(s)
}

func allValidIPs(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		var result []string
		return result
	}

	result := make([]string, 0)

	sl := segmentLengths{1, 1, 1, 1}
	ip, ok := sl.build(s)
	if ok {
		result = append(result, ip)
	}

	return result
}
