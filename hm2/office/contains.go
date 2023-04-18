package office

func Con(array []string, word string) bool {
	for _, j := range array {
		if j == word {
			return true
		}
	}
	return false
}
