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

type UDPSocket struct {
	FileDescriptor int
	ServerAddress syscall.SockaddrInet4
}


func main() {

	buffer := make([]byte, 512) // buffer to store packet data in
	
	var udpSocket UDPSocket // udpSocket object
	var err error	// err variable


	// Create socket, IPv4
	err = createSocket(&udpSocket)
	if err != nil {
		fmt.Println("Error creating socket:", err)
		return
	}


	// set address that socket will bind to
	// Port 8080 and local host ip
	err = setAddress(&udpSocket, 8080, [4]byte{127,0,0,1})
	if err != nil {
		fmt.Println("Error setting address", err)
		return
	}



	/*
	fmt.Println("Binding socket..:")	
	err = syscall.Bind(udpSocket, addr) // bind the socket to the address/port. this will route traffic to the socket
	if err != nil {
		fmt.Println("Error binding sockets:", err)
		return
	}

	fmt.Println("Listening for packets...")
	bytes, clientAddress, err :=	syscall.Recvfrom(udpSocket, buffer, 0)
	if err != nil {
		fmt.Println("Error receiving data:", err)
		return
	}

	fmt.Println("Packet received!")
	clientAddressIPv4, ok := clientAddress.(*syscall.SockaddrInet4)
	if !ok {
		fmt.Println("Client address type incorrect")
		return
	}


	tempPacket := udpPacket{bytes, *clientAddressIPv4, err}
	fmt.Println(tempPacket)
	message := string(buffer[:tempPacket.BytesRecieved])
	fmt.Println("Message from client: " + message)
	syscall.Close(udpSocket)
	fmt.Println("Closing server...")*/
}


// Function creates socket and sets the filedescriptor within the object 
func createSocket(udpSocket *UDPSocket) error {

	fmt.Println("Creating socket...")
	var err error
	udpSocket.FileDescriptor, err = syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP)
	if err != nil {
		fmt.Println("Error creating socket", err)
		return err
	}
	fmt.Println("Socket created!")
	return nil
}

// Function sets the address of the socket
// checks if the port number is valid
func setAddress(udpSocket *UDPSocket, port int, ip [4]byte) error {

	fmt.Println("Setting address...")	
	if port < 0 || port > 0xFFFF {
		return fmt.Errorf("Invalid port number: %d, must between 0 and %d", port, 0xFFFF)
	}

	udpSocket.ServerAddress.Port = port
	udpSocket.ServerAddress.Addr = ip
	fmt.Println("Address set!")
	return nil
}

func bindSocket(udpSocket *UDPSocket) error {
	fmt.Println("Binding socket...")
	var err error
	err = syscall.Bind(udpSocket.FileDescriptor, &udpSocket.ServerAddress)
	if err != nil {
		fmt.Println("Error binding socket:", err)
		return err
	}
	return nil 
}