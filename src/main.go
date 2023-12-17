package main

import (
	"Christmas-2023/src/day1"
	"Christmas-2023/src/day10"
	"Christmas-2023/src/day11"
	"Christmas-2023/src/day13"
	"Christmas-2023/src/day14"
	"Christmas-2023/src/day15"
	"Christmas-2023/src/day17"
	"Christmas-2023/src/day2"
	"Christmas-2023/src/day3"
	"Christmas-2023/src/day4"
	"Christmas-2023/src/day5"
	"Christmas-2023/src/day9"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func printDay(day int, fns ...func()) {
	fmt.Printf("***************Day  %v********************\n", day)
	for index, task := range fns {
		fmt.Printf("---------------Task %v-------------------\n", index+1)
		task()
		fmt.Println()
	}
}

func main() {
	go func() {
		log.Fatal(http.ListenAndServe(":7777", nil))
	}()

	printDay(1, day1.TaskOne, day1.TaskTwo)
	printDay(2, day2.TaskOne, day2.TaskTwo)
	printDay(3, day3.TaskOne)
	printDay(4, day4.TaskOne, day4.TaskTwo)
	printDay(5, day5.TaskOne) //, day5.TaskTwo)
	printDay(9, day9.TaskOne, day9.TaskTwo)
	printDay(10, day10.TaskOne)
	printDay(11, day11.TaskOne, day11.TaskTwo)
	//printDay(12, day12.TaskOne)
	printDay(13, day13.TaskOne, day13.TaskTwo)
	printDay(14, day14.TaskOne, day14.TaskTwo)
	printDay(15, day15.TaskOne, day15.TaskTwo)
	//printDay(16, day16.TaskOne, day16.TaskTwo)
	printDay(17, day17.TaskOne)
}
