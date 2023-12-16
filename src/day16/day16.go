package day16

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type point struct {
	x,
	y int
}

type beam struct {
	position point
	dx,
	dy int
}

func (b beam) step(grid []string, visitedFields *[]beam) {
	if b.dx+b.dy > 1 {
		panic("INVALID VELOCITY!")
	}

	b.position.x += b.dx
	b.position.y += b.dy

	if slices.ContainsFunc(*visitedFields, func(b2 beam) bool {
		return b2.position == b.position && b.dx == b2.dx && b.dy == b2.dy
	}) {
		return
	}

	//stepped outside the grid
	if b.position.x < 0 || b.position.x >= len(grid[0]) || b.position.y < 0 || b.position.y >= len(grid) {
		return
	}

	*visitedFields = append(*visitedFields, b)

	char := grid[b.position.y][b.position.x]

	if char == '\\' {
		b2 := b
		b.dy = b2.dx
		b.dx = b2.dy
		b.step(grid, visitedFields)
		return
	}

	if char == '/' {
		b2 := b
		b.dy = b2.dx * -1
		b.dx = b2.dy * -1
		b.step(grid, visitedFields)
		return
	}

	if char == '-' && b.dy != 0 {
		b.dy = 0
		b.dx = 1
		b2 := b
		b2.dx = -1
		b.step(grid, visitedFields)
		b2.step(grid, visitedFields)
		return
	}

	if char == '|' && b.dx != 0 {
		b.dx = 0
		b.dy = 1
		b2 := b
		b2.dy = -1
		b.step(grid, visitedFields)
		b2.step(grid, visitedFields)
		return
	}

	b.step(grid, visitedFields)
}

func TaskOne() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("couldnt get pwd\nwith err: %v", err.Error())
	}

	inputFile, err := os.ReadFile(dir + "/src/day16/input")
	if err != nil {
		log.Fatalf("failed to open file input\nwith err: %v", err)
	}

	grid := strings.Split(string(inputFile), "\n")
	startBeam := beam{
		dx: 1,
		position: point{
			x: -1,
		},
	}

	var visitedArr []beam
	visited := make(map[point]bool)
	startBeam.step(grid, &visitedArr)

	for _, p := range visitedArr {
		visited[p.position] = true
	}

	fmt.Println(len(visited))
}

func TaskTwo() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("couldnt get pwd\nwith err: %v", err.Error())
	}

	inputFile, err := os.ReadFile(dir + "/src/day16/input")
	if err != nil {
		log.Fatalf("failed to open file input\nwith err: %v", err)
	}

	grid := strings.Split(string(inputFile), "\n")

	maxVisited := 0

	// all edges on the top
	for x := 0; x < len(grid[0]); x++ {
		startBeam := beam{
			dy: 1,
			position: point{
				x: x,
				y: -1,
			},
		}

		var visitedArr []beam
		visitedMap := make(map[point]bool)
		startBeam.step(grid, &visitedArr)

		for _, p := range visitedArr {
			visitedMap[p.position] = true
		}

		visited := len(visitedMap)
		if visited > maxVisited {
			maxVisited = visited
		}
	}

	// all edges on the bottom
	for x := 0; x < len(grid[0]); x++ {
		startBeam := beam{
			dy: -1,
			position: point{
				x: x,
				y: len(grid),
			},
		}

		var visitedArr []beam
		visitedMap := make(map[point]bool)
		startBeam.step(grid, &visitedArr)

		for _, p := range visitedArr {
			visitedMap[p.position] = true
		}

		visited := len(visitedMap)
		if visited > maxVisited {
			maxVisited = visited
		}
	}

	// all edges on the left
	for y := 0; y < len(grid); y++ {
		startBeam := beam{
			dx: 1,
			position: point{
				x: -1,
				y: y,
			},
		}

		var visitedArr []beam
		visitedMap := make(map[point]bool)
		startBeam.step(grid, &visitedArr)

		for _, p := range visitedArr {
			visitedMap[p.position] = true
		}

		visited := len(visitedMap)
		if visited > maxVisited {
			maxVisited = visited
		}
	}

	// all edges on the right
	for y := 0; y < len(grid); y++ {
		startBeam := beam{
			dx: -1,
			position: point{
				x: len(grid[0]),
				y: y,
			},
		}

		var visitedArr []beam
		visitedMap := make(map[point]bool)
		startBeam.step(grid, &visitedArr)

		for _, p := range visitedArr {
			visitedMap[p.position] = true
		}

		visited := len(visitedMap)
		if visited > maxVisited {
			maxVisited = visited
		}
	}

	fmt.Println(maxVisited)
}
