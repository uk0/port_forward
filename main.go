package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s  0.0.0.0:3389  123.46.12.91:3389  udp or tcp\n", os.Args[0])
		os.Exit(1)
	}
	ip := os.Args[1]
	rip := os.Args[2]
	ConnType := os.Args[3]

	switch strings.ToUpper(ConnType) {
	case "UDP":
		fmt.Printf("Type UDP")
		forwarder, err := UdpForward(ip, rip, DefaultTimeout)
		if err != nil {
			panic(err)
		}
		forwarder.Connected()
	case "TCP":
		fmt.Printf("Type TCP")
		TcpForward(ip, rip)
	}

	select {}
}
