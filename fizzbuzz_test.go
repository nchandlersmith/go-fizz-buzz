package go_fizz_buzz

import "testing"

func TestFizzBuzz_1(t *testing.T) {
	want := "1"
	got := FizzBuzz(1)
	if got != want {
		t.Errorf("TestFizzBuzz(1) = %v want %v", got, want)
	}
}
