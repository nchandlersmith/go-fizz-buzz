package go_fizz_buzz

import "testing"

func TestFizzBuzz_returns_integer_as_string(t *testing.T) {
	want := "1"
	got := FizzBuzz(1)
	if got != want {
		t.Errorf("got: %v want: %v", got, want)
	}
}
