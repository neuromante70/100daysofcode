/*
Consider an array/list of sheep where some sheep may be missing from their place.
We need a function that counts the number of sheep present in the array (true means present).

For example,
[true,  true,  true,  false,
 true,  true,  true,  true ,
 true,  false, true,  false,
 true,  false, false, true ,
 true,  true,  true,  true ,
 false, false, true,  true]
The correct answer would be 17.

Hint: Don't forget to check for bad values like null/undefined
*/
package main

import (
	"fmt"
	"reflect"
)

// func CountSheeps(numbers []bool) int {
// 	return 0 // your code here
//   }

func countSheeps(numbers []bool) int {
	total := 0
	for noSheep, count := range numbers {
		if reflect.ValueOf(count).Kind().String() == "bool" {
			if count {
				total++
			}
		} else if count {
			return noSheep
		}
	}
	return total
}

func main() {
	// fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal
	sheeps := []bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	}
	result := countSheeps(sheeps)
	if result == 0 {
		fmt.Println("You intered not a bool value")
	} else {
		fmt.Println("Total sheep are: ", result)
	}
	// fmt.Println(sheeps)
}
