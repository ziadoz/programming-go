// A slice cannot be used as a map key, since slices aren't directly comparable.
// Instead this can be by creating a map whose keys are a string of the entire slice.
package main

import "fmt"

func main() {
	// Add this slice once.
	slice1 := []string{"foo", "bar"}
	Add(slice1)

	// Then this one twice.
	slice2 := []string{"hello", "world"}
	Add(slice2)
	Add(slice2)

	// And then this one with emojis three times.
	slice3 := []string{"🇬🇧", "🇨🇦", "🇪🇺"}
	Add(slice3)
	Add(slice3)
	Add(slice3)

	// Check each one appears the correct number of times.
	fmt.Println("Slice 1 appears once", Count(slice1) == 3)
	fmt.Println("Slice 2 appears twice ", Count(slice2) == 2)
	fmt.Println("Slice 3 appears thrice", Count(slice3) == 3)

	// Dump out the map so we can see how each slice is represent in the map keys as a quoted string.
	fmt.Println("The map with slices as keys looks like this: \n", sliceMap)
}

// A map whose keys will be a quoted string that accurately represents a slice.
// Records the number of times each slice appears in the map.
var sliceMap = make(map[string]int)

// Returns a string of the entire slice which can be used as a map key.
// The string is quoted, so string boundaries are recorded faithfully,
func Key(list []string) string {
	return fmt.Sprintf("%q", list)
}

// Add a slice to the map, if it's already there we increment the total times it appears.
func Add(list []string) {
	sliceMap[Key(list)]++
}

// Return the total number of times a slice appears in the map.
func Count(list []string) int {
	return sliceMap[Key(list)]
}
