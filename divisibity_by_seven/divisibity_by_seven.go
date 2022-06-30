/* A number m of the form 10x + y is divisible by 7 if and only if x − 2y is divisible by 7. In other words, subtract twice the last digit from the number formed by the remaining digits. Continue to do this until a number known to be divisible by 7 is obtained; you can stop when this number has at most 2 digits because you are supposed to know if a number of at most 2 digits is divisible by 7 or not.

The original number is divisible by 7 if and only if the last number obtained using this procedure is divisible by 7.

Examples:
1) m = 371 -> 37 − (2×1) -> 37 − 2 = 35 ; thus, since 35 is divisible by 7, 371 is divisible by 7.

The number of steps to get the result is 1.

2) m = 1603 -> 160 - (2 x 3) -> 154 -> 15 - 8 = 7 and 7 is divisible by 7.

3) m = 372 -> 37 − (2×2) -> 37 − 4 = 33 ; thus, since 33 is not divisible by 7, 372 is not divisible by 7.

4) m = 477557101->47755708->4775554->477547->47740->4774->469->28 and 28 is divisible by 7, so is 477557101. The number of steps is 7.

Task:
Your task is to return to the function seven(m) (m integer >= 0) an array (or a pair, depending on the language) of numbers, the first being the last number m with at most 2 digits obtained by your function (this last m will be divisible or not by 7), the second one being the number of steps to get the result.

Forth Note:
Return on the stack number-of-steps, last-number-m-with-at-most-2-digits

Examples:
seven(371) should return [35, 1]
seven(1603) should return [7, 2]
seven(477557101) should return [28, 7]
seven(477557102) should return [47, 7]
*/

package main

import "fmt"

/*func reverseInt(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func splitToDigits(n int) []int {
	var numSplit []int

	for n != 0 {
		numSplit = append(numSplit, n%10)
		n /= 10
	}
	reverseInt(numSplit)
	return numSplit
}*/

func howLongIsNumber(n int64) int64 {
	s := fmt.Sprintf("%d", n)
	return int64(len(s))
}

func Seven(n int64) (result []int) {
	// start number: 371
	cycleNr := howLongIsNumber(n)
	result = make([]int, 2)
	var i int64
	var o int64
	for i = 1; n >= cycleNr; i++ {
		// m = n % 10                               // result 1, remainder of the division by 10; 371 / 10 = 37, mod 1
		// o = (n - m) / 10                         // result 37, orNum - units divided by 10; (371 - 1) / 10 = 37
		// p = o - m*2                              // result 35; 37 - 1 * 2 = 35
		o = (n-n%10)/10 - (n % 10 * 2)
		if o%7 == 0 && howLongIsNumber(o) <= 2 { // 35 / 7 = 5, mod 0
			result[0] = int(o)
			result[1] = int(i)
			return result
			break
		} else if howLongIsNumber(o) > 2 {
			n = o
		} else {
			result[0] = int(o)
			result[1] = int(i)
			return result
			break
		}
	}
	return result
}

func main() {
	givenNumber := int64(477557102)
	// myResult := []int{}
	myResult := Seven(givenNumber)
	fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal
	fmt.Println("Your number is", myResult[0], "divisible for 7 and it takes", myResult[1], "passages")

}
