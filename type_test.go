package main

import "testing"

const (
	_T_INT   = 0x00
	_T_FLOAT = 0x01
)

type typ struct {
	t    uint8
	data interface{}
}

func _BenchmarkType1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan *typ, 100)
		go func() {
			for j := 0; j < 1000; j++ {
				ch <- &typ{
					t:    _T_INT,
					data: j,
				}
			}
			close(ch)
		}()

		sum := 0
		for j := range ch {
			switch j.t {
			case _T_INT:
				sum += j.data.(int)
			}
		}
	}
}

func _BenchmarkType2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ch := make(chan *typ, 100)
		go func() {
			for j := 0; j < 1000; j++ {
				ch <- &typ{
					t:    _T_INT,
					data: j,
				}
			}
			close(ch)
		}()

		sum := 0
		for j := range ch {
			switch j.data.(type) {
			case int:
				sum += j.data.(int)
			}
		}
	}
}
