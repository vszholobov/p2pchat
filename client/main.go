package main

import (
	"fmt"
	"net"
	"time"
)

const (
	ServerHost = "172.25.0.1"
	ServerPort = "9988"
	ServerType = "udp"
)

func main() {
	//establish connection
	raddr, err := net.ResolveUDPAddr(ServerType, ServerHost+":"+ServerPort)
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP(ServerType, nil, raddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for {
		_, err = conn.Write([]byte("HELLO FROM CLIENT!"))

		if err != nil {
			fmt.Println("Error listening:", err.Error())
			return
		}

		buffer := make([]byte, 1024)

		// Set a deadline for the ReadOperation so that we don't
		// wait forever for a server that might not respond on
		// a resonable amount of time.
		deadline := time.Now().Add(500)
		err = conn.SetReadDeadline(deadline)
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			return
		}

		nRead, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error listening:", err.Error())
			return
		}

		fmt.Printf("Packet-Received: bytes=%d from=%s\n", nRead, addr.String())
		time.Sleep(time.Duration(5) * time.Second)
	}
}
