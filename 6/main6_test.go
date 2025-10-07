package generator

import (
	"fmt"
	"testing"
	"time"
)

func TestGenerate1(t *testing.T) {
	ch := make(chan int)

	go GenerateRandomInt(ch, 5)

	for c := range ch {
		fmt.Printf("Случайное число :%d\n", c)
		time.Sleep(time.Millisecond * 150)
	}
	fmt.Println("Генерация завершена")
}

func TestGenerate2(t *testing.T) {
	ch := make(chan int)

	go GenerateRandomInt(ch, 100)

	for c := range ch {
		fmt.Printf("Случайное число :%d\n", c)
		time.Sleep(time.Millisecond * 150)
	}
	fmt.Println("Генерация завершена")
}
