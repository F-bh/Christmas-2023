package day17

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x,
	y,
	heat int
}

func (p point) cost(endX, endY int) int {
	return (endX - p.x) + (endY - p.y) + p.heat
}

func TaskOne() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("couldnt get pwd\nwith err: %v", err.Error())
	}

	inputFile, err := os.ReadFile(dir + "/src/day17/input")
	if err != nil {
		log.Fatalf("failed to open file input\nwith err: %v", err)
	}

	var grid [][]int
	for _, line := range strings.Split(string(inputFile), "\n") {
		var items []int
		for _, char := range line {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				log.Fatalf("couldn't parse int: %v", char)
			}
			items = append(items, num)
		}
		grid = append(grid, items)
	}

	path := findShortestPath(grid)
	fmt.Println(path)
}

type direction int

const (
	N direction = iota
	E
	S
	W
)

func findShortestPath(grid [][]int) []point {
	endY := len(grid) - 1
	endX := len(grid[0])

	var visited []point
	var lastDirection *direction
	stepsInDir := 0

	for {
		var current point
		if len(visited) != 0 {
			current = visited[len(visited)-1]
			if current.x == endX && current.y == endY {
				break
			}
		} else {
			current = point{}
		}

		var possible [4]*point
		//check above^
		if ((lastDirection != nil && *lastDirection != N && stepsInDir < 3) || lastDirection == nil) && current.y > 0 {
			if lastDirection != nil && *lastDirection != S || lastDirection == nil {
				heat := grid[current.y-1][current.x]
				possible[0] = &point{x: current.x, y: current.y - 1, heat: heat}
			}
		}

		//check right
		if ((lastDirection != nil && *lastDirection != E && stepsInDir < 3) || lastDirection == nil) && current.x < len(grid[0])-1 {
			if lastDirection != nil && *lastDirection != W || lastDirection == nil {
				heat := grid[current.y][current.x+1]
				possible[1] = &point{x: current.x + 1, y: current.y, heat: heat}
			}
		}

		//check below
		if ((lastDirection != nil && *lastDirection != S && stepsInDir < 3) || lastDirection == nil) && current.y < len(grid)-1 {
			if lastDirection != nil && *lastDirection != N || lastDirection == nil {
				heat := grid[current.y+1][current.x]
				possible[2] = &point{x: current.x, y: current.y + 1, heat: heat}
			}
		}

		//check left
		if ((lastDirection != nil && *lastDirection != W && stepsInDir < 3) || lastDirection == nil) && current.x > 0 {
			if lastDirection != nil && *lastDirection != E || lastDirection == nil {
				heat := grid[current.y][current.x-1]
				possible[3] = &point{x: current.x - 1, y: current.y, heat: heat}
			}
		}

		var cheapest *point
		var dir direction
		for d, p := range possible {
			if p == nil {
				continue
			}

			if cheapest == nil {
				cheapest = p
				dir = direction(d)
				continue
			}

			if p.cost(endX, endY) < cheapest.cost(endX, endY) {
				cheapest = p
				dir = direction(d)
			}
		}

		if lastDirection != nil && dir != *lastDirection {
			stepsInDir = 1
		} else {
			stepsInDir++
		}

		lastDirection = &dir
		visited = append(visited, *cheapest)
	}

	return visited
}
