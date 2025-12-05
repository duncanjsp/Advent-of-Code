package main_test

import (
	"os"
	"testing"

	main "github.com/duncanjsp/Advent-of-Code/2025/day05"
)

func Test_Part1(t *testing.T) {
	raw, err := os.ReadFile("sample.txt")
	if err != nil {
		t.Fatalf("failed to read sample input: %v", err)
	}

	got, _ := main.Solve(string(raw))

	want := 0

	if got != want {
		t.Fatalf("Part 1 solution failed with %d, want %d", got, want)
	}
}

// func Test_Part2(t *testing.T) {
// 	raw, err := os.ReadFile("sample.txt")
// 	if err != nil {
// 		t.Fatalf("failed to read sample input: %v", err)
// 	}

// 	_, got := main.Solve(string(raw))

// 	want := 0

// 	if got != want {
// 		t.Fatalf("Part 2 solution failed with %d, want %d", got, want)
// 	}
// }
