package main

import "fmt"

func main() {
	sum := 10
	for sum < 100 {
		sum += sum
	}
	fmt.Println(sum)
}
