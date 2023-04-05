package utills

import (
	"strings"
)

func MatchString(str string, substr string) bool {
	if strings.Contains(str, substr) {
		return true
	}
	return false
}
