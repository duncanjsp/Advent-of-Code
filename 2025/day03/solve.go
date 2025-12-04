package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve(input string) (int, int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	part1 := 0
	part2 := 0

	for _, l := range lines {
		part1 += maxJoltage(l, 2)
		part2 += maxJoltage(l, 12)
	}

	return part1, part2
}

// Return the maximum joltage possible in the
// bank of batteries using n batteries.
func maxJoltage(bank string, n int) int {
	ba := maxJoltageBytes(bank, n)
	ans, err := strconv.Atoi(string(ba))
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return ans
}

// Return the maximum joltage possible in the
// bank of batteries using n batteries as a
// byte array.
func maxJoltageBytes(bank string, n int) []byte {
	bats := len(bank)
	if n > bats {
		fmt.Printf("%s\nCannot activate %d batteries. Bank contains only %d batteries.\n", bank, n, len(bank))
		return []byte{}
	}
	a := bank[0]
	r := 1
	for l := 1; l <= bats-n; l++ {
		if bank[l] > a {
			a = bank[l]
			r = l + 1
		}
	}

	if n == 1 {
		return []byte{a}
	}

	return append([]byte{a}, maxJoltageBytes(bank[r:], n-1)...)
}
