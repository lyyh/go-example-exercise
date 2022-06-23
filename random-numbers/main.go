package main

import (
	"fmt"
	"math/rand"
)

func main() {
	s2 := rand.NewSource(42)
	fmt.Println(s2)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
}
