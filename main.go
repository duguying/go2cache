// Copyright 2018. All rights reserved.
// This file is part of go2cache project
// Created by duguying on 2018/5/11.

package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {

	go dial([]string{"127.0.0.1:9981", "192.168.2.227:60000"})

	go listen("127.0.0.1:9981")

	select {}
}

func listen(addr string) {
	ls, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		fmt.Println(err)
	}

	listener, err := net.ListenUDP("udp", &net.UDPAddr{IP: ls.IP, Port: ls.Port})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())

	data := make([]byte, 1024)

	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			fmt.Printf("error during read: %s", err)
		}
		fmt.Printf("<%s> %s\n", remoteAddr, data[:n])
		_, err = listener.WriteToUDP([]byte("world"), remoteAddr)
		if err != nil {
			fmt.Printf(err.Error())
		}
	}
}

func dial(addrs []string) {
	var writers []io.Writer
	for _, addr := range addrs {
		dst, err := net.ResolveUDPAddr("udp", addr)
		if err != nil {
			fmt.Println(err)
		}
		srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
		dstAddr := &net.UDPAddr{IP: dst.IP, Port: dst.Port}

		conn, err := net.DialUDP("udp", srcAddr, dstAddr)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("connect to: <%s>\n", conn.RemoteAddr())
		}

		writers = append(writers, conn)
	}

	connections := io.MultiWriter(writers...)

	defer func(writers []io.Writer) {
		for _, writer := range writers {
			conn := writer.(*net.UDPConn)
			conn.Close()
		}
	}(writers)

	for {
		connections.Write([]byte("hello"))
		time.Sleep(time.Second * 3)
	}
}
