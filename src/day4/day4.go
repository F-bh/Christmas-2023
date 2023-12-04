package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func TaskOne() {
	inputFile, err := os.Open("/home/feyez/coding/Christmas-2023/src/day4/input4")
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
		res := 0
		line := fileScanner.Text()
		split := strings.Split(line, " | ")
		winning := strings.Split(split[1], " ")
		game := strings.Split(strings.Split(split[0], ": ")[1], " ")

	guesses:
		for _, entry := range game {
			for _, winner := range winning {
				entry := strings.Trim(entry, " ")
				winner = strings.Trim(winner, " ")

				if entry != "" && winner != "" && entry == winner {
					if res == 0 {
						res = 1
					} else {
						res *= 2
					}
					continue guesses
				}
			}
		}

		result += res
	}

	fmt.Printf("Result:%+v", result)
}

type card struct {
	wins int
	next []*card
}

func (c *card) nodes() int {
	result := c.wins
	for _, c2 := range c.next {
		result += c2.nodes()
	}
	return result
}

func TaskTwo() {
	inputFile, err := os.Open("/home/feyez/coding/Christmas-2023/src/day4/input4")
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

	var cards []card

	for fileScanner.Scan() {
		res := 0
		line := fileScanner.Text()
		split := strings.Split(line, " | ")
		winning := strings.Split(split[1], " ")
		game := strings.Split(strings.Split(split[0], ": ")[1], " ")

	guesses:
		for _, entry := range game {
			for _, winner := range winning {
				entry := strings.Trim(entry, " ")
				winner = strings.Trim(winner, " ")

				if entry != "" && winner != "" && entry == winner {
					res += 1
					continue guesses
				}
			}
		}
		cards = append(cards, card{
			wins: res,
		})
	}

	//populate trees
	for x := 0; x < len(cards)-1; x++ {
		c := &cards[x]
		for y := x + 1; y <= x+c.wins; y++ {
			pointsTo := &cards[y]
			c.next = append(c.next, pointsTo)
		}
	}

	//count nodes
	for _, card := range cards {
		result += 1
		result += card.nodes()
	}

	fmt.Printf("Result:%+v", result)
}
