package day3

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func TaskOne() {
	inputFile, err := os.ReadFile("/home/feyez/coding/Christmas-2023/src/day3/input3")
	if err != nil {
		log.Fatalf("failed to read file input1\nwith err: %v", err.Error())
	}
	input := string(inputFile)
	lines := strings.Split(input, "\n")
	result := 0

	var parts []string

	//parse
	for line, lineContent := range lines {
		number := ""
		startIndex := -1

		for index := 0; index < len(lineContent); index++ {
			char := rune(lineContent[index])

			if number == "" && !unicode.IsNumber(char) {
				continue
			}

			//build number
			if unicode.IsNumber(char) {
				if number == "" {
					startIndex = index
				}
				number += string(char)
				if index != len(lineContent)-1 {
					continue
				}
			}

			//check if symbol is adjacent on same line
			if char != '.' && !unicode.IsNumber(char) {
				parts = append(parts, number)
				number = ""
				startIndex = -1
				continue
			}

			if char == '.' || index == len(lineContent)-1 {
				var prevChar *rune
				if startIndex > 0 {
					tmp := rune(lineContent[startIndex-1])
					prevChar = &tmp
				}

				if prevChar != nil && *prevChar != '.' && !unicode.IsNumber(*prevChar) {
					parts = append(parts, number)
					number = ""
					startIndex = -1
					continue
				}

				containsSymbol := func(checkContent string) bool {
					checkIndex := startIndex - 1
					endCheckIndex := index + 1
					if startIndex <= 0 {
						checkIndex = 0
					}

					if index == len(lineContent)-1 {
						endCheckIndex = index
					}

					for _, char := range checkContent[checkIndex:endCheckIndex] {
						if char != '.' && !unicode.IsNumber(char) {
							return true
						}
					}
					return false
				}

				//check if symbol is adjacent on other lines
				if line != 0 {
					aboveLineContent := lines[line-1]
					if containsSymbol(aboveLineContent) {
						parts = append(parts, number)
						number = ""
						startIndex = -1
						continue
					}
				}

				if line != len(lines)-1 {
					belowLineContent := lines[line+1]
					if containsSymbol(belowLineContent) {
						parts = append(parts, number)
						number = ""
						startIndex = -1
						continue
					}
				}
				number = ""
				startIndex = -1
			}
		}
	}

	for _, part := range parts {
		tmp, err := strconv.Atoi(part)
		if err != nil {
			log.Fatalf("failed to parse part number: %v", part)
		}
		result += tmp
	}
	fmt.Printf("Result:%+v", result)
}

func TaskTwo() {
	inputFile, err := os.ReadFile("/home/feyez/coding/Christmas-2023/src/day3/input3")
	if err != nil {
		log.Fatalf("failed to read file input1\nwith err: %v", err.Error())
	}
	input := string(inputFile)
	_ = input
	var result int

	fmt.Printf("Result:%+v", result)
}
