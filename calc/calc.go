package calc

import "fmt"

// a, b both int
func Multiply(a, b int) int {
	return a * b
}

func Add(numbers ...int) int {
	total := 0
	for index, number := range numbers {
		fmt.Println(index)
		total += number
	}

	for i := 0; i < len(numbers); i++ {
		// ..
	}

	return total
}
