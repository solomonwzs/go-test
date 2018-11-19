package main

import (
	"bytes"
	"strings"
	"testing"
)

var _s0 = "0123456789"

func strAppend(n int) string {
	var tmp = ""
	for i := 0; i < n; i++ {
		tmp += _s0
	}
	return tmp
}

func sliceAppend(n int) string {
	tmp := []string{}
	for i := 0; i < n; i++ {
		tmp = append(tmp, _s0)
	}
	return strings.Join(tmp, "")
}

func buffAppend(n int) string {
	var buffer bytes.Buffer
	for i := 0; i < n; i++ {
		buffer.WriteString(_s0)
	}
	return buffer.String()
}

func _BenchmarkTestStrAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strAppend(1000)
	}
}

func _BenchmarkTestSliceAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceAppend(1000)
	}
}

func _BenchmarkTestBuffAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buffAppend(1000)
	}
}
