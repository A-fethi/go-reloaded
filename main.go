package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	goreloaded "goreloaded/ressources"
)

func isNumeric(s string) (int, bool) {
	if len(s) == 0 || s[len(s)-1] != ')' {
		return 0, false
	}
	s = s[:len(s)-1] // Remove the closing parenthesis
	n, err := strconv.Atoi(s)
	if err != nil || n < 0 {
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
					if SplitedStr[j] == "\n" {
						break
					}
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
					if SplitedStr[j] == "\n" {
						break
					}
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
					if SplitedStr[j] == "\n" {
						break
					}
					SplitedStr[j] = goreloaded.Capitalize(SplitedStr[j])
				}
				SplitedStr = append(SplitedStr[:i], SplitedStr[i+2:]...)
				i--
			}
		}
	}
	return SplitedStr, i
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input_file> <output_file>")
	} else {
		inputFile := os.Args[1]
		outputFile := os.Args[2]
		if strings.HasSuffix(inputFile, ".txt") {
			content, err := os.ReadFile(inputFile)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}

			str := string(content)
			splitedStr := strings.Split(str, "\n")
			var SplitedStr []string
			for i, line := range splitedStr {
				words := strings.Split(line, " ")
				SplitedStr = append(SplitedStr, words...)
				if i < len(splitedStr)-1 {
					SplitedStr = append(SplitedStr, "\n")
				}
			}
			goreloaded.Atoan(SplitedStr)
			for i := 0; i < len(SplitedStr); i++ {
				if SplitedStr[i] == "(hex)" {
					if i > 0 {
						SplitedStr[i-1] = goreloaded.HexConv(SplitedStr[i-1])
					}
					SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
					i--
				} else if SplitedStr[i] == "(bin)" {
					if i > 0 {
						SplitedStr[i-1] = goreloaded.BinConv(SplitedStr[i-1])
					}
					SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
					i--
				} else if SplitedStr[i] == "(up)" {
					if i > 0 {
						SplitedStr[i-1] = strings.ToUpper(SplitedStr[i-1])
					}
					SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
					i--
				} else if SplitedStr[i] == "(low)" {
					if i > 0 {
						SplitedStr[i-1] = strings.ToLower(SplitedStr[i-1])
					}
					SplitedStr = append(SplitedStr[:i], SplitedStr[i+1:]...)
					i--
				} else if SplitedStr[i] == "(cap)" {
					if i > 0 {
						SplitedStr[i-1] = goreloaded.Capitalize(SplitedStr[i-1])
					}
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
			var output string
			for i, str := range SplitedStr {
				output += str
				if str != "\n" && i < len(SplitedStr)-1 && SplitedStr[i+1] != "\n" {
					output += " "
				}
			}
			output = goreloaded.Punctuations(output)
			output = goreloaded.Quotes(output)
			if strings.HasSuffix(outputFile, ".txt") {
				err = os.WriteFile(outputFile, []byte(output), 0o644)
				if err != nil {
					fmt.Println("Error writing to file:", err)
					return
				}
			} else {
				fmt.Println("Error: outputFile should be .txt format")
			}
			fmt.Println((output))
		} else {
			fmt.Println("Error: inputFile should be .txt format")
		}
	}
}
