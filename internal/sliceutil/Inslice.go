package sliceutil

// InSliceInt64 checks if slice contains int64 value.
func InSliceInt64(slice []int64, needle int64) bool {
	for _, v := range slice {
		if v == needle {
			return true
		}
	}
	return false
}

// InSliceString checks if slice contains string value.
func InSliceString(slice []string, needle string) bool {
	for _, v := range slice {
		if v == needle {
			return true
		}
	}
	return false
}
