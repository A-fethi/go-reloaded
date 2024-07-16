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
				/*--------handling quotes--------*/
				if i < len(text)-1 {
					if string(text[i-1]) != " " && string(text[i-1]) != "\n" && string(text[i+1]) != " " {
						result += string(char)
					} else {
						/*--------handling singe quotes--------*/
						if count == 0  {
							count++
						} else {
							result = strings.TrimSuffix(result, " ")
							count--
						}
						result += string(char)
					}
				} else {
					/*--------handling singe quotes--------*/
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

	if count != 0 {
		panic("Error: missing closed quote")
	}
	return result
}
