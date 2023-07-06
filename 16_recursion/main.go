package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	fmt.Println("factorial:", factorial(5))
	fmt.Println("fibonacci:", fibonacci(7))
}
