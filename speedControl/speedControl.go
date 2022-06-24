/*
Speed Control
In John's car the GPS records every s seconds the distance travelled from an origin (distances are measured in an arbitrary but consistent unit). For example, below is part of a record with s = 15

The sections are:
0.0-0.19, 0.19-0.5, 0.5-0.75, 0.75-1.0, 1.0-1.25, 1.25-1.50, 1.5-1.75, 1.75-2.0, 2.0-2.25
and  the Delta is:
x = [0.0, 0.19, 0.5, 0.75, 1.0, 1.25, 1.5, 1.75, 2.0, 2.25]


To calculate hourly speed you can use:
(3600 * delta_distance) / s

We can calculate John's average hourly speed on every section and we get:
[0.0,  0.19, 0.5,  0.75, 1.0,  1.25, 1.5,  1.75, 2.0, 2.25]
[45.6, 74.4, 60.0, 60.0, 60.0, 60.0, 60.0, 60.0, 60.0w]

Given s and x the task is to return as an integer the *floor* of the maximum average speed per hour obtained on the sections of x. If x length is less than or equal to 1 return 0 since the car didn't move.

Example:
with the above data your function gps(s, x)should return 74

Note
With floats it can happen that results depends on the operations order.
*/

package main

import "log"

func Gps(s int, segments []float64) int {
	var maxSpeed float64

	// lenghtSegment := len(segments)
	for x := 0; x < len(segments)-1; x++ {
		delta := segments[x+1] - segments[x]
		hourly_speed := 3600.00 * delta / float64(s)
		if hourly_speed > maxSpeed {
			maxSpeed = hourly_speed
		}
	}
	return int(maxSpeed)
}

func main() {
	// fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal
	// var gpsCoordinates = []float64{0, 0.19, 0.5, 0.75, 1, 1.25, 1.5, 1.75, 2, 2.25}
	var gpsCoordinates = []float64{0.0, 0.11, 0.22, 0.33, 0.44, 0.65, 1.08, 1.26, 1.68, 1.89, 2.1, 2.31, 2.52, 3.25}
	second := 12

	result := Gps(second, gpsCoordinates)
	log.Println("The result is: ", result)
}
