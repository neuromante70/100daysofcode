package main

import "fmt"

func factorial_calc(number int) int {
	if number == 0 || number == 1 { //base case
		return 1
	}

	if number < 0 { //if number is negative
		fmt.Println("Invalid Number")
		return -1
	}

	return number * factorial_calc(number-1)
	// recursive function until operation reaches the base case
}

func main() {
	answer1 := factorial_calc(0)
	//where 0 is the parameter to be passed
	fmt.Println(answer1)

	answer2 := factorial_calc(4)
	fmt.Println(answer2)

	answer3 := factorial_calc(-1)
	fmt.Println(answer3)

	answer4 := factorial_calc(5)
	fmt.Println(answer4)

}
