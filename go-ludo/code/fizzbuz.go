package kata

import (
	"fmt"
	"strconv"
)

// FizzBuzz main function
func FizzBuzz(max int) string {
	fmt.Println("Kata FizzBuzz")

	var result string

	for i := 1; i < max; i++ {
		result += Display(i) + "\n"
	}
	return result
}

// Display number
func Display(number int) string {

	fizz := IsMultipleOf(3, number)
	buzz := IsMultipleOf(5, number)

	if fizz && buzz {
		return "FizzBuzz"
	}
	if fizz {
		return "Fizz"
	}
	if buzz {
		return "Buzz"
	}

	return strconv.Itoa(number)
}

// IsMultipleOf divisor for number
func IsMultipleOf(divisor int, number int) bool {
	return number%divisor == 0
}
