package main

import "fmt"

func zeroVal(iVal int) {
	iVal = 0
}

func zeroPtr(iPtr *int) {
	*iPtr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i) // 1

	zeroVal(i)
	fmt.Println("zeroVal:", i) // 1

	zeroPtr(&i)
	fmt.Println("zeroPtr:", i) // 0

	fmt.Println("pointer:", &i) // pointer address
}
