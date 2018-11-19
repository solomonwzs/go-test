package main

import (
	"fmt"
	"testing"
	"unsafe"
)

type ptrStruct struct {
	a  int8
	b  int16
	ch chan int
}

func TestPtr(t *testing.T) {
	ss := ptrStruct{}
	st := &ptrStruct{1, 2, make(chan int)}

	bp := &(st.ch)
	st0 := (*ptrStruct)(
		unsafe.Pointer(uintptr(unsafe.Pointer(bp)) - unsafe.Offsetof(ss.ch)))

	fmt.Printf("%v %v\n", st, st0)
}
