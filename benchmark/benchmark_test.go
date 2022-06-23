package main

import "testing"

func main() {

}

func IntMain(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMain(1, 2)
	}
}
