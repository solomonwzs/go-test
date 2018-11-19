package main

import (
	"fmt"
	"regexp"
	"testing"
)

func _TestRe(t *testing.T) {
	re := regexp.MustCompile("byes=(\\d*)-(\\d*)")

	s1 := "bytes=-1000, 2000-6576, 19000-"
	fmt.Println(re.FindAllStringSubmatch(s1, -1))
}
