package main

import (
	"adventofcode2024/days/day01"
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
	part := flag.Int("part", 1, "part to solve")
	flag.Parse()

	switch *day {
	case 1:
		day01.Solve(*part)
	default:
		log.Printf("Day %d, part: %d not solved yet", *day, *part)
	}
}
