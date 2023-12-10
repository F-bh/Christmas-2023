package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type tile struct {
	x    int
	y    int
	pipe string
}

type step struct {
	tile
	velocityX int
	velocityY int
}

func stepForward(input [][]string, lastStep step) step {
	if lastStep.velocityY > 1 || lastStep.velocityX > 1 || lastStep.velocityY < -1 || lastStep.velocityX < -1 {
		log.Fatalf("invalid input step:\n%v\n", lastStep)
	}
	currentStep := step{
		tile: tile{
			x: lastStep.x + lastStep.velocityX,
			y: lastStep.y + lastStep.velocityY,
		},
	}
	currentStep.pipe = input[currentStep.y][currentStep.x]

	switch currentStep.pipe {
	case "S":
		currentStep.velocityX = lastStep.velocityX
		currentStep.velocityY = lastStep.velocityY
	case "-":
		currentStep.velocityX = lastStep.velocityX
	case "|":
		currentStep.velocityY = lastStep.velocityY
	case "F":
		//last step below
		if lastStep.velocityY == -1 {
			currentStep.velocityX = 1
			break
		}

		//last step from right
		currentStep.velocityY = 1
	case "7":
		//last step below
		if lastStep.velocityY == -1 {
			currentStep.velocityX = -1
			break
		}

		currentStep.velocityY = 1
	case "J":
		//last step above
		if lastStep.velocityY == 1 {
			currentStep.velocityX = -1
			break
		}

		//last step from left
		currentStep.velocityY = -1
	case "L":
		//last step above
		if lastStep.velocityY == 1 {
			currentStep.velocityX = 1
			break
		}

		//last step from right
		currentStep.velocityY = -1
	default:
		panic("invalid loop encountered")
	}

	return currentStep
}

func getMainLoop(input [][]string) (loop []step) {
	startTile := tile{
		pipe: "S",
	}

searchStart:
	for rowIndex, row := range input {
		for columnIndex, column := range row {
			if column == "S" {
				startTile.y, startTile.x = rowIndex, columnIndex
				break searchStart
			}
		}
	}

	stepRow, stepCol := 0, 0
	//check left
	if startTile.x != 0 && strings.Contains("-FL", input[startTile.y][startTile.y-1]) {
		stepRow, stepCol = startTile.y, startTile.x-1
	}

	//check right
	if startTile.x != len(input[startTile.y])-1 && strings.Contains("-7J", input[startTile.y][startTile.x+1]) {
		stepRow, stepCol = startTile.y, startTile.x-1
	}

	//check above
	if startTile.y != 0 && strings.Contains("|F7", input[startTile.y-1][startTile.x]) {
		stepRow, stepCol = startTile.y-1, startTile.x
	}

	//check below
	if startTile.y != len(input)-1 && strings.Contains("|LJ", input[startTile.y+1][startTile.x]) {
		stepRow, stepCol = startTile.y+1, startTile.x
	}

	lastStep := step{
		tile:      startTile,
		velocityX: stepCol - startTile.x,
		velocityY: stepRow - startTile.y,
	}

	lastStep = stepForward(input, lastStep)
	for lastStep.pipe != "S" {
		loop = append(loop, lastStep)
		lastStep = stepForward(input, lastStep)
	}

	return loop
}

func TaskOne() {
	inputFile, err := os.Open("/home/feyez/coding/Christmas-2023/src/day10/input")
	defer func() {
		if err := inputFile.Close(); err != nil {
			panic(err.Error())
		}
	}()

	if err != nil {
		log.Fatalf("failed to open file input1\nwith err: %v", err.Error())
	}

	fileScanner := bufio.NewScanner(inputFile)

	var pipes [][]string
	for fileScanner.Scan() {
		var pipesInRow []string
		for _, r := range fileScanner.Text() {
			pipesInRow = append(pipesInRow, string(r))
		}

		pipes = append(pipes, pipesInRow)
	}

	loop := getMainLoop(pipes)

	fmt.Printf("Result: %v", len(loop)/2+1)
}
