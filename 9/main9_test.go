package conveyor

import (
	"fmt"
	"testing"
)

func TestConveyor(t *testing.T) {
	in := make(chan uint8)

	go Conveyor(in, 10)

	for ch := range in {
		fmt.Printf("%d, ", ch)
	}
}

func TestCubing(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)

	go func() {
		for i := 0; i < 10; i++ {
			in <- uint8(i)
		}
		close(in)
	}()

	go Cubing(in, out)
	for ch := range out {
		fmt.Printf("%.2f, ", ch)
	}
}
