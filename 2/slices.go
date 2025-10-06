package myslice

import (
	"fmt"
	"math/rand"
)

func OriginalSlice() []int {
	var res []int
	for i := 0; i < 10; i++ {
		res = append(res, rand.Intn(100))
	}
	return res
}

func SliceExample(s []int) []int {
	var res []int
	for _, j := range s {
		if j%2 == 0 {
			res = append(res, j)
		}
	}
	return res
}

func AddElements(s []int, target int) []int {
	var res []int

	res = append(res, s...)
	res = append(res, target)

	return res
}

func CopySlice(s []int) []int {
	res := s
	return res
}

func RemoveElement(s []int, removedIndex int) ([]int, error) {
	if len(s) <= removedIndex {
		return s, fmt.Errorf("removed index > len slice")
	}
	res := s

	copy(res[removedIndex:], res[removedIndex+1:])

	return res, nil
}
