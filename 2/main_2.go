package main

import (
	"fmt"
	"math/rand"
)

func originalSlice() []int {
	var res []int
	for i := 0; i < 10; i++ {
		res = append(res, rand.Intn(100))
	}
	return res
}

func sliceExample(s []int) []int {
	var res []int
	for _, j := range s {
		if j%2 == 0 {
			res = append(res, j)
		}
	}
	return res
}

func addElements(s []int, target int) []int {
	var res []int

	res = append(res, s...)
	res = append(res, target)

	return res
}

func copySlice(s []int) []int {
	res := s
	return res
}

func removeElement(s []int, removedIndex int) []int {
	res := s

	copy(res[removedIndex:], res[removedIndex+1:])

	return res
}

func main() {
	s1 := originalSlice()
	fmt.Println(s1)

	s2 := sliceExample(s1)
	fmt.Println(s2)

	s3 := addElements(s2, 87)
	fmt.Println(s3)

	s4 := copySlice(s3)
	s4 = addElements(s4, 111)
	fmt.Printf("original: %d\n", s3)
	fmt.Printf("copy: %d\n", s4)

	fmt.Printf("before: %d\n", s4)
	s4 = removeElement(s4, 2)
	fmt.Printf("after: %d\n", s4)
}
