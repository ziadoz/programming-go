package main

import (
	"sort"
)

type Palindrome []rune

func (p Palindrome) Len() int {
	return len(p)
}

func (p Palindrome) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p Palindrome) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Palindrome) String() string {
	return string(p)
}

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if eq := !s.Less(i, j) && !s.Less(j, i); !eq {
			return false
		}
	}
	return true
}

func main() {
	values := []Palindrome{
		Palindrome([]rune("madam")),
		Palindrome([]rune("palindrome")),
	}

	for _, value := range values {
		println(value.String(), IsPalindrome(value))
	}
}
