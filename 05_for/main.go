package main

import "fmt"

func main() {
	i := 0
	for i < 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Variable in for statement has different scope
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println(i) // 3

	for {
		fmt.Println("loop")
		break
	}

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i) // Prints only odd number
	}
}
