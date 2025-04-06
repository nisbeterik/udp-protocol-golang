package main

import (
	"fmt"
	// "os"
	"github.com/nisbeterik/tcp-udp-golang/udp/packetHandler"
	"net"
	"os"
	"strconv"
	"sync"
	"syscall"
)

type UDPSocket struct {
	FileDescriptor int
	ServerAddress  syscall.SockaddrInet4
}

func main() {

	buffer := make([]byte, 512) // buffer to store packet data in
	var port int
	var ip [4]byte
	var err error

	if len(os.Args) < 3 {
		port = 8080
		ip = [4]byte{127, 0, 0, 1}
	} else {
		port, ip, err = parseArguments()
		if err != nil {
			fmt.Println("Error parsing arguments:", err)
			return
		}
	}
	var udpSocket UDPSocket // udpSocket object
	var wg sync.WaitGroup
	// Create socket, IPv4
	err = createSocket(&udpSocket)
	if err != nil {
		fmt.Println("Error creating socket:", err)
		return
	}

	// set address that socket will bind to
	// Port 8080 and local host ip
	err = setAddress(&udpSocket, port, ip)
	if err != nil {
		fmt.Println("Error setting address", err)
		return
	}

	err = bindSocket(&udpSocket)
	if err != nil {
		fmt.Println("Error binding socket:", err)
		return
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("Listening for packets...")
		bytes, clientAddress, err := syscall.Recvfrom(udpSocket.FileDescriptor, buffer, 0)
		if err != nil {
			fmt.Println("Error receiving data:", err)
			break
		}
		fmt.Println("Packet received!")
		wg.Add(1)
		go func(bytes int, clientAddress syscall.Sockaddr, data []byte, err error) {
			defer wg.Done()
			packetHandler.ProcessPacket(udpSocket.FileDescriptor, bytes, clientAddress, data, err)
		}(bytes, clientAddress, buffer[:bytes], err)

	}
	wg.Wait()
	syscall.Close(udpSocket.FileDescriptor) // close socket when server is done
	fmt.Println("Closing server...")
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
		return fmt.Errorf("invalid port number: %d, must between 0 and %d", port, 0xFFFF)
	}

	udpSocket.ServerAddress.Port = port
	udpSocket.ServerAddress.Addr = ip
	fmt.Println("Address set!")
	return nil
}

func bindSocket(udpSocket *UDPSocket) error {
	fmt.Println("Binding socket...")
	err := syscall.Bind(udpSocket.FileDescriptor, &udpSocket.ServerAddress)
	if err != nil {
		fmt.Println("Error binding socket:", err)
		return err
	}
	return nil
}

func parseArguments() (int, [4]byte, error) {
	num, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return -1, [4]byte{}, err
	}

	ip := net.ParseIP(os.Args[2])
	if ip == nil {
		fmt.Println("Error parsing IP:", err)
		return -1, [4]byte{}, err
	}

	ip4 := ip.To4()
	if ip4 == nil {
		fmt.Println("Error converting to IP4:", err)
		return -1, [4]byte{}, err
	}
	var ipArr [4]byte
	copy(ipArr[:], ip4)
	return num, ipArr, nil
}
