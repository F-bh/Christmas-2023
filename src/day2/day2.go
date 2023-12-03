package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func TaskOne() {
	allowed := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	inputFile, err := os.Open("/home/feyez/coding/Christmas-2023/src/day2/input2")
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
		splitByColon := strings.Split(line, ":")
		gameId := strings.Split(splitByColon[0], " ")
		gameIdInt, err := strconv.Atoi(gameId[1])
		if err != nil {
			log.Fatalf("failed to parse game Id int read: %v", gameId)
		}

		showings := strings.Split(splitByColon[1], ";")
		gameOk := true
	game:
		for _, showing := range showings {
			colours := strings.Split(showing, ",")
			found := false

		colourLoop:
			for _, colour := range colours {
				colour := colour[1:] //remove leading white space
				for lookupColour, allowedAmount := range allowed {
					if strings.Contains(colour, lookupColour) {
						found = true
						amountStr := strings.Split(colour, " ")[0]
						amount, err := strconv.Atoi(amountStr)
						if err != nil {
							log.Fatalf("failed to parse colour amount str to int, read: %v", amountStr)
						}

						if amount > allowedAmount {
							gameOk = false
							break game
						}
						continue colourLoop
					}
				}
				if !found {
					log.Fatalf("no colour found in showing: %v", showing)
				}
			}
		}
		if gameOk {
			result += gameIdInt
		}

	}
	fmt.Printf("Result:%v", result)
}

func TaskTwo() {
	inputFile, err := os.Open("/home/feyez/coding/Christmas-2023/src/day2/input2")
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
		splitByColon := strings.Split(line, ":")

		showings := strings.Split(splitByColon[1], ";")

		maxMap := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, showing := range showings {
			colours := strings.Split(showing, ",")
			for _, colour := range colours {
				colour := colour[1:] //remove leading white space
				for lookupColour, maxFound := range maxMap {
					if strings.Contains(colour, lookupColour) {
						amountStr := strings.Split(colour, " ")[0]
						amount, err := strconv.Atoi(amountStr)
						if err != nil {
							log.Fatalf("failed to parse colour amount str to int, read: %v", amountStr)
						}

						if amount > maxFound {
							maxMap[lookupColour] = amount
							break
						}
					}
				}
			}
		}

		result += func() (power int) {
			for _, v := range maxMap {
				if power == 0 {
					power = v
					continue
				}
				power *= v
			}
			return power
		}()
	}
	fmt.Printf("Result:%v", result)
}
