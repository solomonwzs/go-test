package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func _TestTime(t *testing.T) {
	f := "2006-01-02 15:04:05"
	fmt.Println(time.Parse(f, "2018-02-31 18:02:32"))

	f = "15:04:05"
	fmt.Println(time.Parse(f, "18:02:32"))

	fmt.Println(time.Time{})

	var i int = 0xffffffff + 1
	fmt.Println(i)
	// var c1 chan int = nil
	// select {
	// case <-c1:
	// 	fmt.Println(1)
	// case <-time.After(1 * time.Second):
	// 	fmt.Println(2)
	// }

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			time.Sleep(100 * time.Millisecond)
			fmt.Println(j)
		}(i)
	}
	wg.Wait()
}

func _TestTimer(t *testing.T) {
	b := true
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; b; i++ {
			fmt.Printf("%d ", i)
			time.Sleep(5 * time.Millisecond)
		}
		fmt.Printf("\n")
	}()
	time.Sleep(1 * time.Second)
	b = false
	wg.Wait()
}
