/*
Grasshopper - Check for factor
This function should test if the factor is a factor of base.
Return true if it is a factor or false if it is not.

About factors
Factors are numbers you can multiply together to get another number.

2 and 3 are factors of 6 because: 2 * 3 = 6

You can find a factor by dividing numbers. If the remainder is 0 then the number is a factor.
You can use the mod operator (%) in most languages to check for a remainder
For example 2 is not a factor of 7 because: 7 % 2 = 1

Note: base is a non-negative number, factor is a positive number.
*/

package main

import (
	"fmt"
	"strconv"
)

func isNumeric(s string) int {
	var r int
out:
	for {
		getInt, err := strconv.Atoi(s)
		if err != nil || getInt < 0 {
			fmt.Println("Not a (positive) integer number. Please insert a valid input:")
			fmt.Scan(&s)
		} else {
			if err == nil && getInt > 0 {
				r, _ = strconv.Atoi(s)
				break out
			}
		}
	}
	return r
}

func main() {
	var s string

	fmt.Println("Enter a integer number (base): ")
	fmt.Scan(&s)
	var base = isNumeric(s)
	fmt.Println("Enter a integer number (factor): ")
	fmt.Scan(&s)
	var factor = isNumeric(s)

	if factor%base == 0 {
		fmt.Println(factor, "is a factor of", base)
	} else {
		fmt.Println(factor, "is NOT a factor of", base)
	}
}
