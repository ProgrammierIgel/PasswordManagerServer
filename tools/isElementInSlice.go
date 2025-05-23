package tools

import "fmt"

func IsElementInSlice(element string, slice []string) error {
	for _, elementInSlice := range slice {
		if element == elementInSlice {
			return nil
		}
	}

	return fmt.Errorf("element not included")

}
