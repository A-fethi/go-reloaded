package goreloaded

import (
	"strings"
)

func Quotes(text string) string {
	count := 0
	result := ""
	for i, char := range text {
		if i == 0 {
			result += string(char)
			if string(char) == "'" {
				count++
			}
		} else {
			if char != '\'' && char != ' ' {
				result += string(char)
			} else if char == '\'' {
				/*--------handling quotes in the middle of the word--------*/
				if i < len(text)-1 {
					if isChar(string(text[i-1])) && isChar(string(text[i+1])) {
						result += string(char)
					} else {
						/*--------handling opening quote--------*/
						if count == 0 {
							count++
						} else {
							/*--------handling closing quote--------*/
							result = strings.TrimSuffix(result, " ")
							count--
						}
						result += string(char)
					}
				} else {
					/*--------handling quotes at the end of the text--------*/
					if count == 1 {
						result = strings.TrimSuffix(result, " ")
						result += string(char)
						count--
					} else {
						result += string(char)
						count++
					}
				}
			} else if char == ' ' {
				if string(result[len(result)-1]) == "'" && count == 1 {
					continue
				} else {
					result += string(char)
				}
			}
		}
	}
	return result
}

func isChar(s string) bool {
	for i := range s {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= 'À' && s[i] <= 'ÿ') {
			return true
		}
	}
	return false
}
