package conveyor

func Cubing(ch1 chan uint8, ch2 chan float64) {
	for in := range ch1 {
		f := float64(in)
		ch2 <- f * f * f
	}
	close(ch2)
}

func Conveyor(ch chan uint8, count int) {
	for i := 0; i <= count; i++ {
		ch <- uint8(i)
	}
	close(ch)
}
