package mergechan

import (
	"fmt"
	"testing"
)

func TestMergeChan(t *testing.T) {
	ch1 := Produce(40)
	ch2 := Produce(-3)
	ch3 := Produce(12)
	ch4 := Produce(1)
	ch5 := Produce(123)

	res := MergeChan(ch1, ch2, ch3, ch4, ch5)

	for v := range res {
		fmt.Printf("%d, ", v)
	}
}
