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
	fmt.Println("Udp socket: ", udpSocket)
}