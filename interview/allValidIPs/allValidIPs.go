package allValidIPs

// Given an input string of integers, return a slice of all valid
// IPv4 addresses that can be created by inserting 3 x "." in the input string

import (
	"go-sharpen-blade/interview/util/validation"
	"strings"
)

const (
	// IPv4 has 4 segments: "seg1.seg2.seg3.seg4"
	ipV4SegmentCount = 4

	minSegmentLength = 1
	maxSegmentLength = 3

	minIPv4Digits   = len("1111")
	maxIPv4Digits   = len("123123123123")
	countIPv4Digits = maxIPv4Digits - minIPv4Digits + 1
)

type segmentLengths [ipV4SegmentCount]int

// allSegmentLengths returns all the possibly valid segment lengths for
// IPv4 addresses of various lengths.
var allSegmentLengths = func() map[int][]segmentLengths {
	result := make(map[int][]segmentLengths, countIPv4Digits)

	for length := minIPv4Digits; length <= maxIPv4Digits; length++ {
		for i := minSegmentLength; i <= maxSegmentLength; i++ {
			for j := minSegmentLength; j <= maxSegmentLength; j++ {
				for k := minSegmentLength; k <= maxSegmentLength; k++ {
					for l := minSegmentLength; l <= maxSegmentLength; l++ {
						if length == i+j+k+l {
							result[length] = append(result[length], segmentLengths{i, j, k, l})
						}
					}
				}
			}
		}
	}

	return result
}()

func (sl segmentLengths) build(s string) (string, bool) {
	if !sl.valid(s) {
		return "", false
	}

	segments := make([]string, 0, ipV4SegmentCount)
	start := 0
	for _, length := range sl {
		segment := s[start : start+length]
		if !validation.IsValidIPv4Segment(segment) {
			return "", false
		}
		segments = append(segments, segment)
		start += length
	}

	return strings.Join(segments, "."), true
}

// valid returns true if the sum of the segment lengths consume the whole string s
// and the individual segment lengths are valid for an IPv4
func (sl segmentLengths) valid(s string) bool {
	sum := 0

	for _, length := range sl {
		if length < minSegmentLength || length > maxSegmentLength {
			return false
		}

		sum += length
	}

	return sum == len(s)
}

func allValidIPs(s string) []string {
	if len(s) < minIPv4Digits || len(s) > maxIPv4Digits {
		var result []string
		return result
	}

	result := make([]string, 0, len(allSegmentLengths[len(s)]))

	for _, sl := range allSegmentLengths[len(s)] {
		ip, ok := sl.build(s)
		if ok {
			result = append(result, ip)
		}
	}

	return result
}
