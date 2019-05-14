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

func (s *sliceUtil) sliceStringCompare(s1, s2 []string) bool {
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
