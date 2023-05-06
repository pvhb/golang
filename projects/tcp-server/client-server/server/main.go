package main

import (
	"fmt"
	"net"
	"os"
)

const (
	SERVER_TYPE = "tcp"
	SERVER_HOST = "localhost"
	SERVER_PORT = "5443"
)

func main() {
	fmt.Println("Starting the server...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting : ", err)
		}
		fmt.Println("Client connected")
		go processClient(connection)
	}

}

func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	len, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading : ", err)
	}
	fmt.Println("Received msg : ", string(buffer[:len]))

	fmt.Println("Printing the msg 3 times...")
	for i := 0; i < 3; i++ {
		fmt.Printf("Printing msg %d time : %s\n", i+1, string(buffer[:len]))
	}
	_, err = connection.Write([]byte("Thanks, got your msg : '" + string(buffer[:len]) + "'"))
	if err != nil {
		fmt.Println("Error writing : ", err)
	}
	connection.Close()

}
