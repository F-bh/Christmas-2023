package day11

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type point struct {
	x,
	y int
}

type couple struct {
	start,
	goal point
}

func (c couple) shortestPathDistanceWithExpansion(expansion int, colsToExpand []int, rowsToExpand []int) int {
	xGoal := c.goal.x
	xStart := c.start.x

	for _, col := range colsToExpand {
		if col < c.goal.x {
			xGoal += expansion
		}

		if col < c.start.x {
			xStart += expansion
		}
	}

	yGoal := c.goal.y
	yStart := c.start.y

	for _, row := range rowsToExpand {
		if row < c.goal.y {
			yGoal += expansion
		}

		if row < c.start.y {
			yStart += expansion
		}
	}

	xdiff := xGoal - xStart
	if xdiff < 0 {
		xdiff = xStart - xGoal
	}

	ydiff := yGoal - yStart
	if ydiff < 0 {
		ydiff = yStart - yGoal
	}

	return xdiff + ydiff
}

func getCouplings(input []string) (output []couple) {
	var points []point

	for lineIndex, line := range input {
		for charIndex, char := range line {
			if char == '#' {
				points = append(points, point{x: charIndex, y: lineIndex})
			}
		}
	}

	for pointIndex, currPoint := range points {
		for _, nextPoint := range points[pointIndex+1:] {
			output = append(output, couple{currPoint, nextPoint})
		}
	}
	return
}

func getRowIndexesToExpand(input []string) (output []int) {
	for rowIndex, row := range input {
		if strings.Count(row, "#") == 0 {
			output = append(output, rowIndex)
		}
	}
	return
}

func getColIndexesToExpand(input []string) (output []int) {
	var columns []string
	for _, line := range input {
		for columnIndex, char := range line {
			if len(columns) < columnIndex+1 {
				columns = append(columns, "")
			}
			columns[columnIndex] = columns[columnIndex] + string(char)
		}
	}

	for columnIndex, column := range columns {
		if strings.Count(column, "#") == 0 {
			output = append(output, columnIndex)
		}
	}
	return
}

func TaskOne() {
	inputFile, err := os.ReadFile("/home/feyez/coding/Christmas-2023/src/day11/input")
	if err != nil {
		log.Fatalf("failed to open file input\nwith err: %v", err.Error())
	}

	lines := strings.Split(string(inputFile), "\n")
	columnsToExpand := getColIndexesToExpand(lines)
	rowsToExpand := getRowIndexesToExpand(lines)

	couplings := getCouplings(lines)

	res := 0
	for _, c := range couplings {
		res += c.shortestPathDistanceWithExpansion(1, columnsToExpand, rowsToExpand)
	}

	fmt.Printf("Result: %v", res)
}

func TaskTwo() {
	inputFile, err := os.ReadFile("/home/feyez/coding/Christmas-2023/src/day11/input")
	if err != nil {
		log.Fatalf("failed to open file input\nwith err: %v", err.Error())
	}

	lines := strings.Split(string(inputFile), "\n")
	columnsToExpand := getColIndexesToExpand(lines)
	rowsToExpand := getRowIndexesToExpand(lines)

	couplings := getCouplings(lines)

	res := 0
	for _, c := range couplings {
		res += c.shortestPathDistanceWithExpansion(1000000-1, columnsToExpand, rowsToExpand)
	}

	fmt.Printf("Result: %v", res)
}
