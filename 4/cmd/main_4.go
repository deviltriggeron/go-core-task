package main

import (
	"fmt"
	"nce"
)

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	resSlice := nce.NotContainingElem(slice1, slice2)
	fmt.Println(resSlice)
}
