package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	text := "I was sitting over there ,and then BAMM ! ! ! ! I was thinking ... You were right"

	formattedText := formatText(text)
	fmt.Println(formattedText)
	fmt.Println("I was sitting over there, and then BAMM!!!! I was thinking... You were right")
	fmt.Println(len("I was sitting over there, and then BAMM!!!! I was thinking... You were right"))
	Text := Punctuations(text)
	fmt.Println(Text)
	fmt.Println(len(Text))
}

func formatText(text string) string {
	// Regular expression to match punctuation groups like ..., !?, etc.
	reGroups := regexp.MustCompile(`(\.{3}|!{1,2}\?{0,1})`)

	// Replace groups of punctuation first
	formattedText := reGroups.ReplaceAllStringFunc(text, func(match string) string {
		return strings.TrimSpace(match)
	})

	// Regular expression to match individual punctuation
	reIndividual := regexp.MustCompile(`\s*([.,!?;:])\s*`)

	// Replace individual punctuation
	formattedText = reIndividual.ReplaceAllString(formattedText, "$1")

	return formattedText
}

func Punctuations(el string) string {
    result := []rune(el)
    i := 0

    for i < len(result) {
        if i > 0 && IsPunc(result[i]) {
            // Move punctuation left as long as there are spaces before it
            j := i
            for j > 0 && result[j-1] == ' ' {
                result[j-1], result[j] = result[j], result[j-1]
                j--
            }
        }
        i++
    }

    return strings.TrimSpace(string(result))
}

func IsPunc(el rune) bool {
	if el == '.' || el == ',' || el == ':' || el == '?' || el == '!' || el == ';' {
		return true
	}
	return false
}