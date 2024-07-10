package goreloaded

import "strings"

// fix punctuation position
func Punctuations(el string) string {
	result := []rune(el)

	for i := 0; i < len(result)-1; i++ {
		if (IsPunc(result[i+1]) && result[i] == ' ') || (IsPunc(result[i+1]) && result[i] == '') {
			result[i+1], result[i] = result[i], result[i+1]
		}
	}
	return strings.Join(strings.Fields(string(result)), " ")
}

func IsPunc(el rune) bool {
	if el == '.' || el == ',' || el == ':' || el == '?' || el == '!' || el == ';' {
		return true
	}
	return false
}