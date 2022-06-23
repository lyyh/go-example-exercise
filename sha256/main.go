package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 is string"
	h := sha256.New()
	h.Write([]byte(s))

	bs := h.Sum(nil)
	fmt.Printf("%x\n", bs)
}
