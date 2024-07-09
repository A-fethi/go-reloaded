package goreloaded

import (
	"strings"
)

func Capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + strings.ToLower(s[1:])
}
