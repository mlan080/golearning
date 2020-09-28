package example

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssertPalindrome(t *testing.T) {
	// unit test for Palindrome function
	// an example of a table-driven test using the testify/assert package
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
		// using assert instead of an "if" statement:
		assert.Equal(t, c.want, got)
	}
}
