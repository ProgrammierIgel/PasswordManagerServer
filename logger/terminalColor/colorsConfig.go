package terminalcolor

var colors = []string{"Red", "Green", "Yellow", "Blue", "Purple", "Cyan", "Gray", "White"}
var colorsWithCodes = map[string]string{

	"Red":    "\033[31m",
	"Green":  "\033[32m",
	"Yellow": "\033[33m",
	"Blue":   "\033[34m",
	"Purple": "\033[35m",
	"Cyan":   "\033[36m",
	"Gray":   "\033[37m",
	"White":  "\033[97m",
}
var Reset = "\033[0m"
