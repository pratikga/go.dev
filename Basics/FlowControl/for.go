package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 8; i++ {
		sum = sum + i
	}
	fmt.Println(sum)
}
