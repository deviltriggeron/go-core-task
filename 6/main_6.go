package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomInt(ch chan int, count int) {
	for i := 0; i < count; i++ {
		r := rand.Intn(100)
		ch <- r
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

func main() {
	ch := make(chan int)

	go generateRandomInt(ch, 50)

	for c := range ch {
		fmt.Printf("Случайное число :%d\n", c)
		time.Sleep(time.Millisecond * 150)
	}
	fmt.Println("Генерация завершена")
}
