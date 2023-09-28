package probe

import (
	"os"
)

func CheckCross() (string, bool) {
	crossEnv := os.Getenv("CROSS_COMPILE")
	if crossEnv == "1" {
		return "", true
	} else if crossEnv != "" {
		return crossEnv, true
	}
	return "", false
}
