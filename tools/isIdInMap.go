package tools

import "github.com/programmierigel/pwmanager/manager"

func IsIDInMap(id string, Map map[string]manager.Token) bool {
	for currentID := range Map {
		if currentID == id {
			return true
		}
	}
	return false
}
