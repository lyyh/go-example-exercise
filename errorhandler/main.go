package main

import (
	"errors"
	"fmt"
)

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func f1(arg int) (int, error) {
	if arg == 1 {
		return -1, errors.New("can't work with 42")
	}
	return arg, nil
}
func main() {
	for _, v := range []int{1, 2} {
		if ret, err := f1(v); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ret)
		}
	}

	_, e := f2(42)
	// 类型断言
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
