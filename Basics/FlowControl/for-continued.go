package main

import "fmt"

func main() {
	sum := 1
	// for ; sum < 10 ; { semi-colon can be ignored
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)
}
