package main

import (
	"fmt"
	"math"
)

func howBigIsNumber(n int) int {
	var i int
	for i = 0; n > 10; i++ {
		n /= 10
	}
	return i
}

func howBigIsNumber2(n int) int {
	s := fmt.Sprintf("%d", n)
	return len(s) - 1
}

func main() {
	fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal
	myNumber := 12401
	magnitude := howBigIsNumber(myNumber)
	powerOf := math.Pow(10, float64(magnitude))
	fmt.Println("Your number is ", magnitude, "x 10 big and approximately", powerOf)
	fmt.Println("Second version: ", howBigIsNumber2(myNumber))
}
