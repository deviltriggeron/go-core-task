package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func createVal() []interface{} {
	var res []interface{}

	var numDecimal int = 42           // Десятичная система
	var numOctal int = 052            // Восьмеричная система
	var numHexadecimal int = 0x2A     // Шестнадцатиричная система
	var pi float64 = 3.14             // Тип float64
	var name string = "Golang"        // Тип string
	var isActive bool = true          // Тип bool
	var complexNum complex64 = 1 + 2i // Тип complex64

	res = append(res, numDecimal)
	res = append(res, numOctal)
	res = append(res, numHexadecimal)
	res = append(res, pi)
	res = append(res, name)
	res = append(res, isActive)
	res = append(res, complexNum)

	return res
}

func typeOf(v any) reflect.Type {
	return reflect.TypeOf(v)
}

func concatenation(varuables []interface{}) string {
	var res string
	for _, v := range varuables {
		switch val := v.(type) {
		case int:
			res += strconv.Itoa(val)
		case float64:
			res += strconv.FormatFloat(val, 'f', -1, 64)
		case string:
			res += val
		case bool:
			res += strconv.FormatBool(val)
		case complex64:
			res += fmt.Sprintf("(%g+%gi)", real(val), imag(val))
		}
	}
	return res
}

func hashWithSalt() {

}

func main() {
	varuables := createVal()

	fmt.Println("Varuables:")
	for _, v := range varuables {
		fmt.Println(v)
	}

	fmt.Println("Varuables type:")
	for _, v := range varuables {
		fmt.Println(typeOf(v))
	}

	fmt.Println("Varuables typecast to string and concatenation:")
	str := concatenation(varuables)
	fmt.Println(str)

	fmt.Println("string typecast to rune:")
	runes := []rune(str)
	fmt.Printf("Count rune:\n%d\n", len(runes))
}

/*
1. Создает несколько переменных различных типов данных:
```
int (три числа в десятичной, восьмеричной и шеснадцатиричной системах)
float64
string
bool
complex64
```
2. Определяет тип каждой переменной и выводит его на экран.
3. Преобразует все переменные в строковый тип и объединяет их в одну строку.
4. Преобразовать эту строку в срез рун.
5. Захэшировать этот срез рун SHA256, добавив в середину соль "go-2024" и вывести результат.

* Напишите unit тесты к созданным функциям

Напишите main функцию, в которой протестируете весь вышеописанный функционал. Выведите результаты на экран.

Входные числа из пункта 1 могут быть:
```
var numDecimal int = 42           // Десятичная система
var numOctal int = 052            // Восьмеричная система
var numHexadecimal int = 0x2A     // Шестнадцатиричная система
var pi float64 = 3.14             // Тип float64
var name string = "Golang"         // Тип string
var isActive bool = true           // Тип bool
var complexNum complex64 = 1 + 2i  // Тип complex64
```

---
*/
