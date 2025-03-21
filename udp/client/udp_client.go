package main

import (
	"fmt"
	"syscall"
)


func main() {

	// Client socket to use to send data
	// Same structure as server socket
	fmt.Println("Creating client socket")
	clientUdpSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		fmt.Println("Error creating socket:", err)
		return
	}

	// Address client will send packet to
	serverAddr := &syscall.SockaddrInet4{
		Port: 8080,
		Addr: [4]byte{127,0,0,1},
	}

	message := []byte("This is a message from the client") // Message to send

	fmt.Println("Sending packet...")
	err = syscall.Sendto(clientUdpSocket, message, 0, serverAddr)
	if err != nil {
		fmt.Println("Error sending packet")
		return
	}

	fmt.Println("Packet sent!")
 	syscall.Close(clientUdpSocket)
	fmt.Println("Closing client...")
}

