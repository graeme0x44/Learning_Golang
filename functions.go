package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	re := plus(1, 2)
	fmt.Println(re)

	res := plusPlus(1, 2, 3)
	fmt.Println(res)
}
