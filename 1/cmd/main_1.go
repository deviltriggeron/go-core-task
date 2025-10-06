package main

import (
	"fmt"
	variable "variable"
)

func main() {
	varuables := variable.CreateVal()

	fmt.Println("Varuables:")
	for _, v := range varuables {
		fmt.Println(v)
	}

	fmt.Println("Varuables type:")
	for _, v := range varuables {
		fmt.Println(variable.TypeOf(v))
	}

	fmt.Println("Varuables typecast to string and concatenation:")
	str := variable.Concatenation(varuables)
	fmt.Println(str)

	fmt.Println("string typecast to rune:")
	runes := []rune(str)
	fmt.Printf("Count rune:\n%d\n", len(runes))

	res := variable.HashWithSalt(runes)
	fmt.Println(len(res))
	fmt.Println(res)
}
