package main

import (
	"fmt"
	"math/rand"
	"syscall"
	"time"
)

func main() {
	// TODO: Send packet using raw sockets that form valid UDP headers.
	buffer := make([]byte, 2) // buffer to parse server response
	ip, port, err := enterDetails()
	if err != nil {
		fmt.Println("Error entering details:", err)
		return
	}

	message := genMessage()

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
	err = getResponse(clientUdpSocket, buffer)
	if err != nil {
		fmt.Println("Error waiting for response", err)
		return
	}

	syscall.Close(clientUdpSocket)
	fmt.Println("Closing client...")
}

func enterDetails() ([4]byte, int, error) {
	fmt.Println("Enter the server IP address")
	ip := [4]byte{0, 0, 0, 0}

	fmt.Scanln(&ip[0], &ip[1], &ip[2], &ip[3])
	fmt.Println("Enter the server port")
	port := 0
	fmt.Scanln(&port)

	return ip, port, nil

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

func getResponse(fd int, buf []byte) error {
	bytes, serverAddress, err := syscall.Recvfrom(fd, buf, 0)
	if err != nil {
		return fmt.Errorf("error receiving response")
	}
	fmt.Println("Message received!")
	fmt.Println(string(buf[:bytes]))
	addr := serverAddress.(*syscall.SockaddrInet4)
	fmt.Printf("From server: %v:%v\n", addr.Addr, addr.Port)
	return nil
}

func genMessage() []byte {
	packet := make([]byte, 16)

	packet[0] = 0b11100011 // LI = 3, VN = 4, Mode = 3
	packet[1] = 0          // Stratum
	packet[2] = 6          // Poll Interval
	packet[3] = 0xEC       // Precision
	packet[4], packet[5], packet[6], packet[7] = 0, 0, 0, 0
	packet[8], packet[9], packet[10], packet[11] = 0, 0, 0, 0
	packet[12], packet[13], packet[14], packet[15] = 49, 0x4E, 49, 52 // "IN14"

	return packet
}
