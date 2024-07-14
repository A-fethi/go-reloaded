package goreloaded

import "strings"

func Quotes(s string) string {
	result := ""
	open := false
	var start int

	for i, char := range s {
		if char == '\'' && (i == 0 || s[i-1] == ' ') && (i == len(s)-1 || s[i+1] == ' ') {
			if !open {
				open = true
				start = i
			} else {
				// Close the quote
				content := strings.TrimSpace(s[start+1 : i])
				result += "'" + content + "'"
				open = false
			}
		} else if !open {
			result += string(char)
		}
	}

	// If there's an unclosed quote, add the rest of the string
	if open {
		result += s[start:]
	}

	return DeleteSpaces(result)
}

func DeleteSpaces(s string) string {
	myRune := []rune(s)
	for i := 0; i < len(myRune)-1; i++ {
		if myRune[i] == ' ' && myRune[i+1] == ' ' {
			myRune = append(myRune[:i], append([]rune{'\u0000'}, myRune[i+1:]...)...)
		}
	}
	return string(myRune)
}