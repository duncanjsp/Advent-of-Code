package main

import (
	"strconv"
	"strings"
)

func Solve(input string) (int, int) {
	ranges := strings.Split(strings.TrimSpace(input), ",")
	part1 := 0
	part2 := 0

	for _, idRange := range ranges {
		a, b, ok := strings.Cut(idRange, "-")
		if !ok {
			panic(0)
		}

		i, _ := strconv.Atoi(a)
		j, _ := strconv.Atoi(b)
		for i <= j {
			sti := strconv.Itoa(i)
			// Only values with even number of digits can be invalid
			sl := len(sti)
			if sl%2 != 0 {
				i++
				continue
			}

			if sti[0:sl/2] == sti[sl/2:] {
				part1 += i
			}

			i++
		}
	}

	return part1, part2
}
