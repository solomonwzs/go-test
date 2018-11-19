package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
)

type config struct {
	Title string `toml:"title"`
	Owner *owner `toml:"owner"`
}

type owner struct {
	Name string `toml:"name"`
	Org  string `toml:"organization"`
}

func _TestToml(t *testing.T) {
	f, err := os.Open("./conf.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	var c config
	var m toml.MetaData
	if m, err = toml.Decode(string(b), &c); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", c)
	fmt.Printf("%+v\n", m)
}
