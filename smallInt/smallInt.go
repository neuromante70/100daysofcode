/*
Given an array of integers your solution should find the smallest integer.

For example:
Given [34, 15, 88, 2] your solution will return 2
Given [34, -345, -1, 100] your solution will return -345
You can assume, for the purpose of this kata, that the supplied array will not be empty.
*/

package main

import (
	"fmt"
	//"sort"
)

func SmallestIntegerFinder(numbers []int) int {
	len := len(numbers)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if numbers[j] < numbers[minIndex] {
				numbers[j], numbers[minIndex] = numbers[minIndex], numbers[j]
			}
		}
	}
	return numbers[0] // your code here
}

// func SmallestIntegerFinder(numbers []int) int {
// 	sort.Ints(numbers)
// 	return numbers[0] // your code here
// }

func main() {
	numbers := []int{-144, -68, 11, 41, 143, -116, -133, 84, 23, -53, 17, 143, -20, 12}
	solution := SmallestIntegerFinder(numbers)
	fmt.Println("\nThe smallest number is: ", solution)
}
