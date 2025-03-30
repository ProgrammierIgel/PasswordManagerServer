package tools

func RemoveMapFromMap(mapToFind map[string]map[string]string, element string) map[string] map[string]string {
	newMap := make(map[string]map[string]string)
	for key, value := range mapToFind{
		if key == element{
			continue
		}
		newMap[key] = value
	}

	return newMap

}
