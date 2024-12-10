package day10

import (
	"strings"
	"testing"
)

func TestExample(t *testing.T) {
	example := strings.Split(`
89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`, "\n")[1:]
	example1 := strings.Split(`
0123
1234
8765
9876`, "\n")[1:]
	example2 := strings.Split(`
9990999
9991999
9992999
6543456
7999997
8111198
9111119`, "\n")[1:]

	t.Run("part1", func(t *testing.T) {
		tests := []struct {
			input []string
			want  int
		}{
			{example, 36},
			{example1, 1},
			{example2, 3},
		}
		for _, test := range tests {
			want := test.want
			got := part1(parseInput(test.input))
			if want != got {
				t.Errorf("want %d got %d", want, got)
			}

		}
	})
}
