package day04

import (
	"adventofcode2024/pkg/aoc"
	"testing"
)

func TestExample(t *testing.T) {
	input, _ := aoc.Get(4)
	field := NewField(input)

	t.Run("part1", func(t *testing.T) {
		want := 2530
		got := part1(&field)

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})
	t.Run("part2", func(t *testing.T) {
		want := 1921
		got := part2(&field)

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})
}
