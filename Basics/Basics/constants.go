package main

import "fmt"

const pi = 3.14

func main() {
	const world = "India"
	fmt.Println("hello", world)
	fmt.Println("happy", pi, "day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
