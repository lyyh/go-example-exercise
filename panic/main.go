package main

import "os"

func main() {
	// panic("xxxxx")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
