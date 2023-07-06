package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

type argError struct {
	arg     int
	problem string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.problem)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, v := range []int{7, 42} {
		if res, err := f1(v); err != nil {
			fmt.Println("f1 failed:", err)
		} else {
			fmt.Println("f1 worked:", res)
		}
	}

	for _, v := range []int{7, 42} {
		if res, err := f2(v); err != nil {
			fmt.Println("f2 failed:", err)
		} else {
			fmt.Println("f2 worked:", res)
		}
	}

	_, err := f2(42)
	if argErr, ok := err.(*argError); ok {
		fmt.Println(argErr.arg)
		fmt.Println(argErr.problem)
	}
}
