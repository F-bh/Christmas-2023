package day15

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func hash(input string) int {
	entrySum := 0
	for _, char := range input {
		entrySum += int(char)
		if int(char) > 255 {
			panic("invald ascii character: " + string(char))
		}

		entrySum *= 17
		entrySum %= 256
	}
	return entrySum
}

func TaskOne() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("couldnt get pwd\nwith err: %v", err.Error())
	}

	inputFile, err := os.ReadFile(dir + "/src/day15/input")
	if err != nil {
		log.Fatalf("failed to open file input\nwith err: %v", err)
	}

	entries := strings.Split(string(inputFile), ",")

	sum := 0
	for _, entry := range entries {
		sum += hash(entry)
	}

	fmt.Println(sum)
}

type lens struct {
	focal int
	label string
}

func TaskTwo() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("couldnt get pwd\nwith err: %v", err.Error())
	}

	inputFile, err := os.ReadFile(dir + "/src/day15/input")
	if err != nil {
		log.Fatalf("failed to open file input\nwith err: %v", err)
	}

	entries := strings.Split(string(inputFile), ",")

	boxes := make(map[int][]lens)

	for _, entry := range entries {
		var boxHash int
		var newLens lens

		if entry[len(entry)-1] == '-' {
			newLens.label = entry[:len(entry)-1]
		} else {
			split := strings.Split(entry, "=")
			label, focalStr := split[0], split[1]
			focal, err := strconv.Atoi(focalStr)
			if err != nil {
				log.Fatalf("couldn't parse focal length: %v of box %v", focalStr, entry)
			}

			newLens = lens{
				focal,
				label,
			}
		}

		boxHash = hash(newLens.label)
		lenses, ok := boxes[boxHash]
		if !ok {
			boxes[boxHash] = []lens{}
		}

		//remove from box
		if entry[len(entry)-1] == '-' {
			newLenses := slices.DeleteFunc(lenses, func(in lens) bool { return in.label == newLens.label })
			boxes[boxHash] = newLenses
			fmt.Printf("\nOp: %v - boxes:\n", entry)
			for h, ls := range boxes {
				fmt.Printf("Box: %v\nContains: %v\n\n", h, ls)
			}

			continue
		}

		// add to box
		lenses = append(lenses, newLens)
		boxes[boxHash] = lenses

		fmt.Printf("\nOp: %v - boxes:\n", entry)
		for h, ls := range boxes {
			fmt.Printf("Box: %v\nContains: %v\n\n", h, ls)
		}
	}
}
