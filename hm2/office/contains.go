package office

func Con[T comparable](array []T, word T) bool {
	for _, j := range array {
		if j == word {
			return true
		}
	}
	return false
}
