package main

import "testing"

func Benchmark_FibIterative_10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fiboI(10)
	}
}

func Benchmark_FibTailRecursive_10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fiboTail(10)
	}
}

func Benchmark_FibRecursive_10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fiboR(10)
	}
}
