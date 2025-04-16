package tools

import "github.com/programmierigel/pwmanager/manager"

func RemoveMapFromMap(mapToFind map[string]map[string]manager.Secret, element string) map[string]map[string]manager.Secret {
	newMap := make(map[string]map[string]manager.Secret)
	for key, value := range mapToFind{
		if key == element{
			continue
		}
		newMap[key] = value
	}

	return newMap

}
