package main

import (
	"fmt"
	"io"
	"unsafe"
)

type MyStruct struct {
	Name  string
	Value int
}

func (s *MyStruct) Close() error {
	return nil
}

func main2() {
	m := make(map[int]int)
	delete(m, 1)
}

func main1() {
	var s interface{}
	s = new(MyStruct)

	if _, ok := s.(io.Closer); ok {
		fmt.Println(1)
	}

	// m := new(MyStruct)
}

func main3() {
	b0 := [4]byte{0, 1, 2, 3}
	b1 := []byte{0, 1, 2, 3}
	m := map[int]int{1: 2}
	s := "abcd"
	fmt.Println(b0, b1, m, s)
}

func foo(s string) {
	s += "111"
}

func bar(i byte) {
	fmt.Println(i)
}

func printlnString(ptr unsafe.Pointer, size int) {
	fmt.Printf("string: ")
	for offset := uintptr(0); offset < uintptr(size); offset++ {
		fmt.Printf("%c", *(*byte)(unsafe.Pointer(uintptr(ptr) + offset)))
	}
	fmt.Printf("\n")
}
