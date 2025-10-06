package mymap

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	m := NewMyMap[string, int]()
	fmt.Println(reflect.TypeOf(m))
	mm := NewMyMap[float32, bool]()
	fmt.Println(reflect.TypeOf(mm))
}

func TestAddElem(t *testing.T) {
	m := NewMyMap[string, int]()
	m.Add("key", 12)
	res, ok := m.Get("key")
	if !ok {
		fmt.Println("key does not exist")
	}
	fmt.Println(res)
}

func TestRemove(t *testing.T) {
	m := NewMyMap[string, int]()
	m.Add("key", 12)

	if m.Exists("key") {
		fmt.Println("key exist")
	} else {
		fmt.Println("key does not exist")
	}

	m.Remove("key")

	if m.Exists("key") {
		fmt.Println("key exist")
	} else {
		fmt.Println("key does not exist")
	}
}

func TestCopy(t *testing.T) {
	m := NewMyMap[string, int]()
	m.Add("key", 12)

	copy := m.Copy()

	copy.Add("key", 2)

	res, ok := m.Get("key")
	if !ok {
		fmt.Println("key does not exist")
	}
	fmt.Printf("Original map: %d\n", res)

	copied, ok := copy.Get("key")
	if !ok {
		fmt.Println("key does not exist")
	}
	fmt.Printf("Copy changed map: %d\n", copied)
}
