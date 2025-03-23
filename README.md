# tcp-udp-golang

This repo is a monorepo of a TCP server and a UDP server written i golang. 
Both implementations have their respective client implementations to interact with the servers.

## Requirements
- [Golang 1.22.1](https://go.dev/)
- Linux OS

## Usage

### UDP Server

- Open a new terminal
- Navigate to where you stored the root of the project e.g. ``cd Desktop/tcp-udp-golang``
- Start the server by running ``go run ./udp/server/udp_server.go``
- Now your server is running and will wait for packets!

### UDP Client

- Open a new terminal
- Navigate to where you stored the root of the project e.g. ``cd Desktop/tcp-udp-golang``
- Start the client by running ``go run ./udp/client/udp_client.go``
- Now your client is running
- Follow the instructions by entering "127 0 0 1" as the server IP address
- Then enter "8080" as the server port
- Enter your message to the server
- Wait for the server's response!
