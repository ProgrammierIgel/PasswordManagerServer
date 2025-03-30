package tools

import "github.com/programmierigel/pwmanager/manager"

func RemovePasswordFromMap(mapToFind map[string]manager.Password, element string) map[string]manager.Password {
	newMap := make(map[string]manager.Password)
	for key, value := range mapToFind {
		if key == element {
			continue
		}
		newMap[key] = value
	}

	return newMap

}
