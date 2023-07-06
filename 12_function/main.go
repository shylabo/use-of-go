package main

import "fmt"

func sumTwoNums(a int, b int) int {
	return a + b
}

func sumThreeNums(a, b, c int) int {
	return a + b + c
}

func main() {
	result := sumTwoNums(1, 2)
	fmt.Println("1 + 2 =", result)

	result = sumThreeNums(1, 2, 3)
	fmt.Println("1 + 2 + 3 =", result)
}
