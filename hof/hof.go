package main

import (
	"fmt"
)

func sum(a, b int) int {
	return a + b
}

func partialSum(a int) func(int) int {
	return func(b int) int {
		return sum(a, b)
	}
}

func main() {
	partial := partialSum(2)
	fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal
	fmt.Println(partial(5))
}
