package tools

func IsElementInMap(element string, mapInput map[string]string) bool {
	for name := range mapInput {
		if element == name {
			return true
		}
	}

	return false

}
