package terminalcolor

import "fmt"

func isElementInSlice(element string, slice []string) error {
	for _, elementInSlice := range slice {
		if element == elementInSlice {
			return nil
		}
	}

	return fmt.Errorf("element not included")

}
