package main

import (
	"strconv"
	"strings"
)

func Solve(input string) (int, int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	part1 := 0
	part2 := 0

	for _, l := range lines {
		part1 += largestJoltage(l)
	}

	return part1, part2
}

func largestJoltage(bank string) int {
	a := bank[0]
	r := 1
	for l := 1; l < len(bank)-1; l++ {
		if bank[l] > a {
			a = bank[l]
			r = l + 1
		}
	}
	b := bank[r]
	for ; r < len(bank); r++ {
		if bank[r] > b {
			b = bank[r]
		}
	}
	ba := []byte{a, b}
	ans, _ := strconv.Atoi(string(ba))
	return ans
}
