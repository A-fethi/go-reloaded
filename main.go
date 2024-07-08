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
		// str = goreloaded.Punctuations(str)
		SplitedStr := strings.Split(str, " ")
		goreloaded.Atoan(SplitedStr)
		for i := 0; i <= len(SplitedStr)-1; i++ {
			if SplitedStr[i] == "(hex)" {
				SplitedStr[i-1] = goreloaded.HexConv(SplitedStr[i-1])
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
				i--
			} else if SplitedStr[i] == "(bin)" {
				SplitedStr[i-1] = goreloaded.BinConv(SplitedStr[i-1])
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
				i--
			} else if SplitedStr[i] == "(up)" {
				SplitedStr[i-1] = strings.ToUpper(SplitedStr[i-1])
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
				i--
			} else if SplitedStr[i] == "(low)" {
				SplitedStr[i-1] = strings.ToLower(SplitedStr[i-1])
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
				i--
			} else if SplitedStr[i] == "(cap)" {
				SplitedStr[i-1] = goreloaded.Capitalize(SplitedStr[i-1])
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
				i--
			} else if SplitedStr[i] == "(up," {
				SplitedStr, i = instancesProcess(SplitedStr, i)
			} else if SplitedStr[i] == "(low," {
				SplitedStr, i = instancesProcess(SplitedStr, i)
			} else if SplitedStr[i] == "(cap," {
				SplitedStr, i = instancesProcess(SplitedStr, i)
			}
		}
		output := strings.Join(SplitedStr, " ")
		output = goreloaded.Punctuations(output)
		fields := strings.Fields(output)
		output = strings.Join(fields, " ")
		fmt.Println(output)
		// fmt.Println(len(SplitedStr))
		// outputFile, _ := os.Create(os.Args[2])
		// defer outputFile.Close()
		// output := os.WriteFile(os.Args[2], content, 0o666)
		// if output != nil {
		// 	fmt.Println(output)
		// 	return
		// }
	}

	// fmt.Println(SplitedStr)
}

func isNumeric(s string) (int, bool) {
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

func instancesProcess(SplitedStr []string, i int) ([]string, int) {
	if SplitedStr[i] == "(up," {
		if i+1 < len(SplitedStr) {
			nb, valid := isNumeric(SplitedStr[i+1])
			if valid {
				end := i - nb
				if end < 0 {
					end = 0
				}
				for j := i - 1; j >= end; j-- {
					SplitedStr[j] = strings.ToUpper(SplitedStr[j])
				}
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+2:]...)
				i--
			}
		}
	} else if SplitedStr[i] == "(low," {
		if i+1 < len(SplitedStr) {
			nb, valid := isNumeric(SplitedStr[i+1])
			if valid {
				end := i - nb
				if end < 0 {
					end = 0
				}
				for j := i - 1; j >= end; j-- {
					SplitedStr[j] = strings.ToLower(SplitedStr[j])
				}
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+2:]...)
				i--
			}
		}
	} else if SplitedStr[i] == "(cap," {
		if i+1 < len(SplitedStr) {
			nb, valid := isNumeric(SplitedStr[i+1])
			if valid {
				end := i - nb
				if end < 0 {
					end = 0
				}
				for j := i - 1; j >= end; j-- {
					SplitedStr[j] = goreloaded.Capitalize(SplitedStr[j])
				}
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+2:]...)
				i--
			}
		}
	}
	return SplitedStr, i
}

// output := strings.Join(splitedStr, " ")
// err = os.WriteFile(os.Args[2], []byte(output), 0644)
// if err != nil {
// 	fmt.Println("Error writing to file:", err)
// 	return
// }
