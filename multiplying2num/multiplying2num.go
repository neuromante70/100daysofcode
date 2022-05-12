package main

import "fmt"

func mul2num(nums []int) int {
	return nums[0] * nums[1]
}

func main() {
	var nums = []int{5, 3}
	result := mul2num(nums)
	fmt.Println("Your result is: ", result)
}
