package main

import (
	"fmt"
	"strconv"
)

// func isNumeric(s string) bool {
// 	_, err := strconv.Atoi(s)
// 	if err != nil {
// 		return false
// 	} else {
// 		return true
// 	}
// }

func main() {
	var s string

	// var i int
	fmt.Println("Enter a integer number: ")
	fmt.Scan(&s)

	for {
		getint, err := strconv.Atoi(s)

		if err != nil || getint < 0 {
			fmt.Println("Not an integer number. Please insert a valid input:")
			fmt.Scan(&s)
		} else {
			if getint > 0 {
				fmt.Println("Well done! Job finished")
				break
			}

		}
	}
	fmt.Println("Uscito dal for, prog terminato")
}
