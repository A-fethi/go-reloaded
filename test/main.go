package main

import "fmt"

// func Quotes(s string) string {
// 	myRune := []rune(s)
// 	for i, j := 0, len(myRune)-1; i < j; i, j = i+1, j-1 {
// 		if myRune[i] == '\'' && myRune[i+1] == ' ' {
// 			myRune = append(myRune[:i+1], myRune[i+2:]...)
// 		}
// 		if j < len(myRune)-1 && myRune[j] == '\'' && myRune[j-1] == ' ' {
// 			myRune = append(myRune[:j-1], myRune[j:]...)
// 		}
// 	}
// 	return string(myRune)
// }

func Quotes(s string) string {
	myRune := []rune(s)
	result := []rune{}
	open := false
	lastWasQuote := false

	for i := 0; i < len(myRune); i++ {
		char := myRune[i]

		if char == '\'' {
			if !open {
				result = append(result, char)
				open = true
			} else {
				if lastWasQuote {
					result = append(result, char)
					open = false
				} else {
					result = append(result, char)
					open = false
				}
			}
			lastWasQuote = true
		} else if char == ' ' {
			if !open || (open && !lastWasQuote) {
				result = append(result, char)
			}
			lastWasQuote = false
		} else {
			result = append(result, char)
			lastWasQuote = false
		}
	}

	return string(result)
}

func main() {
	s := "'   ' '   ' '   ' '   '"
	fmt.Println(Quotes(s))
}
