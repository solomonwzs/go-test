package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type hello_s struct {
	Hello string `json:"hello"`
	World string `json:"world"`
}

func _BenchmarkJson0(b *testing.B) {
	m := map[string]string{
		"hello": "123",
		"world": "345",
	}
	for i := 0; i < b.N; i++ {
		json.Marshal(m)
	}
}

func _BenchmarkJson1(b *testing.B) {
	m := hello_s{"123", "456"}
	for i := 0; i < b.N; i++ {
		json.Marshal(m)
	}
}

func _BenchmarkJson2(b *testing.B) {
	f := `{"hello":"%s","world":"%s"}`
	for i := 0; i < b.N; i++ {
		_ = []byte(fmt.Sprintf(f, "123", "456"))
	}
}

func _BenchmarkJson3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := &bytes.Buffer{}
		buf.WriteString(`{"hello":"`)
		buf.WriteString("123")
		buf.WriteString(`","world":"`)
		buf.WriteString("456")
		buf.WriteString(`"}`)
		_ = buf.Bytes()
	}
}

func _BenchmarkJson4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b := &strings.Builder{}
		b.WriteString(`{"hello":"`)
		b.WriteString("123")
		b.WriteString(`","world":"`)
		b.WriteString("456")
		b.WriteString(`"}`)
		_ = []byte(b.String())
	}
}
