package main

import "fmt"

func intSeq() func() int {
	acc := 0
	return func() int {
		acc++
		return acc
	}
}

func main() {
	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	newInts := intSeq()
	fmt.Println(newInts()) // 1
}
