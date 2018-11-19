package main

import (
	"fmt"
	"net"
	"os"
	"sync"
	"testing"
	"time"
)

const (
	_NAME = "/tmp/unixdomain"
	_TYPE = "unixgram"
)

var (
	wg sync.WaitGroup
)

func server() {
	fmt.Println("server start...")

	conn, err := net.ListenUnixgram(_TYPE, &net.UnixAddr{_NAME, _TYPE})
	if err != nil {
		panic(err)
	}
	defer os.Remove(_NAME)

	buf := make([]byte, 1024)
	if n, err := conn.Read(buf); err != nil {
		panic(err)
	} else {
		fmt.Printf(">> %s\n", string(buf[:n]))
	}
	conn.Close()
}

func client() {
	fmt.Println("client start...")

	name := "/tmp/unixdomaincli"
	laddr := &net.UnixAddr{name, _TYPE}
	saddr := &net.UnixAddr{_NAME, _TYPE}
	conn, err := net.DialUnix(_TYPE, laddr, saddr)
	if err != nil {
		panic(err)
	}
	defer os.Remove(name)

	_, err = conn.Write([]byte("hello"))
	// _, err = conn.WriteTo([]byte("hello"), saddr)
	if err != nil {
		fmt.Println(err.(*net.OpError).Source)
		panic(err.(*net.OpError))
	}
	conn.Close()
}

func _TestUnixSocket(t *testing.T) {
	go func() {
		wg.Add(1)
		defer wg.Done()

		server()
	}()

	time.Sleep(1 * time.Second)

	go func() {
		wg.Add(1)
		defer wg.Done()

		client()
	}()

	wg.Wait()
}
