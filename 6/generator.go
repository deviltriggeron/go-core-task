package generator

import (
	"math/rand"
	"time"
)

func GenerateRandomInt(ch chan int, count int) {
	for i := 0; i < count; i++ {
		r := rand.Intn(100)
		ch <- r
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}
