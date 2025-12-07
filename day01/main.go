package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maxTicks = 100
var startPosition = 50

func solve(input string) int {
	var lines = strings.Split(input, "\n")
	var position = startPosition
	var result = 0
	for _, line := range lines {
		var direction = line[0]
		var ticks, _ = strconv.Atoi(line[1:])
		switch direction {
		case 'L':
			fmt.Printf("Turn left %v ticks\n", ticks)
			position -= ticks
			position %= maxTicks
		case 'R':
			fmt.Printf("Turn right %v ticks\n", ticks)
			position += ticks
			position %= maxTicks
		default:
			println("Should not happen")
		}
		fmt.Printf("position: %v \n", position)
		if position == 0 {
			result++
			fmt.Printf("0 found, new result: %v \n", result)
		}
	}
	return result
}

func solve2(input string) int {
	var lines = strings.Split(input, "\n")
	var position = startPosition
	var result = 0

	for _, line := range lines {
		var direction = line[0]
		var ticks, _ = strconv.Atoi(line[1:])

		var full_rotations = ticks / maxTicks
		var remainder = ticks % maxTicks

		if full_rotations > 0 {
			result += full_rotations
			fmt.Printf("%v full rotations, result incremented to %v\n", full_rotations, result)
		}

		switch direction {
		case 'L':
			fmt.Printf("Turn left %v ticks\n", remainder)
			if position == 0 {
				position = maxTicks
			}
			position -= remainder

			if position < 0 {
				result++
				fmt.Printf("Crossed 0, result incremented to: %v\n", result)
				position += maxTicks
			}
		case 'R':
			fmt.Printf("Turn right %v ticks\n", remainder)
			if position == maxTicks {
				position = 0
			}
			position += remainder

			if position > maxTicks {
				result++
				fmt.Printf("Crossed 0, result incremented to: %v\n", result)
				position -= maxTicks
			}
		default:
			println("Should not happen")
		}
		fmt.Printf("position: %v \n", position)
		if position == 0 || position == 100 {
			result++
			fmt.Printf("0 found, new result: %v \n", result)
		}
	}
	return result
}

var test_input = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`

func main() {
	input, err := os.ReadFile("day01/input.txt")
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
