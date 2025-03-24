package main

import (
	"bufio"
	"fmt"
	"os"
	"syscall"
)

func main() {

	buffer := make([]byte, 512) // buffer to parse server response
	ip, port, message, err := enterDetails()
	if err != nil {
		fmt.Println("Error entering details:", err)
		return
	}

	clientUdpSocket, err := createSocket()
	if err != nil {
		fmt.Println("Error creating client socket:", err)
		return
	}

	// Address client will send packet to
	serverAddr := setServerAddress(port, ip)

	err = sendPacket(clientUdpSocket, message, serverAddr)
	if err != nil {
		fmt.Println("Error sending packet:", err)
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

func enterDetails() ([4]byte, int, []byte, error){
	fmt.Println("Enter the server IP address")
	ip := [4]byte{0, 0, 0, 0}

	fmt.Scanln(&ip[0], &ip[1], &ip[2], &ip[3])

	fmt.Println("Enter the server port")
	port := 0
	fmt.Scanln(&port)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a message you want to send")
	message, err := reader.ReadBytes('\n')
	if err != nil {

		fmt.Println("error entering message:", err)
		return ip, -1, nil, err
	}
	return ip, port, message, nil

}

func createSocket() (int, error) {
	fmt.Println("Creating client socket")
	clientUdpSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		fmt.Println("Error creating socket:", err)
		return -1, err
	}
	return clientUdpSocket, nil
}

func setServerAddress(port int, ip [4]byte) *syscall.SockaddrInet4 {
	return &syscall.SockaddrInet4{
		Port: port,
		Addr: ip,
	}
}

func sendPacket(fd int, message []byte, addr *syscall.SockaddrInet4) error {
	fmt.Println("Sending packet...")
	return syscall.Sendto(fd, message, 0, addr)
}