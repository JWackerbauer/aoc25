package main

import (
	"fmt"
	"slices"
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
