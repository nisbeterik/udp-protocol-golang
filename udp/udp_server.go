package main

import (
	"fmt"
	// "os"
	"syscall"
	)


func main() {
	//  Create socket
	// IPv4, UDP Socket, UDP
	udpSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		fmt.Println("Error creating socket", err)
		return
	}

	addr := &syscall.SockaddrInet4{
		Port: 8080,					// bind socket to port 8080
		Addr: [4]byte{127,0,0,1},	// set it to interface only with localhost
	}
	
	syscall.Bind(udpSocket, addr) // bind the socket to the address/port. this will route traffic to the socket

}