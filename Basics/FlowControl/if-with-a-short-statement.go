package main

import (
	"fmt"
	"math"
)

func pow(a, b, lim float64) float64 {
	if v := math.Pow(a, b); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(pow(2, 4, 6))
	fmt.Println(pow(2, 2, 10))
}
