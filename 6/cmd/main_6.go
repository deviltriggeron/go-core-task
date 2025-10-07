package main

import (
	"fmt"
	g "generator"
	"time"
)

func main() {
	ch := make(chan int)

	go g.GenerateRandomInt(ch, 50)

	for c := range ch {
		fmt.Printf("Случайное число :%d\n", c)
		time.Sleep(time.Millisecond * 150)
	}
	fmt.Println("Генерация завершена")
}
