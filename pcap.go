package main

import (
	"time"

	"github.com/solomonwzs/goxutil/pcap"
)

func main() {
	h, err := pcap.OpenLive("eno1", 512, true, 1*time.Second)
	if err != nil {
		return
	}
	defer h.Close()

	h.SetFilter("port 80")
}
