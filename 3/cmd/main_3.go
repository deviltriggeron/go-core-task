package main

import (
	"fmt"
	"mymap"
)

func main() {
	m := mymap.NewMyMap[string, float32]()
	m.Add("asd", 12)
	fmt.Println(m.Get("asda"))
}
