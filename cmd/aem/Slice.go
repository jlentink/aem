package main

type SliceUtil struct{}

func (s *SliceUtil) InSliceInt64(slice []int64, needle int64) bool {
	for _, v := range slice {
		if v == needle {
			return true
		}
	}
	return false
}

func (s *SliceUtil) InSliceString(slice []string, needle string) bool {
	for _, v := range slice {
		if v == needle {
			return true
		}
	}
	return false
}
