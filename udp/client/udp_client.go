package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
)

func main() {

	// TODO: refactor this file to 
	// be more modular and readable
	buffer := make([]byte, 512)
	fmt.Println("Enter the server IP address")
	ip := [4]byte{0, 0, 0, 0}

	fmt.Scanln(&ip[0], &ip[1], &ip[2], &ip[3])

	fmt.Println("Enter the server port")
	port := 0
	fmt.Scanln(&port)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a message you want to send")
	message, _ := reader.ReadBytes('\n')

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
		Port: port,
		Addr: ip,
	}

	fmt.Println("Sending packet...")
	err = syscall.Sendto(clientUdpSocket, message, 0, serverAddr)
	if err != nil {
		fmt.Println("Error sending packet")
		return
	}

	fmt.Println("Packet sent!")
	fmt.Println("Waiting for response...")
	bytes, serverAddress, err := syscall.Recvfrom(clientUdpSocket, buffer, 0)
	if err != nil {
		fmt.Println("Error receiving response:", err)
		return
	}
	fmt.Println("Message received!")
	fmt.Println(string(buffer[:bytes]))
	addr := serverAddress.(*syscall.SockaddrInet4)
	fmt.Printf("From server: %v:%v\n", addr.Addr, addr.Port)

	syscall.Close(clientUdpSocket)
	fmt.Println("Closing client...")
}
