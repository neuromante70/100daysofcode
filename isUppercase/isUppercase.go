// 65-90 capital letters
// 97-122 lowercase letters
package main

import "fmt"

var MyString string

func main() {
	MyString = "AbCD"
	myRune := []rune(MyString)
	isLowercase := false
	for _, num := range myRune {
		if int(num) < 65 || int(num) > 90 {
			// fmt.Println("not capital letter", int(num))
			isLowercase = true
			break
		}
	}
	if isLowercase {
		fmt.Println("not all capital letter")
	} else {
		fmt.Println("ALL CAPITAL letter")
	}
}
