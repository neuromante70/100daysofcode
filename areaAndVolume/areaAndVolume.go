package main

import "fmt"

func GetSize(w, h, d int) [2]int {
	// Surface = 2wh + 2wd + 2hd
	// Area = w * h * d
	return [2]int{(2 * w * h) + (2 * w * d) + (2 * h * d), w * h * d}
}

func main() {
	var areaAndSurface [2]int = GetSize(2, 2, 2)
	fmt.Println(areaAndSurface[0], areaAndSurface[1])
}
