package day14

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func TaskOne() {
	inputFile, err := os.ReadFile("/home/feyez/coding/Christmas-2023/src/day14/input")
	if err != nil {
		log.Fatalf("failed to open file input\nwith err: %v", err.Error())
	}

	var entities [][]string
	for _, line := range strings.Split(string(inputFile), "\n") {
		entities = append(entities, strings.Split(line, ""))
	}

	moveRocksUp(entities)

	for _, e := range entities {
		fmt.Println(e)
	}

	sum := 0
	slices.Reverse(entities)

	for load, line := range entities {
		for _, item := range line {
			if item == "O" {
				sum += load + 1
			}
		}
	}

	fmt.Println(sum)
}

func TaskTwo() {
	inputFile, err := os.ReadFile("/home/feyez/coding/Christmas-2023/src/day14/input")
	if err != nil {
		log.Fatalf("failed to open file input\nwith err: %v", err.Error())
	}

	var entities [][]string
	for _, line := range strings.Split(string(inputFile), "\n") {
		entities = append(entities, strings.Split(line, ""))
	}

	prev := entities
	eqs := 0
	for x := 0; x < 1000000000 && eqs < 100; x++ {
		if x%100 == 0 {
			println(x)
		}
		for y := 0; y < 4; y++ {
			moveRocksUp(entities)
			entities = rotateRight(entities)
			prev = entities
		}

		if matrixEqual(entities, prev) {
			eqs++
		}
	}

	for _, e := range entities {
		fmt.Println(e)
	}

	sum := 0
	slices.Reverse(entities)

	for load, line := range entities {
		for _, item := range line {
			if item == "O" {
				sum += load + 1
			}
		}
	}

	fmt.Println(sum)

}

func moveRocksUp(input [][]string) {
	for y := range input {
		for x, char := range input[y] {
			if char == "O" {
				for y2 := y - 1; y2 >= 0; y2-- {
					if input[y2][x] == "O" || input[y2][x] == "#" {
						input[y2+1][x] = char
						if y2+1 != y {
							input[y][x] = "."
						}
						break
					} else if y2 == 0 && input[y2][x] == "." {
						input[y2][x] = char
						if y2 != y {
							input[y][x] = "."
						}
						break
					}
				}
			}
		}
	}
}

func matrixEqual(a, b [][]string) bool {
	for index, line := range a {
		if !slices.Equal(line, b[index]) {
			return false
		}
	}
	return true
}

func rotateRight(input [][]string) [][]string {
	inputFlipped := make([][]string, len(input[0]), len(input[0]))
	for _, line := range input {
		slices.Reverse(line)
		for colIndex, col := range line {
			inputFlipped[colIndex] = append(inputFlipped[colIndex], col)
		}
	}

	for x := 0; x < len(inputFlipped); x++ {
		slices.Reverse(inputFlipped[x])
	}

	slices.Reverse(inputFlipped)

	return inputFlipped
}
