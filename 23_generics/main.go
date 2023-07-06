package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type element[T any] struct {
	next *element[T]
	val  T
}

type List[T any] struct {
	head *element[T]
	tail *element[T]
}

func (list *List[T]) Push(v T) {
	if list.tail == nil {
		list.head = &element[T]{val: v}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: v}
		list.tail = list.tail.next
	}
}

func (list *List[T]) GetAll() []T {
	var elems []T
	for elem := list.head; elem != nil; elem = elem.next {
		elems = append(elems, elem.val)
	}
	return elems
}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys:", MapKeys(m))

	// we don’t have to specify the types for K and V when calling MapKeys but technically we can.
	_ = MapKeys[int, string](m)

	list := List[int]{}
	list.Push(1)
	list.Push(2)
	list.Push(3)
	fmt.Println("list:", list.GetAll())
}
