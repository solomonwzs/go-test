package main

import (
	"bytes"
	"io"
	"testing"
	"unsafe"
)

const (
	PH_NULL = 0x00
	PH_BASE = 0x01
)

type PayloadHeaderInfo struct {
	Version   uint8
	Type      uint8
	Unused    uint8
	ExtType   uint8
	ExtLength uint32
}

type PayloadHBase struct {
	RequestID uint64
	TimeSec   uint64
}

func headerReader() io.Reader {
	buf := []byte{
		0, 1, 1, 0, 27, 0, 0, 0, 0, 0, 0, 0, 21, 205, 91, 7,
		0, 0, 0, 0, 127, 249, 206, 90, 0, 0, 0, 0, 123, 34,
		97, 34, 58, 49, 50, 51, 44, 34, 98, 34, 58, 34, 104,
		101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 34, 125}
	return bytes.NewBuffer(buf)
}

func TestDecode(t *testing.T) {
	r := headerReader()

	pInfo := make([]byte, 8)
	r.Read(pInfo)
	info := (*PayloadHeaderInfo)(unsafe.Pointer(&pInfo[0]))

	if info.Type == PH_BASE {
		pBase := make([]byte, 16)
		r.Read(pBase)
	}
}
