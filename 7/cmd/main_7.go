package main

import (
	"fmt"
	merge "mergechan"
)

func main() {
	ch1 := merge.Produce(3)
	ch2 := merge.Produce(10)
	ch3 := merge.Produce(15)

	res := merge.MergeChan(ch1, ch2, ch3)

	for v := range res {
		fmt.Printf("%d, ", v)
	}
}
