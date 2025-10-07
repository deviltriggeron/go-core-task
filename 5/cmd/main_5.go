package main

import (
	c "contains"
	"fmt"
)

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	s, ok := c.Contains(a, b)
	if ok {
		fmt.Println(s)
	}
}
