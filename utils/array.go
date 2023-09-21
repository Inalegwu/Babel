package utils

func Find[T comparable](array []T, check T) bool {
	var curr T
	var found bool
	for i := 0; i < len(array)-1; i++ {
		curr = array[i]
		if curr == check {
			found = true
		} else {
			// found = false
		}
	}
	return found
}
