package tools

import "github.com/programmierigel/pwmanager/manager"

func RemoveTokenFromMap(token string, Map map[string]manager.Token) map[string]manager.Token {
	newMap := make(map[string]manager.Token)
	for key, value := range Map {
		if key == token {
			continue
		}
		newMap[key] = value
	}

	return newMap

}
