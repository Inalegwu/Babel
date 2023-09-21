package utils

func Find[T comparable](array []T, check T) bool {
	var found bool
	for i := 0; i < len(array); i++ {
		val := array[i]
		if val == check {
			found = true
		}
	}
	return found
}
