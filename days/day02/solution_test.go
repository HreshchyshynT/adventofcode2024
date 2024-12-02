package day02

import (
	"testing"
)

func TestExampleInput(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		input := parseInput([]string{"7 6 4 2 1",
			"1 2 7 8 9",
			"9 7 6 2 1",
			"1 3 2 4 5",
			"8 6 4 4 1",
			"1 3 6 7 9",
		})
		want := 2
		got := part1(input)

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})
	t.Run("part2", func(t *testing.T) {
		input := parseInput([]string{
			"48 46 47 49 51 54 56",
			"1 1 2 3 4 5",
			"1 2 3 4 5 5",
			"5 1 2 3 4 5",
			"1 4 3 2 1",
			"1 6 7 8 9",
			"1 2 3 4 3",
			"9 8 7 6 7",
			"7 10 8 10 11",
			"29 28 27 25 26 25 22 20"})
		want := 10
		got := part2(input)

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})
}
