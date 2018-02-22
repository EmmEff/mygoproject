package main

import "fmt"

// Subtract subtracts y from x
func Subtract(x int, y int) int {
	return x - y
}

// Sum adds x and y
func Sum(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello")

	fmt.Println(Sum(1, 2))
}
