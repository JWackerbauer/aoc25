package helpers

import (
	"fmt"
	"strconv"
)

// MustAtoi is exported because the first letter is capitalized.
func MustAtoi(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		// Use panic for non-recoverable errors
		panic(fmt.Sprintf("FATAL: Invalid integer value provided: %v", err))
	}
	return val
}
