package main

import (
	"adventofcode2024/pkg/aoc"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("can't read .env: %v\n", err)
	}

	input, err := aoc.Get(1)
	if err != nil {
		log.Fatalf("can't get input: %v\n", err)
	}

	fmt.Printf("received input:\n%v", input)
}
