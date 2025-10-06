package variable

import (
	"crypto/sha256"
	"fmt"
	"reflect"
	"strconv"
)

func CreateVal() []interface{} {
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

func TypeOf(v any) reflect.Type {
	return reflect.TypeOf(v)
}

func Concatenation(varuables []interface{}) string {
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

func HashWithSalt(r []rune) [32]byte {
	salt := "go-2024"
	mid := len(r) / 2

	newRune := make([]rune, len(r)+len(salt))
	newRune = append(newRune, r[:mid]...)
	newRune = append(newRune, []rune(salt)...)
	newRune = append(newRune, r[mid:]...)

	sum := sha256.Sum256([]byte(string(newRune)))

	return sum
}
