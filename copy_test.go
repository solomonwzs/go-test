package main

import (
	"net"
	"sync"
	"testing"
)

const SRC_SIZE = 32 * 1024 * 1024

func setupTcpConn(tb testing.TB) (sConn, cConn net.Conn) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		tb.Fatal(err)
	}
	ch := make(chan net.Conn, 1)
	go func() {
		if conn, err := listener.Accept(); err != nil {
			tb.Fatal(err)
		} else {
			ch <- conn
		}
	}()

	addr := listener.Addr().String()
	cConn, err = net.Dial("tcp", addr)
	if err != nil {
		tb.Fatal(err)
	}
	sConn = <-ch
	return
}

func benchmarkCopy(b *testing.B, size int) {
	src := make([]byte, SRC_SIZE)
	dst := make([]byte, size)
	sc, cc := setupTcpConn(b)

	var wg0 sync.WaitGroup
	wg0.Add(1)
	go func() {
		wg0.Done()
		p := make([]byte, size)
		for {
			if _, err := sc.Read(p); err != nil {
				return
			}
		}
	}()

	for i := 0; i < b.N; i++ {
		ch := make(chan []byte, 1024)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range ch {
				copy(dst, p)
				cc.Write(dst[:len(p)])
			}
		}()

		n := 0
		for n < len(src) {
			m := n + size
			if m > len(src) {
				m = len(src)
			}
			ch <- src[n:m]
			n += size
		}
		close(ch)
		wg.Wait()
	}
	sc.Close()
	cc.Close()
	wg0.Wait()
}

func _BenchmarkCopy0(b *testing.B) {
	benchmarkCopy(b, 0xffff)
}

func _BenchmarkCopy1(b *testing.B) {
	benchmarkCopy(b, 60000)
}
