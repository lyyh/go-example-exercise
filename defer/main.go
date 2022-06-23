package main

import (
	"fmt"
	"os"
)

func main() {
	file := createFile("/tmp/file")
	defer closeFile(file)
	writeFile(file)
}

func createFile(p string) *os.File {
	file, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return file
}

func closeFile(f *os.File) {
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stdout, "err: %v\n", err)
		os.Exit(1)
	}
}

func writeFile(f *os.File) {
	fmt.Fprintln(f, "data")
}
