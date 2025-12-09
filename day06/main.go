package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/JWackerbauer/aoc25/helpers"
)

type Operand int

const (
	Multiplication Operand = iota
	Addition
	NotAnOperand
)

func ToOperand(symbol string) Operand {
	switch symbol {
	case "*":
		return Multiplication
	case "+":
		return Addition
	default:
		return NotAnOperand
	}
}

type Column struct {
	numbers []int
	operand Operand
}

func (col Column) AddCell(value string) Column {
	if value == "*" || value == "+" {
		col.operand = ToOperand(value)
		return col
	}
	col.numbers = append(col.numbers, helpers.MustAtoi(value))
	return col
}

func (col Column) Calculate() int {
	result := col.numbers[0]
	for i := 1; i < len(col.numbers); i++ {
		num := col.numbers[i]
		switch col.operand {
		case Multiplication:
			result *= num
		case Addition:
			result += num
		default:
			panic(fmt.Sprintf("FATAL: Invalid Operand enum value: %v", col.operand))
		}
	}
	return result
}

func ProcessCephalopod(input string) []Column {
	var columns []Column

	lines := strings.Split(input, "\n")

	// Initialize the first column,
	// We will catch NotAnOperand later
	// as soon as we find the first one
	current_col := Column{operand: NotAnOperand}

	for x := 0; x < len(lines[0]); x++ {
		number := ""
		for _, line := range lines {
			current := string(line[x])
			operand := ToOperand(current)
			// We found a new column
			if operand != NotAnOperand {
				// Handle edge case for first column
				if current_col.operand == NotAnOperand {
					current_col.operand = operand
					continue
				}
				columns = append(columns, current_col)
				current_col = Column{operand: operand}
				continue
			}
			if current != " " {
				number += current
			}
		}
		if number == "" {
			continue
		}
		current_col.numbers = append(current_col.numbers, helpers.MustAtoi(number))
	}

	// Add last column
	if current_col.operand != NotAnOperand {
		columns = append(columns, current_col)
	}

	return columns
}

// Get the nth digit of a number
// n is 0 indexed, starting with the most significant digit
// Returns -1 if n >= number of digits in the number
func DigitN(num, n int) int {
	stringrep := strconv.Itoa(num)
	if len(stringrep) <= n {
		return -1
	}
	return helpers.MustAtoi(string(stringrep[n]))
}

func Process(input string) []Column {
	var columns []Column

	for line := range strings.SplitSeq(input, "\n") {
		row := strings.Split(line, " ")
		i := 0
		// Trim all the spaces
		for _, cell := range row {
			cell = strings.Trim(cell, " ")
			// Skip whitespace
			if cell == "" {
				continue
			}
			// Add a column if needed
			if i >= len(columns) {
				columns = append(columns, Column{}.AddCell(cell))
				i++
				continue
			}
			columns[i] = columns[i].AddCell(cell)
			i++
		}
	}
	return columns
}

func solve(input string) int {
	var result = 0

	columns := Process(input)

	for _, col := range columns {
		result += col.Calculate()
	}

	return result
}

func solve2(input string) int {
	var result = 0

	columns := ProcessCephalopod(input)

	for _, col := range columns {
		result += col.Calculate()
	}

	return result
}

var test_input = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

func main() {
	input, err := os.ReadFile("day06/input.txt")
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
