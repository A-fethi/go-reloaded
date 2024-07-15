package goreloaded

import "strings"

func Quotes(text string) string {
	singlequote := false
	result := ""
	for i, char := range text {
		if i == 0 {
			result += string(char)
			if string(char) == "'" {
				singlequote = true
			}
		} else {
			if char != '\'' && char != ' ' {
				result += string(char)
			} else if char == '\'' {
				/*--------handling quotes--------*/
				if i < len(text)-1 {
					if string(text[i-1]) != " " && string(text[i-1]) != "\n" && string(text[i+1]) != " " && string(text[i+1]) != "\n" {
						result += string(char)
					} else {
						/*--------handling singe quotes--------*/
						if !singlequote {
							singlequote = true
						} else {
							singlequote = false
							result = strings.TrimSuffix(result, " ")
						}
						result += string(char)
					}
				} else {
					/*--------handling singe quotes--------*/
					if singlequote {
						result = strings.TrimSuffix(result, " ")
						result += string(char)
					} else {
						result += string(char)
					}
				}
			} else if char == ' ' {
				if string(result[len(result)-1]) == "'" && singlequote {
					continue
				} else {
					result += string(char)
				}
			}
		}
	}

	return result
}
