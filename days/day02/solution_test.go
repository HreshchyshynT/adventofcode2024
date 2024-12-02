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
}
