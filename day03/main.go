package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(input string) int {
	var result = 0

	var lines = strings.Split(input, "\n")

	for _, line := range lines {
		fmt.Printf("%v ", line)
		var digit_1 = -1
		var highest_num_index int
		// Find the first number
		for i, num_rune := range line[:len(line)-1] {
			var num, _ = strconv.Atoi(string(num_rune))
			if num > digit_1 {
				digit_1 = num
				highest_num_index = i
			}
		}
		fmt.Printf("digit_1: %v index: %v ", digit_1, highest_num_index)
		var digit_2 = -1
		// Find the second number
		for _, num_rune := range line[highest_num_index+1:] {
			var num, _ = strconv.Atoi(string(num_rune))
			if num > digit_2 {
				digit_2 = num
			}
		}
		fmt.Printf("digit_2: %v ", digit_2)

		var joltage, _ = strconv.Atoi(fmt.Sprintf("%v%v", digit_1, digit_2))
		fmt.Printf("joltage: %v\n", joltage)
		result += joltage
	}
	return result
}

func solve2(input string) int {
	var result = 0
	const num_batteries = 12

	var lines = strings.Split(input, "\n")

	for _, line := range lines {
		fmt.Printf("%v ", line)

		var digits string
		digit := -1
		digit_pos := 0
		for len(digits) < num_batteries {
			max_digit_pos := len(line) - (num_batteries - 1 - len(digits))
			fmt.Printf("digit_pos: %v max_digit_pos:%v\n", digit_pos, max_digit_pos)
			next_digit_i := 0
			fmt.Printf("considered: %v\n", line[digit_pos:max_digit_pos])
			for i, num_rune := range line[digit_pos:max_digit_pos] {
				var num, _ = strconv.Atoi(string(num_rune))
				if num > digit {
					digit = num
					next_digit_i = i
				}
			}
			digit_pos += next_digit_i + 1
			fmt.Printf("digit: %v digit_pos: %v\n", digit, digit_pos)
			digits += strconv.Itoa(digit)
			digit = -1
		}
		fmt.Printf("digits: %v\n", digits)
		number, _ := strconv.Atoi(digits)
		result += number
	}
	return result
}

var test_input = `987654321111111
811111111111119
234234234234278
818181911112111`

func main() {
	input, err := os.ReadFile("day03/input.txt")
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
