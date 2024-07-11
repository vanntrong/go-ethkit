package utils

func AssertInArrayInt(arr []int, value int) bool {
	for _, b := range arr {
		if b == value {
			return true
		}
	}
	return false
}
