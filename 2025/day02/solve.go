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

			if halvesEqual(sti) {
				// Invalid IDs for pt1 also invalid for pt2
				part1 += i
				part2 += i
			} else if repeatingParts(sti) {
				part2 += i
			}

			i++
		}
	}
	return part1, part2
}

func halvesEqual(s string) bool {
	sl := len(s)
	if sl%2 != 0 {
		return false
	}
	return s[0:sl/2] == s[sl/2:]
}

func repeatingParts(s string) bool {
	sl := len(s)
	for i := 1; i <= sl/2; i++ {
		if sl%i != 0 {
			continue
		}
		part := s[0:i]
		if strings.Repeat(part, sl/i) == s {
			return true
		}
	}
	return false
}
