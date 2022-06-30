/*
Complete the square sum function so that it squares each number passed into it and then sums the results together.
For example, for [1, 2, 2] it should return 9 because 1^2 + 2^2 + 2^2 = 9.
    It("Testing [1,2]",func() {Expect(SquareSum([]int{1,2})).To(Equal(5))})
    It("Testing [0,3,4,5]",func() {Expect(SquareSum([]int{0,3,4,5})).To(Equal(50))})
    It("Testing []",func() {Expect(SquareSum([]int{})).To(Equal(0))})
*/
	// fmt.Print("\033[2J") //clear the screen befor 
package main

import(
	"fmt"
	"os"
	"strconv"
)

/*func function(int) int {
	return 0
}*/

func main(){
	for _,arg := range os.Args[1:] {
			n, err := strconv.Atoi(arg)
			if err != nil {
					fmt.Fprintf(os.Stderr, "Error: %v\n", err)
					os.Exit(1)
			}
			fmt.Printf("%d\n", n)
	}
}