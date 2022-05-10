package main

import (
	"fmt"
	"reflect"
)

func main() {
	testBoolVar := 0
	result := reflect.ValueOf(testBoolVar).Kind().String()

	if result == "bool" {
		fmt.Println("You intered not a bool value")
	} else {
		fmt.Println("Your bool value is: ", result)
	}

}