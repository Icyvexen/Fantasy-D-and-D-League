package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, err := net.Listen("tcp", ":8081")

	if err != nil {
		panic(err)
	}
	// accept connection on port
	conn, err := ln.Accept()

	if err != nil {
		panic(err)
	}

	// run loop forever (or until ctrl-c)
	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')

		// output message received
		fmt.Print("Message Received: ", string(message))

		newmessage := strings.ToUpper(message)
		fmt.Println("newmessage: ", newmessage)
		if strings.Contains(newmessage, "CLOSE") {
			fmt.Println("Close statement received; closing Server")
			conn.Write([]byte("CLOSE\n"))
			conn.Close()
			break
		} else {
			// send new string back to client
			conn.Write([]byte(message + "\n"))
		}
	}
}
