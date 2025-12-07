package main

import (
	"fmt"
	"os"
	"strings"
)

func check(target byte) int {
	if string(target) == "@" {
		return 1
	}
	return 0
}

func can_be_removed(x int, y int, lines []string) bool {
	adjacent := 0
	line := lines[y]
	//left
	if x-1 >= 0 {
		adjacent += check(line[x-1])
		//topleft
		if y-1 >= 0 {
			adjacent += check(lines[y-1][x-1])
		}
		//bottomleft
		if y+1 < len(lines) {
			adjacent += check(lines[y+1][x-1])
		}
	}
	//right
	if x+1 < len(line) {
		adjacent += check(line[x+1])
		//topright
		if y-1 >= 0 {
			adjacent += check(lines[y-1][x+1])
		}
		//bottomright
		if y+1 < len(lines) {
			adjacent += check(lines[y+1][x+1])
		}
	}
	//top
	if y-1 >= 0 {
		adjacent += check(lines[y-1][x])
	}
	//bottom
	if y+1 < len(lines) {
		adjacent += check(lines[y+1][x])
	}
	return adjacent < 4
}

func solve(input string) int {
	var result = 0

	lines := strings.Split(input, "\n")

	for y, line := range lines {
		for x, pos := range line {

			if string(pos) == "@" {
				if can_be_removed(x, y, lines) {
					result++
				}
			}
		}
	}
	return result
}

func remove_rolls(input string, result int) int {
	after := ""
	lines := strings.Split(input, "\n")

	for y, line := range lines {
		for x, pos := range line {
			if string(pos) == "@" && can_be_removed(x, y, lines) {
				result++
				after += "x"
			} else {
				after += string(pos)
			}
		}
		if y < len(lines)-1 {
			after += "\n"
		}
	}
	fmt.Println(result)
	fmt.Println(after)
	if after == input {
		return result
	}
	return remove_rolls(after, result)
}

func solve2(input string) int {
	var result = 0
	fmt.Println(input)
	return remove_rolls(input, result)
}

var test_input = `..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`

func main() {
	input, err := os.ReadFile("day04/input.txt")
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
