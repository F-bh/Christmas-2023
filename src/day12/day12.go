package day12

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type inputLine struct {
	searchString string
	fragments    []int
}

func checkValid(input string, fragments []int) bool {
	fragmentSum := 0
	for _, frag := range fragments {
		fragmentSum += frag
	}

	if strings.Count(input, "!")+strings.Count(input, "#") != fragmentSum {
		return false
	}

	searchString := input
fragmentSearch:
	for _, fragment := range fragments {
		foundFragment := ""
		for index, char := range searchString {
			if char == '!' || char == '#' {
				foundFragment += string(char)
				if index != len(searchString)-1 {
					continue
				}
			}

			if foundFragment != "" {
				if len(foundFragment) != fragment {
					return false
				}
				searchString = searchString[index:]
				continue fragmentSearch
			}

		}
	}

	return true
}

func getPossibleInsertions(input inputLine) (result int) {
	possibleStrings := []string{input.searchString}
	var newPossibleStrings []*string
	newPossibleStrings = append(newPossibleStrings, &input.searchString)

	//outer:
	for {
		if len(newPossibleStrings) == 0 {
			break
		}
		var newNewPossibleStrings []*string
		for _, possibility := range newPossibleStrings {
			possibility := *possibility
			for index, char := range possibility {
				if char == '?' {
					newPossibility := possibility[:index] + "!" + possibility[index+1:]

					if !slices.Contains(possibleStrings, newPossibility) && checkValid(newPossibility, input.fragments) {
						possibleStrings = append(possibleStrings, newPossibility)
						continue
					}

					newNewPossibleStrings = append(newNewPossibleStrings, &newPossibility)
				}
			}
		}

		newPossibleStrings = newNewPossibleStrings
	}

	return len(possibleStrings) - 1
}

func TaskOne() {
	inputFile, err := os.ReadFile("/home/feyez/coding/Christmas-2023/src/day12/input")
	if err != nil {
		log.Fatalf("failed to open file input\nwith err: %v", err.Error())
	}

	lines := strings.Split(string(inputFile), "\n")
	var parsedLines []inputLine

	for _, line := range lines {
		split := strings.Split(line, " ")
		var fragments []int
		for _, fragment := range strings.Split(split[1], ",") {
			frag, err := strconv.Atoi(fragment)
			if err != nil {
				log.Fatalf("failed to parse fragment %v\nwith err: %v", fragment, err.Error())
			}
			fragments = append(fragments, frag)
		}
		parsedLines = append(parsedLines, inputLine{
			searchString: split[0],
			fragments:    fragments,
		})
	}

	var result int
	for index, line := range parsedLines {
		result += getPossibleInsertions(line)
		fmt.Printf("finished line %v\n", index+1)
	}

	fmt.Printf("Result: %v", result)
}
