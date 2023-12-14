package day13

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func TaskOne() {
	inputFile, err := os.ReadFile("/home/feyez/coding/Christmas-2023/src/day13/input")
	if err != nil {
		log.Fatalf("failed to open file input\nwith err: %v", err.Error())
	}

	inputs := strings.Split(string(inputFile), "\n\n")
	var puzzles [][]string

	for _, puzzle := range inputs {
		puzzles = append(puzzles, strings.Split(puzzle, "\n"))
	}

	hSum := 0
	for _, puzzle := range puzzles {
		hSum += findMirror(puzzle)
	}

	//flip to vertical
	var puzzlesFlipped [][]string
	for _, puzzle := range puzzles {
		var lines []string
		for x := 0; x < len(puzzle[0]); x++ {
			lines = append(lines, "")
		}

		for _, puzzleLine := range puzzle {
			for charIndex, char := range puzzleLine {
				lines[charIndex] += string(char)
			}
		}

		puzzlesFlipped = append(puzzlesFlipped, lines)
	}

	vSum := 0
	for _, puzzle := range puzzlesFlipped {
		vSum += findMirror(puzzle)
	}

	fmt.Println(hSum*100 + vSum)
}

func TaskTwo() {
	inputFile, err := os.ReadFile("/home/feyez/coding/Christmas-2023/src/day13/input")
	if err != nil {
		log.Fatalf("failed to open file input\nwith err: %v", err.Error())
	}

	inputs := strings.Split(string(inputFile), "\n\n")
	var puzzles [][]string

	for _, puzzle := range inputs {
		puzzles = append(puzzles, strings.Split(puzzle, "\n"))
	}

	hSum := 0
	for _, puzzle := range puzzles {
		hSum += findMirrorWithSmudge(puzzle)
	}

	//flip to vertical
	var puzzlesFlipped [][]string
	for _, puzzle := range puzzles {
		var lines []string
		for x := 0; x < len(puzzle[0]); x++ {
			lines = append(lines, "")
		}

		for _, puzzleLine := range puzzle {
			for charIndex, char := range puzzleLine {
				lines[charIndex] += string(char)
			}
		}

		puzzlesFlipped = append(puzzlesFlipped, lines)
	}

	vSum := 0
	for _, puzzle := range puzzlesFlipped {
		vSum += findMirrorWithSmudge(puzzle)
	}

	fmt.Println(hSum*100 + vSum)
}

func findMirror(puzzle []string) int {
	for lineIndex := range puzzle {
		if lineIndex == 0 {
			continue
		}

		upper := func() []string {
			var ret []string
			for x := len(puzzle[:lineIndex]) - 1; x >= 0; x-- {
				ret = append(ret, puzzle[:lineIndex][x])
			}
			return ret
		}()
		lower := puzzle[lineIndex:]

		if len(upper) < len(lower) {
			lower = lower[:len(upper)]
		} else {
			upper = upper[:len(lower)]
		}

		if slices.Equal(upper, lower) {
			return lineIndex
		}
	}

	return 0
}

func findMirrorWithSmudge(puzzle []string) int {
	for lineIndex := range puzzle {
		if lineIndex == 0 {
			continue
		}

		upper := func() []string {
			var ret []string
			for x := len(puzzle[:lineIndex]) - 1; x >= 0; x-- {
				ret = append(ret, puzzle[:lineIndex][x])
			}
			return ret
		}()
		lower := puzzle[lineIndex:]

		if len(upper) < len(lower) {
			lower = lower[:len(upper)]
		} else {
			upper = upper[:len(lower)]
		}

		differences := 0
		for x, line := range upper {
			for charIndex, char := range line {
				if char != rune(lower[x][charIndex]) {
					differences++
				}
			}

		}

		if differences == 1 {
			return lineIndex
		}
	}
	return 0
}
