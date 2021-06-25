package sliceutil

import "strings"

// StringCompare Compare string slices.
func StringCompare(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// StringInSliceEqualFold Check if string is in slice with equal fold
func StringInSliceEqualFold(token string, stringSlice []string) bool{
	for _, s := range stringSlice {
		if strings.EqualFold(token, s) {
			return true
		}
	}
	return false
}

