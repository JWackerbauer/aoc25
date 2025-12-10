package main

import (
	"fmt"
	"os"
	"strings"
)

func replaceAtIndex(in string, rep string, i int) string {
	out := []rune(in)
	out[i] = []rune(rep)[0]
	return string(out)
}

func solve(input string) int {
	var result = 0

	lines := strings.Split(input, "\n")
	for y := 1; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			char := string(lines[y][x])
			char_above := string(lines[y-1][x])
			if char_above == "S" || char_above == "|" {
				if char == "^" {
					// split right if there is space
					if x+1 < len(lines[y]) {
						lines[y] = replaceAtIndex(lines[y], "|", x+1)
					}
					//split left if there is space
					if x-1 >= 0 {
						lines[y] = replaceAtIndex(lines[y], "|", x-1)
					}
					result++
					continue
				}
				lines[y] = replaceAtIndex(lines[y], "|", x)
			}
		}
		fmt.Printf("%v\n", lines[y])
	}
	return result
}

// combine x,y into map key for memoization table
func hashCoords(x, y, width int) int {
	return y*width + x
}

// Memoization table
var memo = make(map[int]int)

func TachionParticleTimeline(x int, y_start int, lines []string) int {

	// Check memoization table for result
	key := hashCoords(x, y_start, len(lines[y_start]))
	if result, ok := memo[key]; ok {
		return result
	}

	// Check for splits in timeline
	for y := y_start; y < len(lines); y++ {

		// If the timeline splits; return the result of the split timelines
		if string(lines[y][x]) == "^" {
			result := TachionParticleTimeline(x+1, y+1, lines) +
				TachionParticleTimeline(x-1, y+1, lines)
			// Save the result to memoization table
			memo[key] = result
			return result
		}
	}

	// Return my own timeline
	return 1
}

func solve2(input string) int {

	var x_start int
	for x, char := range input {
		if string(char) == "S" {
			x_start = x
		}
	}

	result := TachionParticleTimeline(x_start, 0, strings.Split(input, "\n"))
	return result
}

var test_input = `.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`

func main() {
	input, err := os.ReadFile("day07/input.txt")
	if err != nil {
		fmt.Print(err)
	}
	var number = "test"
	if len(os.Args) > 1 {
		number = os.Args[1]
	}

	switch number {
	case "1":
		println(solve(string(input)))
	case "2":
		println(solve2(string(input)))
	default:
		println("test 1:")
		println(solve(test_input))
		println("test 2:")
		println(solve2(test_input))
	}
}
