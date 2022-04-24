package main

import (
	"fmt"
	"strconv"
)

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	if err == nil {
		return true
	} else {
		return false
	}
}

func main() {
	var s string

	// var i int
	fmt.Println("Enter a integer number: ")
	fmt.Scan(&s)

	for {
		if isNumeric(s) {
			fmt.Println("Well done! Job finished")
			break
		} else {
			fmt.Println("Enter a integer number: ")
			fmt.Scan(&s)
		}
	}
	fmt.Println("uscito dal for, prog terminato")

	// for {
	// 	_, err := fmt.Scan(&s)
	// 	i, err = strconv.Atoi(s)
	// 	if err != nil {
	// 		fmt.Println("Enter a valid number")
	// 	} else {
	// 		fmt.Println("Got: " + strconv.Itoa(i))
	// 		break
	// 	}
	// }
	// //Todo

}
