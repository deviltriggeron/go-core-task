package contains

import (
	"fmt"
	"testing"
)

func TestContains1(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{6, 7, 8, 9, 10}

	res, ok := Contains(a, b)
	if !ok {
		fmt.Println("no matches")
		t.SkipNow()
	}
	fmt.Println(res)
}

func TestContains2(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 6, 7, 8, 9}

	res, ok := Contains(a, b)
	if !ok {
		fmt.Println("no matches")
		t.SkipNow()
	}
	fmt.Println(res)
}

func TestContains3(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	res, ok := Contains(a, b)
	if !ok {
		fmt.Println("no matches")
		t.SkipNow()
	}
	fmt.Println(res)
}
