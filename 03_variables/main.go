package main

import "fmt"

func main() {
	var a = "coffee"
	fmt.Printf("Type: %T Value: %v\n", a, a) // Type: string Value: coffee

	var b, c int = 1, 2
	fmt.Printf("Type: %T, %T Value: %v, %v\n", b, c, b, c) // Type: int, int Value: 1, 2

	var d = true
	fmt.Printf("Type: %T Value: %v\n", d, d) // Type: bool Value: true

	var e int
	fmt.Printf("Type: %T Value: %v\n", e, e) // Type: int Value: 0 *initialized as zero-value

	var f string
	fmt.Printf("Type: %T Value: %v\n", f, f) // Type: string Value: *(empty string) initialized as zero-value

	var g bool
	fmt.Printf("Type: %T Value: %v\n", g, g) // Type: bool Value: false *initialized as zero-value

	var h float32
	fmt.Printf("Type: %T Value: %v\n", h, h) // Type: float32 Value: 0 *initialized as zero-value

	var i float64
	fmt.Printf("Type: %T Value: %v\n", i, i) // Type: float64 Value: 0 *initialized as zero-value

	j := "cookie"
	fmt.Printf("Type: %T Value: %v\n", j, j) // Type: string Value: cookie
}
