package main

import (
	"fmt"
	"unicode/utf8"
)

func banner(text string, width int) {
	// padding := (width - len(text)) / 2   BUG : len is in bytes
	padding := (width - utf8.RuneCountInString(text))

	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}

	fmt.Println(text)

	for i := 0; i < width; i++ {
		fmt.Print("-")
	}

	fmt.Println()
}

func isPalindrome(s string) bool {
	rs := []rune(s) // Get slice of runes out of s
	for i := 0; i < len(rs)/2; i++ {
		if rs[i] != rs[len(rs) - i - 1] {
			return false
		}
	}

	return true
}

func main() {
	// In Go, len return the no of bytes, not no of visible characters (runes)
	// Go is two bytes. G☺ is four as the unicode character is 3 bytes long. Hence, this may not be visually centered.
	banner("Go", 6)
	banner("G☺", 6)

	s := "G☺"

	for i, r := range s {
		fmt.Println(i, r)
	}

	// Bytes are essentially 8-bit values (alias for uint8). They represent raw data and are ideal for handling binary data or ASCII text.
	// Runes are 32-bit integers (alias for int32) that represent Unicode code points. They allow Go to work with a wide set of characters, beyond just the ASCII set.

	x, y := 1, "1"

	fmt.Printf("x = %v, y = %v\n", x, y)
	fmt.Printf("x = %#v, y = %#v\n", x, y) // Useful for debug logging

	fmt.Println()

	fmt.Println("g", isPalindrome("g"))
}

