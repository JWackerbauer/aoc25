package main

import (
	"fmt"
	"os"
	"strings"
)

func process(input string) {
	var machines []Machine
	for line := range strings.Lines(input) {
		machines = append(machines, CreateMachine(line))
	}

}

func solve(input string) int {
	var result = 0
	process(input)
	return result
}

func solve2(input string) int {
	var result = 0

	return result
}

var test_input = `[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}`

func main() {
	input, err := os.ReadFile("day10/input.txt")
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
