package mywg

import (
	"fmt"
	"testing"
)

func TestWg(t *testing.T) {
	wg := NewCustomWg()
	ch := make(chan int)

	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 1; i <= 15; i++ {
			ch <- i
		}
		close(ch)
	}()

	for c := range ch {
		fmt.Println(c)
	}

	wg.Wait()
	fmt.Println("programm end")
}

func TestWg2(t *testing.T) {
	wg := NewCustomWg()
	ch := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 1; i <= 15; i++ {
			ch <- i
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 15; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	go func() {
		defer wg.Done()
		for i := 1; i <= 15; i++ {
			ch3 <- i
		}
		close(ch3)
	}()

	for c := range ch {
		fmt.Println(c)
	}

	for c := range ch2 {
		fmt.Println(c)
	}

	for c := range ch3 {
		fmt.Println(c)
	}

	wg.Wait()
	fmt.Println("programm end")
}
