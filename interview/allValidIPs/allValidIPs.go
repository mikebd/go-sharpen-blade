package allValidIPs

// Given an input string of integers, return a slice of all valid
// IPv4 addresses that can be created by inserting "." in the input string

import (
	"fmt"
	"go-sharpen-blade/interview/util/validation"
	"math"
	"strings"
)

type segmentLengths [4]int

// allSegmentLengths returns all the possibly valid segment lengths for
// IPv4 addresses of maximum length.  This could be optimized to prebuild
// segment lengths for individual lengths of input strings.
var allSegmentLengths = func() []segmentLengths {
	result := make([]segmentLengths, 0, int(math.Pow(3, 4)))
	fmt.Println("cap", cap(result))

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				for l := 1; l <= 3; l++ {
					result = append(result, segmentLengths{i, j, k, l})
				}
			}
		}
	}
	fmt.Println("segment lengths:", len(result))
	return result
}

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

	for _, sl := range allSegmentLengths() {
		ip, ok := sl.build(s)
		if ok {
			result = append(result, ip)
		}
	}

	return result
}
