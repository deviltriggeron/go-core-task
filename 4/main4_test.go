package nce

import (
	"fmt"
	"testing"
)

func TestNotContainingElem1(t *testing.T) {
	slice1 := []string{"apple"}
	slice2 := []string{"banana", "date", "fig", "after", "slice", "map"}
	res := NotContainingElem(slice1, slice2)
	if res == nil {
		t.Error("No containing")
	}
	fmt.Println(res)
}

func TestNotContainingElem2(t *testing.T) {
	slice1 := []string{"banana", "date", "fig", "after", "slice", "map"}
	slice2 := []string{"apple"}
	res := NotContainingElem(slice1, slice2)
	if res == nil {
		t.Error("No containing")
	}
	fmt.Println(res)
}
