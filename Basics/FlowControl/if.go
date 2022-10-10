package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x)) //Sprint is used to generate and return a formatted string
}
func main() {
	fmt.Println(sqrt(4), sqrt(-9))
}
