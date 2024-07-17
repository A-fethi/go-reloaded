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
		result += string(unicode.ToUpper(myRune[i]))
	}
	return result
}
