package go_fizz_buzz

import "strconv"

func FizzBuzz(n int) string {
	if 3 == n {
		return "Fizz"
	}
	return strconv.Itoa(n)
}
