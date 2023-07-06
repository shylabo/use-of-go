package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["key1"] = 7
	m["key2"] = 12

	fmt.Println("map:", m) // map: map[key1:7 key2:12]

	val1 := m["key1"]
	fmt.Println("val1:", val1) // val1: 7

	val3 := m["key3"]
	fmt.Println("val3:", val3) // val3: 0

	fmt.Println("len:", len(m)) // len: 2

	delete(m, "key2")
	fmt.Println("map:", m) // map: map[key1:7]

	// optional return value tells if the key exists. (if the value will not be used, use _ to ignore)
	_, ok := m["key2"]
	fmt.Println("present:", ok)

	// map can take any types of key and value
	n := map[interface{}]interface{}{"key1": 1, 2: "2", true: false}
	fmt.Println("map:", n)
	fmt.Printf("TYPE: %T, %T, %T\n", n["key1"], n[2], n[true])
	fmt.Printf("VALUE: %v, %v, %v\n", n["key1"], n[2], n[true])
}
