package greeting

import "testing"

func TestGreetingWithName(t *testing.T) {
	got := Greeting("John")
	expect := "Hello, John"

	if got != expect {
		t.Errorf("Did not get the expected result. Wanted %s got %s", got, expect)
	}
}

func TestGreetingWithoutName(t *testing.T) {
	got := Greeting("")
	expect := "Hello, "

	if got != expect {
		t.Errorf("Did not get the expected result. Wanted %s got %s", got, expect)
	}
}

func TestAddTwoNumbers(t *testing.T) {
	got := AddTwoNumbers(1, 2)
	expect := 3

	if got != expect {
		t.Errorf("Did not get the expected result. Wanted %d got %d", got, expect)
	}
}
