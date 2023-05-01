package main

import "fmt"

func avg(numbers ...int) int {
	count := 0
	total := 1
	for i, value := range numbers {
		total += value
		count = i
	}
	count++
	return total / count
}

func main() {
	fmt.Println(avg(1, 2, 3, 4, 5))
	fmt.Println(avg(2, 4, 6, 8, 10, 12, 14, 16, 18, 20))
}
