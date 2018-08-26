package main

import (
	"fmt"
	"net/url"
)

func main() {
	m := url.Values{"lang": {"en"}} // direct construction
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // en
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // 1 (first value)
	fmt.Println(m["item"])     // [1, 2] (direct map access)

	// This could also be expressed as Values(nil).Get("item")
	// But not expressed as nil.Get("item"), because the type of nil is not known.
	m = nil
	fmt.Println(m.Get("item")) // ""
	m.Add("item", "3")         // panic assignment entry in nil map
}
