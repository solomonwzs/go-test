package main

import (
	"testing"
)

const capacity = 1024000

func arraySum(b [capacity]int) int {
	res := 0
	for _, i := range b {
		res += i
	}
	return res
}

func sliceSum(b []int) int {
	res := 0
	for _, i := range b {
		res += i
	}
	return res
}

func _BenchmarkArray(b *testing.B) {
	bb := [capacity]int{}
	for i := 0; i < b.N; i++ {
		_ = arraySum(bb)
	}
}

func _BenchmarkSlice(b *testing.B) {
	bb := make([]int, capacity)
	for i := 0; i < b.N; i++ {
		_ = sliceSum(bb)
	}

}
