package variable

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestType(t *testing.T) {
	m := make(map[string]int)
	tm := TypeOf(m)
	fmt.Println(tm)
}

func TestConcatenation(t *testing.T) {
	var variable []interface{}

	variable = append(variable, 12)
	variable = append(variable, "str")
	variable = append(variable, 12.12)
	res := Concatenation(variable)

	if res == "12str12.12" {
		fmt.Println(res)
	} else {
		t.Errorf("Test failed")
	}
}

func TestHash(t *testing.T) {
	str := "firstsecon"
	runes := []rune(str)
	res := HashWithSalt(runes)
	if res == sha256.Sum256([]byte("firstgo-2024secon")) {
		fmt.Println(res)
	} else {
		t.Errorf("Test failed")
	}
}
