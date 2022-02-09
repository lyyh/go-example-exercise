package main

import "fmt"

func whatType(i interface{}) {
	switch t := i.(type) {
	case string:
		fmt.Println("I'm a string")
	case int:
		fmt.Println("I'm an int")
	default:
		fmt.Printf("Don't know type %T\n", t)
	}
}

func main() {
	var s = "hello"

	whatType(s)
}
