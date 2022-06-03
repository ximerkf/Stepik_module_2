package main

import (
	"fmt"
	"math"
)

var p, v, k float64 = 1, 2, 3

func main() {
	fmt.Println(T())
}

func T() float64 {
	return 6 / W()
}

func W() float64 {
	return math.Sqrt(k / M())
}

func M() float64 {
	return p * v
}
