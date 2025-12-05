package main

import (
	"strings"
)

func Solve(input string) (int, int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	part1 := 0
	part2 := 0
	bogRoll := byte('@')
	for y, l := range lines {
		for x, c := range []byte(l) {
			if c != bogRoll {
				continue
			}
			n := 0
			for yo := -1; yo <= 1; yo++ {
				for xo := -1; xo <= 1; xo++ {
					yc, xc := y+yo, x+xo
					if yc < 0 || yc >= len(lines) || xc < 0 || xc >= len(l) || (yo == 0 && xo == 0) {
						continue
					}
					if lines[yc][xc] == bogRoll {
						n++
					}
				}
			}
			if n < 4 {
				part1++
			}
		}
	}
	return part1, part2
}
