package main

import (
	"fmt"
	"math"
)

func gip(a, b float64) float64 {
	return a*a + b*b
}

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	fmt.Println(math.Sqrt(gip(a, b)))
}
