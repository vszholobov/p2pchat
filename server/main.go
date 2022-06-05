package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	ServerHost = "172.25.0.1"
	ServerPort = "9988"
	ServerType = "udp"
)

func main() {
	fmt.Println("Server Running...")
	server, err := net.ListenPacket(ServerType, ServerHost+":"+ServerPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	fmt.Println("Listening on " + ServerHost + ":" + ServerPort)
	fmt.Println("Waiting for messages...")

	buffer := make([]byte, 1024)

	for {
		n, addr, err := server.ReadFrom(buffer)
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
			return
		}

		fmt.Printf("Packet-Received: bytes=%d from=%s\n", n, addr.String())
		fmt.Println(string(buffer))

		deadline := time.Now().Add(time.Second * time.Duration(10))
		err = server.SetWriteDeadline(deadline)
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
			return
		}

		n, err = server.WriteTo(buffer[:n], addr)
		if err != nil {
			fmt.Println("ERROR: " + err.Error())
			return
		}

		fmt.Printf("Packet-Written: bytes=%d to=%s\n", n, addr.String())
	}
}
