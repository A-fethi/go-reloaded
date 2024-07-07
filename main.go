package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	goreloaded "goreloaded/ressources"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid Arguments")
	} else {
		content, _ := os.ReadFile(os.Args[1])
		str := string(content)
		//str = goreloaded.Punctuations(str)
		//fmt.Println(str)
		SplitedStr := strings.Fields(str)
		//fmt.Println(SplitedStr)
		goreloaded.Atoan(SplitedStr)
		for i := 0; i < len(SplitedStr); i++ {
			if SplitedStr[i] == "(hex)" {
				SplitedStr[i-1] = goreloaded.HexConv(SplitedStr[i-1])
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
			} else if SplitedStr[i] == "(bin)" {
				SplitedStr[i-1] = goreloaded.BinConv(SplitedStr[i-1])
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
			} else if SplitedStr[i] == "(up)" {
				SplitedStr[i-1] = strings.ToUpper(SplitedStr[i-1])
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
			} else if SplitedStr[i] == "(up," {
				if SplitedStr[i] == "(up," {
					if i+1 < len(SplitedStr) {
						nb, valid := IsNumeric(SplitedStr[i+1])
						if valid {
							end := i - nb
							if end < 0 {
								end = 0
							}
							for j := i - 1; j >= end; j-- {
								SplitedStr[j] = strings.ToUpper(SplitedStr[j])
							}
							SplitedStr = append(SplitedStr[:i], SplitedStr[i+2:]...)
							//i-- // Adjust i because we removed two elements
						}
					}
					// nb := 0
					// nb = IsNumeric(SplitedStr[i+1])
					// if i < nb {
					// 	nb = i
					// 	for j := i - 1; j >= nb-i; j-- {
					// 		SplitedStr[j] = strings.ToUpper(SplitedStr[j])
					// 	}
					// } else {
					// 	for j := i - 1; j >= i-nb; j-- {
					// 		SplitedStr[j] = strings.ToUpper(SplitedStr[j])
					// 	}
					// }
					// SplitedStr = append(SplitedStr[:i], SplitedStr[i+2:]...)
				} else if SplitedStr[i] == "(low)" {
					SplitedStr[i-1] = strings.ToLower(SplitedStr[i-1])
					SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
				} else if SplitedStr[i] == "(cap)" {
					SplitedStr[i-1] = goreloaded.Capitalize(SplitedStr[i-1])
					SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
				}
			}
			output := strings.Join(SplitedStr, " ")
			output = goreloaded.Punctuations(output)
			fmt.Println(output)
			//fmt.Println(SplitedStr)
			//fmt.Println(len(SplitedStr))
			// outputFile, _ := os.Create(os.Args[2])
			// defer outputFile.Close()
			// output := os.WriteFile(os.Args[2], content, 0o666)
			// if output != nil {
			// 	fmt.Println(output)
			// 	return
			// }
		}
	}
}

func IsNumeric(s string) (int, bool) {
	s = strings.TrimSpace(s)
	if len(s) == 0 || s[len(s)-1] != ')' {
		return 0, false
	}
	s = s[:len(s)-1] // Remove the closing parenthesis
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	}
	return n, true
}

// func IsNumeric(s string) int {
// 	var n int
// 	for _, i := range s {
// 		if i >= '0' && i <= '9' {
// 			n += int(i - '0')
// 		} else {
// 			continue
// 		}
// 	}
// 	return n
// }

// fmt.Println(splitedStr)

// output := strings.Join(splitedStr, " ")
// err = os.WriteFile(os.Args[2], []byte(output), 0644)
// if err != nil {
// 	fmt.Println("Error writing to file:", err)
// 	return
// }
