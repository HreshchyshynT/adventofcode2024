package main

import (
	"adventofcode2024/days/day01"
	"adventofcode2024/days/day02"
	"adventofcode2024/days/day03"
	"adventofcode2024/days/day04"
	"adventofcode2024/pkg/aoc"
	"flag"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("can't read .env: %v\n", err)
	}

	day := flag.Int("day", 1, "day to solve")
	flag.Parse()

	input, err := aoc.Get(*day)
	if err != nil {
		log.Fatal("error getting input for day 1")
	}

	switch *day {
	case 1:
		day01.Solve(input)
	case 2:
		day02.Solve(input)
	case 3:
		day03.Solve(input)
	case 4:
		day04.Solve(input)
	default:
		log.Printf("Day %d not solved yet", *day)
	}
}
