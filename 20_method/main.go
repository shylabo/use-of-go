package main

import "fmt"

type Rectangle struct {
	width, height int
}

func (r *Rectangle) area() int {
	return r.width * r.height
}

func (r Rectangle) perimeter() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := Rectangle{width: 10, height: 5}

	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perimeter())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perimeter:", rp.perimeter())
}
