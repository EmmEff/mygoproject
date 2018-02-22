package main

import "fmt"

func Subtract(x int, y int) int {
	return x - y
}

func Sum(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello")

	fmt.Println(Sum(1, 2))
}
