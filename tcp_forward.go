package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)


func TcpForward(ip,rip string) {
	fmt.Printf(" Start TCP Target ")
	fmt.Printf("Listening %s\n", ip)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ip)
	lis, err := net.ListenTCP("tcp4", tcpAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			fmt.Println("建立连接错误:%v\n", err)
			continue
		}

		fmt.Println("Connecting from ", conn.RemoteAddr())
		go handle(conn, rip)
	}
}

func handle(sconn net.Conn, rip string) {
	tconn, err := net.Dial("tcp4", rip)
	if err != nil {
		log.Printf("[ERROR]Local dial failed "+rip+"\n", err)
		return
	}
	log.Printf("[%d] connected  to " + rip)
	var wg sync.WaitGroup
	go func(uconn net.Conn, tconn net.Conn) {
		wg.Add(1)
		defer wg.Done()
		_, _ = io.Copy(uconn, tconn)
		_ = uconn.Close()
	}(sconn, tconn)
	go func(uconn net.Conn, tconn net.Conn) {
		wg.Add(1)
		defer wg.Done()
		_, _ = io.Copy(tconn, uconn)
		_ = tconn.Close()
	}(sconn, tconn)
	wg.Wait()
}
