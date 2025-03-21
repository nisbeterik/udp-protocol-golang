package main

import (
	"fmt"
	// "os"
	"syscall"
	)

type udpPacket struct {
	BytesRecieved int
	ClientAddress syscall.SockaddrInet4
	Err error
}


func main() {

	buffer := make([]byte, 512) // buffer to store packet data in
	//  Create socket
	// IPv4, UDP Socket, UDP
	udpSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		fmt.Println("Error creating socket", err)
		return
	}
	// set address that socket will bind to
	addr := &syscall.SockaddrInet4{
		Port: 8080,					// bind socket to port 8080
		Addr: [4]byte{127,0,0,1},	// set it to interface only with localhost
	}

	
	
	err = syscall.Bind(udpSocket, addr) // bind the socket to the address/port. this will route traffic to the socket
	if err != nil {
		fmt.Println("Error binding sockets:", err)
		return
	}

	bytes, clientAddress, err :=	syscall.Recvfrom(udpSocket, buffer, 0)
	if err != nil {
		fmt.Println("Error receiving data:", err)
		return
	}

	clientAddressIPv4, ok := clientAddress.(*syscall.SockaddrInet4)
	if !ok {
		fmt.Println("Client address type incorrect")
		return
	}


	tempPacket := udpPacket{bytes, *clientAddressIPv4, err}
	fmt.Println(tempPacket)
	syscall.Close(udpSocket)
}