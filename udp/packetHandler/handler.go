package packetHandler

import (
	"fmt"
	"syscall"
	"time"
)

type UDPPacket struct {
	BytesRecieved int
	ClientAddress syscall.SockaddrInet4
	Message       []byte
	Err           error
}

func ProcessPacket(fd int, bytes int, clientAddress syscall.Sockaddr, message []byte, err error) {
	if bytes != 16 {
		fmt.Println("Invalid packet size")
		return
	}

	if message[12] != 49 || message[13] != 0x4E || message[14] != 49 || message[15] != 52 {
		fmt.Println("Invalid magic marker")
		return
	}

	var packet UDPPacket
	packet.Err = err

	clientAddressIPv4, err := extractIPv4(clientAddress)
	if err != nil {
		fmt.Println("Error processing packet")
		return
	}

	packet.ClientAddress = *clientAddressIPv4
	packet.BytesRecieved = bytes

	message = packTime(message)
	packet.Message = message

	time.Sleep(3 * time.Second)

	printMessage(packet)

	err = sendResponse(packet, fd)
	if err != nil {
		fmt.Println("Error sending response")
		return
	}
}

func packTime(buf []byte) []byte {
	now := uint32(time.Now().Unix())
	buf[12] = byte(now >> 24)
	buf[13] = byte(now >> 16)
	buf[14] = byte(now >> 8)
	buf[15] = byte(now)
	printTime(now)
	return buf
}

func printTime(now uint32) {
	t := time.Unix(int64(now), 0)

	location, err := time.LoadLocation("Europe/Stockholm")
	if err != nil {

		fmt.Println("Failed to load timezone:", err)
		return
	}

	localTime := t.In(location)

	fmt.Println("Swedish time:", localTime.Format("2006-01-02 15:04:05 MST"))
}

func extractIPv4(clientAddress syscall.Sockaddr) (*syscall.SockaddrInet4, error) {
	clientAddressIPv4, ok := clientAddress.(*syscall.SockaddrInet4)
	if !ok {
		return nil, fmt.Errorf("error extracting IPv4")
	}
	return clientAddressIPv4, nil
}

func printMessage(packet UDPPacket) {
	fmt.Printf("Client address: %d.%d.%d.%d:%d\n", packet.ClientAddress.Addr[0], packet.ClientAddress.Addr[1], packet.ClientAddress.Addr[2], packet.ClientAddress.Addr[3], packet.ClientAddress.Port)
	fmt.Printf("Bytes received: %d\n", packet.BytesRecieved)
	fmt.Print("Message bytes from client (binary): ")
	for i, b := range packet.Message {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%08b", b)
	}
	fmt.Println()

}

// Change to take UDPPacket as parameter
func sendResponse(packet UDPPacket, fd int) error {
	err := syscall.Sendto(fd, packet.Message, 0, &packet.ClientAddress)
	if err != nil {
		return fmt.Errorf("error responding to client")
	}
	return nil
}
