/*
Is the string uppercase?
Task
Create a method IsUpperCase to see whether the string is ALL CAPS. For example:

type MyString string
MyString("c").IsUpperCase() == false
MyString("C").IsUpperCase() == true
MyString("hello I AM DONALD").IsUpperCase() == false
MyString("HELLO I AM DONALD").IsUpperCase() == true
MyString("ACSKLDFJSgSKLDFJSKLDFJ").IsUpperCase() == false
MyString("ACSKLDFJSGSKLDFJSKLDFJ").IsUpperCase() == true

In this Kata, a string is said to be in ALL CAPS whenever it does not contain any lowercase letter so any string containing no letters at all is trivially considered to be in ALL CAPS.
*/

package main

import (
	"fmt"
	//	"unicode"
)

/* type MyString string

func (s MyString) IsUpperCase() bool {
  // Your code here!

return true
} */

func main() {
	// uppercase chars: 065 - 090
	// lowercase chars: 097 - 122

	var myString = "ABCD"
	myRune := []rune(myString)
	//lenR := len(myRune) - 1
	itsUppercase := true
	for j := 0; j <= len(myRune) - 1; j++ {
		if myRune[j] < 65 || myRune[j] > 90 {
			// if unicode.IsUpper(myRune[j]) {
			itsUppercase = false
		}
	}
	fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal
	fmt.Println(myRune)
	fmt.Println("The string", myString, "is uppercase?", itsUppercase)
}
