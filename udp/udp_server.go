package main

import (
	"fmt"
	// "os"
	"syscall"
	)


func main() {
	
	fmt.Println(syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, syscall.IPPROTO_UDP))
	
}