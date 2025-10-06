package main

import (
	"fmt"
	s "myslice"
)

func main() {
	s1 := s.OriginalSlice()
	fmt.Println(s1)

	s2 := s.SliceExample(s1)
	fmt.Println(s2)

	s3 := s.AddElements(s2, 87)
	fmt.Println(s3)

	s4 := s.CopySlice(s3)
	s4 = s.AddElements(s4, 111)
	fmt.Printf("original: %d\n", s3)
	fmt.Printf("copy: %d\n", s4)

	fmt.Printf("before: %d\n", s4)
	s4, err := s.RemoveElement(s4, 2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("after: %d\n", s4)
}
