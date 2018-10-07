// https://stackoverflow.com/questions/28999735/what-is-the-shortest-way-to-simply-sort-an-array-of-structs-by-arbitrary-field
package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"text/tabwriter"
)

type Hobbit struct {
	Name string
	Age  int
}

// Returns a Less() compatible sort function that compares on each of the fields specified using reflection.
func SortFields(fields ...string) func(h1, h2 *Hobbit) bool {
	return func(h1, h2 *Hobbit) bool {
		for _, field := range fields {
			v1 := reflect.Indirect(reflect.ValueOf(h1)).FieldByName(field)
			v2 := reflect.Indirect(reflect.ValueOf(h2)).FieldByName(field)

			switch v1.Kind() {
			case reflect.Int:
				i1 := int(v1.Int())
				i2 := int(v2.Int())
				if i1 != i2 {
					return i1 < i2
				}
			case reflect.String:
				s1 := string(v1.String())
				s2 := string(v2.String())
				if s1 != s2 {
					return s1 < s2
				}
			}
		}

		return false
	}
}

// A Less() compatible function type. A helper type that can be used for sorting.
type By func(h1, h2 *Hobbit) bool

func (by By) Sort(hobbits []Hobbit) {
	hs := &HobbitSorter{
		Hobbits:  hobbits,
		SortFunc: by,
	}

	sort.Sort(hs)
}

// A type for sorting. Accepts the objects to sort and a sorting funciton.
type HobbitSorter struct {
	Hobbits  []Hobbit
	SortFunc func(h1, h2 *Hobbit) bool
}

func (hs *HobbitSorter) Len() int {
	return len(hs.Hobbits)
}

func (hs *HobbitSorter) Swap(i, j int) {
	hs.Hobbits[i], hs.Hobbits[j] = hs.Hobbits[j], hs.Hobbits[i]
}

func (hs *HobbitSorter) Less(i, j int) bool {
	return hs.SortFunc(&hs.Hobbits[i], &hs.Hobbits[j])
}

func PrintHobbits(hobbits []Hobbit) {
	const format = "%s\t%d\t\n"
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	fmt.Fprintf(tw, "%s\t%s\t\n", "Name", "Age")
	fmt.Fprintf(tw, "%s\t%s\t\n", "----", "---")
	for _, hobbit := range hobbits {
		fmt.Fprintf(tw, format, hobbit.Name, hobbit.Age)
	}
	tw.Flush()
}

func main() {
	hobbits := []Hobbit{
		{Name: "Frodo", Age: 33},
		{Name: "Bilbo", Age: 111},
		{Name: "Samwise", Age: 21},
		{Name: "Merry", Age: 16},
		{Name: "Pippin", Age: 16},
	}

	By(SortFields("Age", "Name")).Sort(hobbits)
	PrintHobbits(hobbits)
}
