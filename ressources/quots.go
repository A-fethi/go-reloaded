package goreloaded

// func Quotes(s string) string {
// 	myRune := []rune(s)
// 	for i, j := 0, len(myRune)-1; i < j; i, j = i+1, j-1 {
// 		if myRune[i] == '\'' && myRune[i+1] == ' ' {
// 			myRune = append(myRune[:i+1], myRune[i+2:]...)
// 		}
// 		// if myRune[i] == '\'' && myRune[i-1] == ' ' {
// 		// 	myRune = append(myRune[:i-1], myRune[i:]...)
// 		// }
// 		// if i > 0 && myRune[j] == '\'' && myRune[j-1] == ' ' {
// 		// 	myRune = append(myRune[:j-1], myRune[j:]...)
// 		// }
// 	}
// 	return string(myRune)
// }

// func Quotes(s string) string {
//     runes := []rune(s)
//     result := []rune{}
//     inQuote := false
//     quoteStart := -1
//     lastWordChar := -1

//     for i, char := range runes {
//         if char == '\'' {
//             if !inQuote && (lastWordChar == -1 || i-lastWordChar > 1) {
//                 // Opening quote: only if it's not inside a word
//                 inQuote = true
//                 quoteStart = len(result)
//                 result = append(result, char)
//             } else if inQuote {
//                 // Closing quote
//                 inQuote = false
//                 // Trim spaces at the beginning and end of the quote
//                 for quoteStart+1 < len(result) && result[quoteStart+1] == ' ' {
//                     result = append(result[:quoteStart+1], result[quoteStart+2:]...)
//                 }
//                 for len(result) > 0 && result[len(result)-1] == ' ' {
//                     result = result[:len(result)-1]
//                 }
//                 result = append(result, char)
//             } else {
//                 // Quote inside a word or unclosed quote
//                 result = append(result, char)
//             }
//         } else if inQuote && char == ' ' {
//             // Only add space if it's not at the start or end of the quote
//             if quoteStart+1 < len(result) && result[len(result)-1] != ' ' && i+1 < len(runes) && runes[i+1] != '\'' {
//                 result = append(result, char)
//             }
//         } else {
//             result = append(result, char)
//             if char != ' ' {
//                 lastWordChar = i
//             }
//         }
//     }

//     // If the quote wasn't closed, treat it as regular text
//     if inQuote {
//         result = append([]rune{'\''},  result...)
//     }

//     return string(result)
// }
func Quotes(s string) string {
    runes := []rune(s)
    result := []rune{}
    inQuote := false
    quoteStart := -1
    lastWordChar := -1
    consecutiveQuotes := 0

    for i, char := range runes {
        if char == '\'' {
            consecutiveQuotes++
            if consecutiveQuotes % 2 == 1 {
                if !inQuote && (lastWordChar == -1 || i-lastWordChar > 1) {
                    // Opening quote
                    inQuote = true
                    quoteStart = len(result)
                }
            } else {
                if inQuote {
                    // Closing quote
                    inQuote = false
                    // Trim spaces at the beginning and end of the quote
                    for quoteStart+1 < len(result) && result[quoteStart+1] == ' ' {
                        result = append(result[:quoteStart+1], result[quoteStart+2:]...)
                    }
                    for len(result) > 0 && result[len(result)-1] == ' ' {
                        result = result[:len(result)-1]
                    }
                }
            }
            result = append(result, char)
        } else {
            consecutiveQuotes = 0
            if inQuote && char == ' ' {
                // Only add space if it's not at the start or end of the quote
                if quoteStart+1 < len(result) && result[len(result)-1] != ' ' && i+1 < len(runes) && runes[i+1] != '\'' {
                    result = append(result, char)
                }
            } else {
                result = append(result, char)
                if char != ' ' {
                    lastWordChar = i
                }
            }
        }
    }

    return string(result)
}