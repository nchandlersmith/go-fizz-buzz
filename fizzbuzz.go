package go_fizz_buzz

import "strconv"

func FizzBuzz(n int) string {
	result := ""
	if 0 == n%3 {
		result += "Fizz"
	}
	if 0 == n%5 {
		result += "Buzz"
	}
	if result == "" {
		result = strconv.Itoa(n)
	}
	return result
}
