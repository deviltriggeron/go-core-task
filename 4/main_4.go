package main

import (
	"fmt"
	"slices"
)

func notContainingElem(s1 []string, s2 []string) []string {
	var res []string

	for _, l := range s1 {
		if !slices.Contains(s2, l) {
			res = append(res, l)
		}
	}

	return res
}

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	resSlice := notContainingElem(slice1, slice2)
	fmt.Println(resSlice)
}
