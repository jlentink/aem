package main

type sliceUtil struct{}

func (s *sliceUtil) inSliceInt64(slice []int64, needle int64) bool {
	for _, v := range slice {
		if v == needle {
			return true
		}
	}
	return false
}

func (s *sliceUtil) inSliceString(slice []string, needle string) bool {
	for _, v := range slice {
		if v == needle {
			return true
		}
	}
	return false
}
