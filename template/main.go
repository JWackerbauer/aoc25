package main

import (
	"fmt"
	"os"
)

func solve(input string) int {
	var result = 0

	return result
}

func solve2(input string) int {
	var result = 0

	return result
}

var test_input = ``

func main() {
	input, err := os.ReadFile("dayxx/input.txt")
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
