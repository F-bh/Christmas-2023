package day13

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type mirror struct {
	startLine,
	endLine int
}

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

	var hMirrors []mirror
	for _, puzzle := range puzzles {
		hMirrors = append(hMirrors, findMirrors(puzzle)...)
	}

	hSum := 0
	for _, m := range hMirrors {
		//+1 to handle division by zero
		hSum += ((m.endLine + 1) / (m.startLine + 1)) + m.startLine
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
	var vMirrors []mirror
	for _, puzzle := range puzzlesFlipped {
		vMirrors = append(vMirrors, findMirrors(puzzle)...)
	}

	for _, m := range vMirrors {
		//+1 to handle division by zero
		vSum += ((m.endLine + 1) / (m.startLine + 1)) + m.startLine
	}

	fmt.Println(hSum*100 + vSum)
}

func findMirrors(puzzle []string) []mirror {
	var mirrors []mirror

findMirrored:
	for lineIndex, line := range puzzle {
		possibleMirror := mirror{startLine: lineIndex}
		for nextIndex := lineIndex + 1; nextIndex < len(puzzle); nextIndex++ {
			nextLine := puzzle[nextIndex]

			if line == nextLine {
				possibleMirror.endLine = nextIndex
				mirrors = append(mirrors, possibleMirror)
				continue findMirrored
			}

		}

	}

	//find edges of mirrors (they should already be ordered)
	var validMirrors []mirror
	for mirrorIndex := 0; mirrorIndex <= len(mirrors)-1; mirrorIndex++ {
		m := mirrors[mirrorIndex]

		//only a single line is mirrored
		if len(mirrors) == 1 && m.endLine-m.startLine == 1 {
			validMirrors = append(validMirrors, m)
			mirrorIndex = m.startLine
			break
		}

		//first two lines are mirrored
		if mirrorIndex == 0 && m.endLine-m.startLine == 1 {
			validMirrors = append(validMirrors, m)
			continue
		}

		//last two lines are mirrored
		if mirrorIndex == len(mirrors)-2 && m.endLine-m.startLine == 1 {
			validMirrors = append(validMirrors, m)
			break
		}

		//all lines up to the end of this mirror are also mirrored inside it's start and end
		containedMirrors := 0
		for _, m2 := range mirrors[mirrorIndex+1:] {
			if m2.startLine > m.startLine && m2.endLine < m.endLine {
				containedMirrors++
				continue
			}

			break
		}

		if containedMirrors*2 == m.endLine-m.startLine-1 {
			validMirrors = append(validMirrors, m)
			mirrorIndex += containedMirrors
			continue
		}

		mirrorIndex++
	}

	return validMirrors
}
