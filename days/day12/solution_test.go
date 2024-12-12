package day12

import (
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	t.Run("part1", func(t *testing.T) {
		input := strings.Split(`
AAAA
BBCD
BBCC
EEEC`, "\n")[1:]
		want := 140
		got := part1(parseInput(input))

		if want != got {
			t.Errorf("want %d got %d", want, got)
		}
	})
}
