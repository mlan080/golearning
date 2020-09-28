package example

import "strings"

// Palindrome returns true if the word reads the same backward as forward, such as racecar or Anna.
func Palindrome(word string) bool {
	word = strings.ToLower(word)
	return word == Reverse(word)
}
