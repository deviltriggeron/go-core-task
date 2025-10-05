package main

import (
	"fmt"
	"slices"
)

func contains(s1 []int, s2 []int) ([]int, bool) {
	var res []int
	status := false

	for _, l := range s1 {
		if slices.Contains(s2, l) {
			res = append(res, l)
			status = true
		}
	}

	return res, status
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	s, ok := contains(a, b)
	if ok {
		fmt.Println(s)
	}
}
