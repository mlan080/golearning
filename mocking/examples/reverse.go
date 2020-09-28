package example

// Reverse returns the reverse of its input. Example: "word" becomes "drow".
func Reverse(word string) string {
	var reverse []rune
	for _, r := range word {
		reverse = append(reverse, r)
	}
	n := len(reverse)
	for i := 0; i < n/2; i++ {
		reverse[i], reverse[n-1-i] = reverse[n-1-i], reverse[i]
	}
	return string(reverse)
}
