package main

import (
	"os"
	"sync"
	"testing"
)

func TestFileWrite(t *testing.T) {
	f, err := os.OpenFile("/tmp/ben", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString("hello\n")
	f.Close()

	f, err = os.OpenFile("/tmp/ben", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		t.Fatal(err)
	}
	f.WriteString("world\n")
	f.Close()
}

func _BenchmarkFileWrite(b *testing.B) {
	f, err := os.OpenFile("/tmp/ben", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()

	line := make([]byte, 1024)
	for i := 0; i < b.N; i++ {
		f.Write(line)
	}
}

func _BenchmarkFileWriteWithMutex(b *testing.B) {
	f, err := os.OpenFile("/tmp/ben", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	lock := &sync.Mutex{}

	line := make([]byte, 1024)
	for i := 0; i < b.N; i++ {
		lock.Lock()
		f.Write(line)
		lock.Unlock()
	}
}
