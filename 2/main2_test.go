package myslice

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	s := OriginalSlice()
	fmt.Println(s)
}

func TestExample(t *testing.T) {
	s := OriginalSlice()
	fmt.Print("before")
	fmt.Println(s)

	res := SliceExample(s)
	fmt.Print("after")
	fmt.Println(res)
}

func TestAddElements(t *testing.T) {
	s := OriginalSlice()
	fmt.Print("before")
	fmt.Println(s)

	s = AddElements(s, 155)
	fmt.Print("after")
	fmt.Println(s)
}

func TestCopy(t *testing.T) {
	s := OriginalSlice()
	fmt.Print("original")
	fmt.Println(s)

	ns := CopySlice(s)
	ns = AddElements(ns, 112)
	fmt.Print("copy after add")
	fmt.Println(ns)
}

func TestRemove(t *testing.T) {
	s := OriginalSlice()
	fmt.Print("before")
	fmt.Println(s)

	s, err := RemoveElement(s, 4)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print("after")
	fmt.Println(s)
}
