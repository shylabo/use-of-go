package main

import "fmt"

func sum(nums ...int) int {
	fmt.Println(nums) // [1 2 3 4]

	total := 0
	for _, num := range nums {
		total += num
	}

	return total
}

func main() {
	result := sum(1, 2, 3, 4)
	fmt.Println(result) // 10
}
