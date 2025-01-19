package go_fizz_buzz

import "strconv"

func FizzBuzz(n int) string {
	if 0 == n%5 {
		return "Buzz"
	}
	if 0 == n%3 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}
