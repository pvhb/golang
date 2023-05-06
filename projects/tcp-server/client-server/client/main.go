package main

import (
	"fmt"
	"net"
)

const (
	SERVER_TYPE = "tcp"
	SERVER_HOST = "localhost"
	SERVER_PORT = "5443"
)

func main() {
	fmt.Println("Connecting to the server...")
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}

	var input string
	fmt.Println("Enter any text : ")
	fmt.Scan(&input)
	_, err = conn.Write([]byte(input))
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	buffer := make([]byte, 1024)
	len, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Println("Received : ", string(buffer[:len]))
	defer conn.Close()
}
