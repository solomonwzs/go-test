package main

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

type sCH struct {
	err error
}

func _TestChan(t *testing.T) {
	ch := make(chan int)

	foo := func(id int) {
		for i := range ch {
			fmt.Println(id, i)
		}
	}
	go foo(1)
	go foo(2)

	time.Sleep(1 * time.Second)
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	time.Sleep(1 * time.Second)

	err := errors.New("a")
	s := &sCH{err}
	err = errors.New("b")
	fmt.Println(s.err)
}

func _TestChan1(t *testing.T) {
	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		fmt.Println("close chan")
		close(ch)
	}()

	for {
		select {
		case i, ok := <-ch:
			if !ok {
				return
			}

			fmt.Println(i)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func _TestIO(t *testing.T) {
	fd, err := os.Open("/tmp/1")
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()

	p := make([]byte, 10)
	_, err = fd.Read(p)
	fmt.Println(err)
	_, err = fd.Read(p)
	fmt.Println(err)
}

func _TestChanClose(t *testing.T) {
	ch := make(chan struct{})

	close(ch)
	select {
	case <-ch:
		fmt.Println(1)
	default:
		fmt.Println(2)
	}
}

func _TestChanCounter(t *testing.T) {
	ch := make(chan int, 1000)

	for i := 0; i < 10; i++ {
		ch <- i
	}

	end := false
	for !end {
		select {
		case i := <-ch:
			fmt.Println(i)
		default:
			end = true
		}
	}

	buf := new(bytes.Buffer)
	p := make([]byte, 10)
	fmt.Println(buf.Read(p))
}
