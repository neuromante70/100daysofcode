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

// func function(int) int {
// 	return 0
// }

func main() {
	// fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal

	sample := [6][5]int{}

	rowsLen := len(sample)
	// columnsLen := len(sample[0])

	// sample[1][0] = 1

	// sample[2][0] = 1
	// sample[2][1] = 1

	// sample[3][0] = 1
	// sample[3][1] = 1
	// sample[3][2] = 1

	//fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal

	for row := 1; row < rowsLen; row++ {
		for column := 0; column < row; column++ {
			sample[row][column] = 1
		}
	}

	// print array as a grid with for loop

	// for row := 0; row < 4; row++ {
	// 	for column := 0; column < 3; column++ {
	// 		fmt.Print(sample[row][column], " ")
	// 	}
	// 	fmt.Print("\n")
	// }

	//print array as a grid with for-range loop

	// for _, row := range sample {
	// 	for _, val := range row {
	// 		fmt.Print(val, " ")
	// 	}
	// 	fmt.Print("\n")
	// }

	//print array as a grid with for-range loop, a simplified version

	for _, row := range sample {
		fmt.Print(row, " \n")
	}

	// fmt.Println(rows)
	// fmt.Println(columns)

}
