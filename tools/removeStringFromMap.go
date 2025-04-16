package tools

import "github.com/programmierigel/pwmanager/manager"

func RemoveStringFromMap(mapToFind map[string]manager.Secret, element string) map[string]manager.Secret {
	newMap := make(map[string]manager.Secret)
	for key, value := range mapToFind {
		if key == element {
			continue
		}
		newMap[key] = value
	}

	return newMap

}
