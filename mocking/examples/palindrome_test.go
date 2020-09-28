package example

import "testing"

func TestPalindrome(t *testing.T) {
	// unit test for Palindrome function
	// an example of a table-driven test
	cases := []struct {
		word string
		want bool
	}{
		{"wrong", false},
		{"racecar", true},
		{"Anna", true},
		{"Москва", false},
		{"манекенам", true},
	}
	for _, c := range cases {
		got := Palindrome(c.word)
		if got != c.want {
			t.Errorf("Palindrome(%q) == %v, want %v", c.word, got, c.want)
		}
	}
}
