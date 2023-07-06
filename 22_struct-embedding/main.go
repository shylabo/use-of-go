package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}

func main() {
	container := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("container={num: %v, str: %v}\n", container.num, container.str)

	fmt.Println("also num:", container.base.num) // Alternatively, we can spell out the full path using the embedded type name.

	fmt.Println("describe:", container.describe())

	type describer interface {
		describe() string
	}

	var d describer = container
	fmt.Println("describer:", d.describe())
}
