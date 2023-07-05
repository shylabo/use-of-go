package main

import "fmt"

func main() {
	var s []string
	fmt.Println("uninit: ", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len: ", len(s), "cap: ", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("After addition: ", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copied slice: ", c)

	l := s[2:5] // first index <= x < last index
	fmt.Println("slice l from s: ", l)
	fmt.Println("Original slice s isn't affected: ", s)

	l = s[:5]
	fmt.Println("updated l", l)

	l = s[2:]
	fmt.Println("updated l", l)

	t := []string{"g", "h", "i"}
	fmt.Println("A slice with initial values", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
