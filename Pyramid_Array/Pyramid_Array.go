/*
Write a function that when given a number >= 0, returns an Array of ascending length subarrays.

pyramid(0) => [ ]
pyramid(1) => [ [1] ]
pyramid(2) => [ [1], [1, 1] ]
pyramid(3) => [ [1], [1, 1], [1, 1, 1] ]
Note: the subarrays should be filled with 1s
It("Testing for 0", func() [][]int{}
It("Testing for 1", func() [][]int{[]int{1}}
It("Testing for 2", func() [][]int{[]int{1}, []int{1, 1}}
It("Testing for 3", func()[][]int{[]int{1}, []int{1, 1}, []int{1, 1, 1}}
*/

package main

import "fmt"

func main() {
	// fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal

	// rowsLen := len(sample)
	// columnsLen := len(sample[0])

	for i := 0; i <= 6; i++ {
		value := 1

		// firstLine	:= []int{value}
		// secondLine	:= []int{1, 1}
		// thirdLine	:= []int{1,1,1}
		// fourthLine	:= []int{1,1,1,1}
		// fifthLine	:= []int{1,1,1,1,1}
		// sixthLine	:= []int{1,1,1,1,1,1}
		// seventhLine	:= []int{1,1,1,1,1,1,1}
		// eighthLine	:= []int{1,1,1,1,1,1,1,1}
		// ninthLine	:= []int{1,1,1,1,1,1,1,1,1}
		// tenthLine	:= []int{1,1,1,1,1,1,1,1,1,1}

		// temp := make([][]int, 0)
		// temp = append(temp, firstLine)
		// temp = append(temp, secondLine)

		orders := make([][]int, 0) // create my 2d slice with zeros
		for i := 0; i <= 6; i++ {
			// value := rand.Float64()
			temp := make([]int, 0)
			temp = append(temp, value)
			orders = append(orders, [][]int{temp}...)
		}

		fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal
		fmt.Println(orders)

		// for _, row := range orders {
		// 	fmt.Print(row, " \n")
		// }
	}

}
