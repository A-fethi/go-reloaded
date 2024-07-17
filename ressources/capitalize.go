package goreloaded

import (
	"unicode"
)

func Capitalize(s string) string {
	myRune := []rune(s)
	if string(myRune) == "" {
		return string(myRune)
	}
	var result string
	for i := 0; i < len(myRune); i++ {
		if i == 0 {
			result += string(unicode.ToUpper(myRune[i]))
		} else if i > 0 {
			result += string(unicode.ToLower(myRune[i]))
		}
	}
	return result
}
