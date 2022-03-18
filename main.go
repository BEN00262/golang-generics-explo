package main

import (
	"fmt"
)

// https://go.dev/doc/tutorial/generics (reference)

// Example 1
// func sumNumbers[T int | float64](first T, second T) T {
// 	return first + second
// }

// Example 2
// type Number interface {
// 	int | float64
// }

// func sumNumbers[T Number](first T, second T) T {
// 	return first + second
// }

// Example 3
type Number interface {
	~int | ~float64
}

func sumNumbers[T Number](first T, second T) T {
	return first + second
}

type Numberish float64

func main() {
	var numberish Numberish = 67.8

	fmt.Println(
		sumNumbers(numberish, 8.9),
	)
}
