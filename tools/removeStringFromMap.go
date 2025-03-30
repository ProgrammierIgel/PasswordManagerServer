package tools

func RemoveStringFromMap(mapToFind map[string]string, element string) map[string]string {
	newMap := make(map[string]string)
	for key, value := range mapToFind {
		if key == element {
			continue
		}
		newMap[key] = value
	}

	return newMap

}
