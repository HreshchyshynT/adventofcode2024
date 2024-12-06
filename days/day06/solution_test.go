package day06

import (
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	input := strings.Split(`....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#..`, "\n")
	m, guardX, guardY := parseInput(input)

	t.Run("part1", func(t *testing.T) {
		want := 41
		got := part1(m, guardX, guardY)

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})
	t.Run("part2", func(t *testing.T) {
		want := 6
		got := part2(m, guardX, guardY)

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})
}
