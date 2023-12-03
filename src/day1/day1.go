package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func TaskOne() {
	inputFile, err := os.Open("/home/feyez/coding/Christmas-2023/src/day1/input1")
	defer func() {
		if err := inputFile.Close(); err != nil {
			panic(err.Error())
		}
	}()
	if err != nil {
		log.Fatalf("failed to open file input1\nwith err: %v", err.Error())
	}

	var result int
	fileScanner := bufio.NewScanner(inputFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()

		var res string
		for _, char := range line {
			if unicode.IsDigit(char) {
				res += string(char)
				break
			}
		}

		for x := len(line) - 1; x >= 0; x-- {
			if unicode.IsDigit(rune(line[x])) {
				res += string(line[x])
				break
			}
		}

		num, err := strconv.Atoi(res)
		if err != nil {
			log.Fatalf("could not parse int for string: %v", res)
		}
		result += num
	}

	fmt.Printf("Success!\nResult:%v", result)
}

func TaskTwo() {
	inputFile, err := os.Open("/home/feyez/coding/Christmas-2023/src/day1/input1")
	defer func() {
		if err := inputFile.Close(); err != nil {
			panic(err.Error())
		}
	}()
	if err != nil {
		log.Fatalf("failed to open file input1\nwith err: %v", err)
	}

	digitMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var result int
	fileScanner := bufio.NewScanner(inputFile)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		var res string

		firstDigitIndex := -1
		var firstDigit *string
		for index, char := range line {
			if unicode.IsDigit(char) {
				digit := string(char)
				firstDigit = &digit
				firstDigitIndex = index
				break
			}
		}

		for text, d := range digitMap {
			digit := d
			if index := strings.Index(line, text); index > -1 && index < firstDigitIndex {
				firstDigitIndex = index
				firstDigit = &digit
			}
		}

		var secondDigitIndex *int
		var secondDigit *string
		for x := len(line) - 1; x >= 0; x-- {
			if unicode.IsDigit(rune(line[x])) {
				digit := string(line[x])
				secondDigit = &digit
				secondDigitIndex = &x
				break
			}
		}

		for text, d := range digitMap {
			digit := d
			if index := strings.LastIndex(line, text); secondDigitIndex == nil && index > -1 || index > *secondDigitIndex {
				secondDigitIndex = &index
				secondDigit = &digit
			}
		}

		num, err := strconv.Atoi(*firstDigit + *secondDigit)
		if err != nil {
			log.Fatalf("could not parse int for string: %v", res)
		}
		result += num
	}

	fmt.Printf("Success!\nResult:%v", result)
}
