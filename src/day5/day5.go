package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type mapping struct {
	fromRangeStart,
	toRangeStart,
	len int
}

func TaskOne() {
	inputFile, err := os.Open("/home/feyez/coding/Christmas-2023/src/day5/input5")
	defer func() {
		if err := inputFile.Close(); err != nil {
			panic(err.Error())
		}
	}()
	if err != nil {
		log.Fatalf("failed to open file input1\nwith err: %v", err.Error())
	}

	fileScanner := bufio.NewScanner(inputFile)
	if !fileScanner.Scan() {
		panic("Invalid Input!")
	}

	var seeds []int
	for _, s := range strings.Split(strings.Split(fileScanner.Text(), ": ")[1], " ") {
		tmp, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("failed to parse seed to int: %v\n", s)
		}
		seeds = append(seeds, tmp)
	}

	var mappings [][]mapping

	currentMapping := -1
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Trim(line, " ") == "" {
			continue
		}

		if strings.Contains(line, ":") {
			currentMapping += 1
			continue
		}

		var parsedMapping mapping

		for x, value := range strings.Split(line, " ") {
			tmp, err := strconv.Atoi(value)
			if err != nil {
				log.Fatalf("failed to parse mapping to int: %v\n", tmp)
			}

			switch x {
			case 0:
				parsedMapping.toRangeStart = tmp
			case 1:
				parsedMapping.fromRangeStart = tmp
			case 2:
				parsedMapping.len = tmp
			default:
				log.Fatalf("invalid mapping length, line: %v\n data: %v\n", currentMapping, line)
			}
		}

		if len(mappings)-1 < currentMapping {
			mappings = append(mappings, []mapping{})
		}

		mappings[currentMapping] = append(mappings[currentMapping], parsedMapping)

	}

	for sx, _ := range seeds {
		seed := &seeds[sx]
	mappingType:
		for _, mt := range mappings {
			for _, m := range mt {
				offset := *seed - m.fromRangeStart
				newSeed := m.toRangeStart + offset
				if *seed >= m.fromRangeStart && *seed <= m.fromRangeStart+m.len-1 {
					*seed = newSeed
					continue mappingType
				}
			}
		}
	}

	lowest := -1
	for _, location := range seeds {
		if lowest == -1 {
			lowest = location
		}

		if lowest > location {
			lowest = location
		}
	}

	fmt.Printf("Result:%+v", lowest)
}

func TaskTwo() {
	inputFile, err := os.Open("/home/feyez/coding/Christmas-2023/src/day5/input5")
	defer func() {
		if err := inputFile.Close(); err != nil {
			panic(err.Error())
		}
	}()
	if err != nil {
		log.Fatalf("failed to open file input1\nwith err: %v", err.Error())
	}

	fileScanner := bufio.NewScanner(inputFile)
	if !fileScanner.Scan() {
		panic("Invalid Input!")
	}

	tmpSeedIn := strings.Split(strings.Split(fileScanner.Text(), ": ")[1], " ")

	var mappings [][]mapping
	currentMapping := -1
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if strings.Trim(line, " ") == "" {
			continue
		}

		if strings.Contains(line, ":") {
			currentMapping += 1
			continue
		}

		var parsedMapping mapping

		for x, value := range strings.Split(line, " ") {
			tmp, err := strconv.Atoi(value)
			if err != nil {
				log.Fatalf("failed to parse mapping to int: %v\n", tmp)
			}

			switch x {
			case 0:
				parsedMapping.toRangeStart = tmp
			case 1:
				parsedMapping.fromRangeStart = tmp
			case 2:
				parsedMapping.len = tmp
			default:
				log.Fatalf("invalid mapping length, line: %v\n data: %v\n", currentMapping, line)
			}
		}

		if len(mappings)-1 < currentMapping {
			mappings = append(mappings, []mapping{})
		}

		mappings[currentMapping] = append(mappings[currentMapping], parsedMapping)
	}

	endSeeds := make(chan int, 1000)
	wg := sync.WaitGroup{}
	lowest := -1

	for sx := 0; sx < len(tmpSeedIn); sx = sx + 2 {
		from, err := strconv.Atoi(tmpSeedIn[sx])
		if err != nil {
			log.Fatalf("failed to parse seed start to int: %v\n", tmpSeedIn[sx])
		}

		length, err := strconv.Atoi(tmpSeedIn[sx+1])
		if err != nil {
			log.Fatalf("failed to parse seed len to int: %v\n", tmpSeedIn[sx+1])
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			for seed := from; seed < from+length; seed++ {
				seedCp := seed
			mappingType:
				for _, mt := range mappings {
					for _, m := range mt {
						offset := seedCp - m.fromRangeStart
						newSeed := m.toRangeStart + offset
						if seedCp >= m.fromRangeStart && seedCp <= m.fromRangeStart+m.len-1 {
							seedCp = newSeed
							continue mappingType
						}
					}
				}
				if seedCp < lowest || lowest == -1 {
					endSeeds <- seedCp
				}
			}
		}()

	}

	wg2 := sync.WaitGroup{}
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		for val := range endSeeds {
			fmt.Printf("%v\nchanLen: %v\n", val, len(endSeeds))
			if lowest == -1 {
				lowest = val

			}

			if lowest > val {
				lowest = val
			}
		}
	}()

	wg.Wait()
	close(endSeeds)
	wg2.Wait()

	fmt.Printf("Result:%+v", lowest)
}
