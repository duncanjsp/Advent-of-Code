package main

import (
	"fmt"
	"os"
)

func main() {
	raw, err := os.ReadFile("day05/input.txt")
	if err != nil {
		panic(err)
	}

	p1, p2 := Solve(string(raw))
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}
