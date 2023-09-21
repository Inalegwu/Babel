package utils

func Find[T comparable](array []T, check T) T {
	var curr T
	for i := 0; i < len(array); i++ {
		val := array[i]
		if val == check {
			curr = val
		}
	}
	return curr
}
