package main

import "fmt"

func main() {
	i, j := 10, 20

	p := &i
	fmt.Println("This is the initial value of i", i)
	fmt.Println("This is the address of i", p)
	fmt.Println("This is the initial value of j", j)
	*p = 30
	fmt.Println("This is the new value of i", *p)

	p = &j

	fmt.Println("This is the address of j", p)
	*p = 40
	fmt.Println("This is the modified value  of j", j)

	*p = *p / 4
	fmt.Println("This is the calculated value of j", *p)

}
