package word

import "testing"

func TestPalindrome(t *testing.T) {
	input1 := "detartrated"
	if !IsPalindrome(input1) {
		t.Errorf(`IsPalindrome("%q") = false`, input1)
	}

	input2 := "kayak"
	if !IsPalindrome(input2) {
		t.Errorf(`IsPalindrome(%q) = false`, input2)
	}
}

func TestNonPalindrome(t *testing.T) {
	input := "palindrome"
	if IsPalindrome(input) {
		t.Errorf(`IsPalindrome("%q") = true`, input)
	}
}

func TestFrenchPalindrome(t *testing.T) {
	input := "été"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}

func TestCanalPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
	if !IsPalindrome(input) {
		t.Errorf(`IsPalindrome(%q) = false`, input)
	}
}
