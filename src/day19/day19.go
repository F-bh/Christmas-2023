package day19

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type state struct {
	tasks map[string]task
	rejected,
	accepted []part
}

type part = map[rune]int

type task func(s *state, p part)

func parseTask(input string) (string, task) {
	name := strings.Split(input, "{")[0]

	trimmed := strings.Trim(input, "{}")

	filters := strings.Split(trimmed, ",")

	return name, func(s *state, p part) {
		for index, filter := range filters {
			if index == len(filters)-1 {
				switch filter {
				case "A":
					s.accepted = append(s.accepted, p)
				case "B":
					s.rejected = append(s.rejected, p)
				default:
					log.Fatalf("invalid destination: %v", filter)
				}
			}

			comparison, err := strconv.Atoi(strings.Split(filter[2:], ":")[0])
			destination := strings.Split(filter, ":")[1]

			if err != nil {
				log.Fatalf("couldnt parse comparison string %v\n to: %v", filter, strings.Split(filter, ":")[1])
			}

			val, ok := p[rune(filter[0])]
			if !ok {
				log.Fatalf("couldnt find part attribute: %v\nin:%v", string(filter[0]), s.tasks)
			}

			if filter[1] == '>' {
				if val > comparison {
					s.tasks[destination](s, p)
					return
				}
				continue
			}

			if filter[1] == '<' {
				if val < comparison {
					s.tasks[destination](s, p)
					return
				}
				continue
			}

			log.Fatalf("no comparison found in filter: %v", filter)
		}
		log.Fatalf("how did we end up here?")
	}
}

func TaskOne() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("couldnt get pwd\nwith err: %v", err.Error())
	}

	inputFile, err := os.ReadFile(dir + "/src/day19/input")
	if err != nil {
		log.Fatalf("failed to open file input\nwith err: %v", err)
	}

	split := strings.Split(string(inputFile), "\n\n")

	rules, items := strings.Split(split[0], "\n"), strings.Split(split[1], "\n")

	s := state{
		tasks: make(map[string]task),
	}

	for _, rule := range rules {
		name, f := parseTask(rule)
		s.tasks[name] = f
	}

	for _, item := range items {
		clean := strings.Trim(item, "{}")

		part := make(part)
		for _, entry := range strings.Split(clean, ",") {
			val, err := strconv.Atoi(entry[3:])
			if err != nil {
				log.Fatalf("failed to parse part value for entry: %v", entry)
			}
			part[rune(entry[0])] = val
		}

		for _, task := range s.tasks {
			task(&s, part)
		}
	}
}
