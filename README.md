# udp-protocol-golang

This repo is a UDP server written in Golang

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
- In the console, enter "127 0 0 1" as the server address
- Enter "8080" as the server port
- Client will receive the current time in Sweden
