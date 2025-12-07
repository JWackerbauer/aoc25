package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func solve(input string) int {
	var result = 0

	var id_ranges = strings.Split(input, ",")

	for _, id_range := range id_ranges {
		var bounds = strings.Split(id_range, "-")
		var lower_bound, _ = strconv.Atoi(bounds[0])
		var upper_bound, _ = strconv.Atoi(bounds[1])
		var invalid_amount = 0
		for i := lower_bound; i <= upper_bound; i++ {
			var stringrep = strconv.Itoa(i)

			if len(stringrep)%2 != 0 {
				// ignore asymmetric strings
				continue
			}
			var midpoint = len(stringrep) / 2

			if stringrep[0:midpoint] == stringrep[midpoint:] {
				invalid_amount++
				result += i
			}
		}

		fmt.Printf("%v %v - %v %v: %v\n",
			bounds[0][0:len(bounds[0])/2],
			bounds[0][len(bounds[0])/2:],
			bounds[1][0:len(bounds[1])/2],
			bounds[1][len(bounds[1])/2:],
			invalid_amount,
		)
	}
	return result
}

func solve2(input string) int {
	var result = 0

	var id_ranges = strings.Split(input, ",")

	for _, id_range := range id_ranges {
		var bounds = strings.Split(id_range, "-")
		var lower_bound, _ = strconv.Atoi(bounds[0])
		var upper_bound, _ = strconv.Atoi(bounds[1])
		for id := lower_bound; id <= upper_bound; id++ {
			var id_string = strconv.Itoa(id)
			// Check for repeating patterns;
			// patterns can not be longer than half of the input
			for pattern_length := 1; pattern_length <= len(id_string)/2; pattern_length++ {
				// The input length has to be divisible by pattern_length
				if len(id_string)%pattern_length == 0 {
					// Repeat the pattern and see if it matches the input
					if strings.Repeat(id_string[0:pattern_length], len(id_string)/pattern_length) == id_string {
						fmt.Printf("%v -- %v\n", id_string, id_string[0:pattern_length])
						result += id
						// Break to avoid multiple matches for the same input
						break
					}
				}
			}
		}
	}
	return result
}

var test_input = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

func main() {
	input, err := os.ReadFile("day02/input.txt")
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
