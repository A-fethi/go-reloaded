package main

import (
	"fmt"
	"strings"
	"unicode"
)

func GetApostIdx(s string) []int {
	var indices []int
	runes := []rune(s)
	for i, r := range runes {
		if r == '\'' && (i == 0 || i == len(runes)-1 || 
			(!unicode.IsLetter(runes[i-1]) && !unicode.IsLetter(runes[i+1]))) {
			indices = append(indices, i)
		}
	}
	return indices
}

func HandleSingleQuote(s string) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if len(line) <= 3 {
			continue
		}
		
		indices := GetApostIdx(line)
		if len(indices) == 0 {
			continue
		}
		
		var parts []string
		lastIndex := 0
		for j, idx := range indices {
			if j%2 == 0 {
				parts = append(parts, strings.TrimSpace(line[lastIndex:idx]))
				lastIndex = idx
			} else {
				parts = append(parts, "'"+strings.TrimSpace(line[lastIndex+1:idx])+"'")
				lastIndex = idx + 1
			}
		}
		if len(indices)%2 == 0 {
			parts = append(parts, strings.TrimSpace(line[lastIndex:]))
		}
		
		lines[i] = strings.Join(parts, " ")
	}
	return strings.Join(lines, "\n")
}

// func Quotes(s []string) []string {
// 	count := 0
// 	for i, word := range s {
// 		if word == "'" && count == 0 && !IsAlpha(word) {
// 			count += 1
// 			s[i+1] = word + s[i+1]
// 			// s[i+1] = word + s[i+1]
// 			s = append(s[:i], s[i+1:]...)
// 		}
// 	}

// 	for i, word := range s {
// 		if word == "'" {
// 			s[i-1] = s[i-1] + word
// 			s = append(s[:i], s[i+1:]...)
// 			count = 0
// 		}
// 	}

// 	return s
// }

// func IsAlpha(s string) bool {
// 	for _, i := range s {
// 		if (i < 'a' || i > 'z') && (i < 'A' || i > 'Z') && (i < '0' || i > '9') {
// 			return false
// 		}
// 	}
// 	return true
// }

func main() {
	str := "having a ' great time witfgf ' ' sdfsdfh sdfhsdf ' tgrte  ' dgfdg '."
	fmt.Println(HandleSingleQuote(str))
}
