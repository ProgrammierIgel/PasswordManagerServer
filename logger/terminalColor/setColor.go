package terminalcolor

import (
	"fmt"
)

func SetColor(s string, code string) string {
	err := isElementInSlice(code, colors)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s%s%s", colorsWithCodes[code], s, Reset)
}
