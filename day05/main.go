package main

import (
	"fmt"
	"os"
	"sort"
)

func solve(input string) int {
	var result = 0

	id_ranges, ingredients := Process(input)

	for _, ingredient := range ingredients {
		fmt.Printf("Ingredient %v", ingredient)
		i := sort.Search(len(id_ranges), func(i int) bool {
			return id_ranges[i].upper_bound >= ingredient
		})
		if i < len(id_ranges) && id_ranges[i].lower_bound <= ingredient {
			fmt.Printf(" ...found in %v-%v", id_ranges[i].lower_bound, id_ranges[i].upper_bound)
			result++
		}
		fmt.Printf("\n")
	}

	return result
}

func solve2(input string) int {
	var result = 0

	id_ranges, _ := Process(input)

	for _, id_range := range id_ranges {
		result += id_range.upper_bound - id_range.lower_bound + 1
	}

	return result
}

var test_input = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

func main() {
	input, err := os.ReadFile("day05/input.txt")
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
