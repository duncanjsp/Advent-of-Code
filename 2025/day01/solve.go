package main

import (
	"strconv"
	"strings"

	t "github.com/duncanjsp/Advent-of-Code/2025/tools"
)

func Solve(input string) (int, int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	part1 := 0
	part2 := 0

	pos := 50
	for _, l := range lines {
		dir := l[0:1]
		n, _ := strconv.Atoi(l[1:])
		switch dir {
		case "R":
			newPos, zeros := right(pos, n)
			pos = newPos
			part2 += zeros

		case "L":
			newPos, zeros := left(pos, n)
			pos = newPos
			part2 += zeros
		}
		if pos == 0 {
			part1++
		}
	}

	return part1, part2
}

func right(pos, n int) (new, pt2 int) {
	new = t.Mod(pos+n, 100)
	pt2 = (pos + n) / 100

	return new, pt2
}

func left(pos, n int) (new, pt2 int) {
	new = t.Mod(pos-n, 100)

	if n > pos {
		pt2++
		remainder := n - pos
		pt2 += remainder / 100
		if pos == 0 {
			pt2--
		}
	} else if new == 0 {
		pt2++
	}

	return new, pt2
}
