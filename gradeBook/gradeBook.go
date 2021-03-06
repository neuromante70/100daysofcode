/*
Grade book
Complete the function so that it finds the average of the three scores passed to it
and returns the letter value associated with that grade.
Numerical Score	Letter Grade
90 <= score <= 100	'A'
80 <= score < 90	'B'
70 <= score < 80	'C'
60 <= score < 70	'D'
0 <= score < 60	'F'
Tested values are all between 0 and 100. Theres is no need to check for negative values
or values greater than 100.
*/

package main

import "fmt"

func GetGrade(a, b, c int) rune {
	// var grade rune

	// average := (a + b + c) / 3
	// 90 <= score <= 100 'A'
	switch average := (a + b + c) / 3; {
	case average >= 90:
		return 'A'
	// 80 <= score < 90	'B'
	case average >= 80:
		return 'B'
	// 70 <= score < 80	'C'
	case average >= 70:
		return 'C'
	// 60 <= score < 70	'D'
	case average >= 60:
		return 'D'
	// 0 <= score < 60	'F'
	default:
		return 'F'
	}
}

func main() {
	// average := (a + b + c) / 3
	result := GetGrade(40, 40, 40)
	fmt.Println("")
	fmt.Println("Your grade is ", string(result))
}
