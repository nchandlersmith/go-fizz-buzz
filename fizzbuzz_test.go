package go_fizz_buzz

import "testing"

func TestFizzBuzz_returns_integer_as_string(t *testing.T) {
	want := "1"
	got := FizzBuzz(1)
	assertEquals(got, want, t)
}

func TestFizzBuzz_returns_Fizz_when_3(t *testing.T) {
	want := "Fizz"
	got := FizzBuzz(3)
	assertEquals(got, want, t)
}

func TestFizzBuzz_returns_Fizz_when_divisible_by_3(t *testing.T) {
	test_cases := []int{3, 6, 9, 12}
	for _, in := range test_cases {
		got := FizzBuzz(in)
		assertEquals(got, "Fizz", t)
	}
}

func assertEquals(got string, want string, t *testing.T) {
	if got != want {
		t.Errorf("Expected got: %v to equal want: %v", got, want)
	}
}
