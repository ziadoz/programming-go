package main

import (
	"fmt"
	"strings"
)

func main() {
	// Anagrams.
	fmt.Println(anagram("murder", "redrum"))
	fmt.Println(anagram("listen", "silent"))
	fmt.Println(anagram("ƑƚǐƤ", "ƤƚǐƑ")) // UTF-8.

	// Not anagrams.
	fmt.Println(anagram("poop", "scoop"))
	fmt.Println(anagram("flip", "flop"))
	fmt.Println(anagram("ƓƕƔ", "ƩƝƔ"))    // UTF-8, shared one identical character.
	fmt.Println(anagram("foo", "foobar")) // Different lengths.
}

// Return whether or not two strings are anagrams of one another.
func anagram(s1, s2 string) bool {
	// Loop over the second string's runes.
	for _, r := range s2 {
		// If the rune is not within the first string, there can't be an anagram, so bail.
		if !strings.ContainsRune(s1, r) {
			return false
		}

		// The rune was found, so now remove it from the first string so we're keeping track.
		s1 = strings.Replace(s1, string(r), "", 1)
	}

	// If every character was removed from the first string, we have an anagram.
	return s1 == ""
}
