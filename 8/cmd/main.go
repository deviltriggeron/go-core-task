package main

import (
	"fmt"
	"mywg"
)

func main() {
	wg := mywg.NewCustomWg()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i <= 3; i++ {
			fmt.Println(i)
		}
	}()

	wg.Wait()
	fmt.Println("programm end")
}
