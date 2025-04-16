package tools

import "github.com/programmierigel/pwmanager/manager"

func IsElementInMap(element string, mapInput map[string]manager.Secret) bool {
	for name := range mapInput {
		if element == name {
			return true
		}
	}

	return false

}
