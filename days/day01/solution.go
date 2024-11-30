package day01

import (
	"adventofcode2024/pkg/aoc"
	"log"
)

func Solve(part int) {
	log.Printf("solving 1st day, part: %d", part)
	input, err := aoc.Get(1)
	if err != nil {
		log.Fatal("error getting input for day 1")
	}
	switch part {
	case 1:
		part1(input)
	case 2:
		part2(input)
	}

}

func part1(input []string) {

}

func part2(input []string) {

}
