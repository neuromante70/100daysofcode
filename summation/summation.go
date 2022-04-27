/*
Summation
Write a program that finds the summation of every n from 1 to num.
The n will always be a positive integer greater than 0.
For example:
summation(2) -> 3
1 + 2

summation(8) -> 36
1 + 2 + 3 + 4 + 5 + 6 + 7 + 8
*/

package main

import "fmt"

func Summation(n int) int {
	if n == 1 { //base case
		return 1
	}

	if n <= 0 { //if n is 0 or negative
		// Invalid number
		return 0
	}

	return n + Summation(n-1)
	// recursive function until operation reaches the base case
}

func main() {
	answer1 := Summation(0)
	//where 0 is the parameter to be passed
	fmt.Println(answer1)

	answer2 := Summation(4)
	fmt.Println(answer2)

	answer3 := Summation(-1)
	fmt.Println(answer3)

	answer4 := Summation(5)
	fmt.Println(answer4)
}
