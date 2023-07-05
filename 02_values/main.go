package main

import "fmt"

func main() {
	fmt.Println("Hello" + " " + "World") // Hello World
	fmt.Println("1 + 1 = ", 1+1)         // 1 + 1 =  2
	fmt.Println("8.0 / 3.0 = ", 8.0/3.0) // 8.0 / 3.0 =  2.6666666666666665

	fmt.Println(true && false) // false
	fmt.Println(true || false) // true
	fmt.Println(!false)        // true
}
