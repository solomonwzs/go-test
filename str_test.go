package main

import (
	"fmt"
	"strings"
	"testing"
	"unsafe"
)

func str2bytes(s string) []byte {
	x := (*[2]unsafe.Pointer)(unsafe.Pointer(&s))
	h := [3]unsafe.Pointer{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

var _s = strings.Repeat("a", 10240)

func _BenchmarkNormalConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := []byte(_s)
		_ = string(b)
	}
}

func _BenchmarkPointerConvert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := str2bytes(_s)
		_ = bytes2str(b)
	}
}

func TestStr2Bytes(t *testing.T) {
	fmt.Println(str2bytes("123"))
}
