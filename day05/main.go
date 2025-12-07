package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/JWackerbauer/aoc25/helpers"
)

type IdRange struct {
	lower_bound int
	upper_bound int
	deleted     bool
}

func Process(input string) ([]IdRange, []int) {
	split_input := strings.Split(input, "\n\n")
	var ingredients []int
	for ingredient := range strings.SplitSeq(split_input[1], "\n") {
		ingredient_int, _ := strconv.Atoi(string(ingredient))
		ingredients = append(ingredients, ingredient_int)
	}
	return CreateSortedRanges(split_input[0]), ingredients
}

func CreateSortedRanges(input string) []IdRange {
	var id_ranges []IdRange
	// Collect all ranges
	for line := range strings.SplitSeq(input, "\n") {
		bounds := strings.Split(line, "-")

		new_range := IdRange{
			lower_bound: helpers.MustAtoi(bounds[0]),
			upper_bound: helpers.MustAtoi(bounds[1]),
		}
		id_ranges = append(id_ranges, new_range)
	}
	fmt.Printf("unsorted: %v\n", id_ranges)
	// Sort ranges
	slices.SortFunc(id_ranges, func(a, b IdRange) int {
		return a.lower_bound - b.lower_bound
	})
	fmt.Printf("sorted: %v\n", id_ranges)
	// Coalesce; merge overlapping ranges
	for current := 1; current < len(id_ranges); current++ {
		// Skip deleted ranges
		previous_undeleted := current - 1
		for j := previous_undeleted; j >= 0; j-- {
			if !id_ranges[j].deleted {
				previous_undeleted = j
				break
			}
		}
		// Check for intersection or adjacency
		if id_ranges[current].lower_bound <= id_ranges[previous_undeleted].upper_bound+1 {
			// If current range is not fully contained, extend previous
			if id_ranges[current].upper_bound > id_ranges[previous_undeleted].upper_bound {
				id_ranges[previous_undeleted].upper_bound = id_ranges[current].upper_bound
			}
			// Mark for deletion
			id_ranges[current].deleted = true
		}
	}
	fmt.Printf("coalesced: %v\n", id_ranges)
	// Contract; delete marked ranges
	id_ranges = slices.DeleteFunc(id_ranges, func(n IdRange) bool {
		return n.deleted
	})
	fmt.Printf("contracted: %v\n", id_ranges)
	return id_ranges
}

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
