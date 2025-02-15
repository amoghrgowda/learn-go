package main

import "math"

func conv() (z uint) {
	var x, y int = 10, 20
	var f float64 = math.Sqrt(float64(x*x + y*y))
	z = uint(f)
	return
}
