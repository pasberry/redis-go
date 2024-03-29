package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	//ensure the teardown of the server on application termination.
	defer listener.Close()
	fmt.Println("The server is up and listening on port 6379")

	for {
		conns := make(chan net.Conn)
		go func() {
			for {
				conn, err := listener.Accept()
				if err != nil {
					fmt.Println("Error accepting connection: ", err.Error())
					os.Exit(1)
				}
				conns <- conn
			}
		}()

		for c := range conns {
			go handleClient(c)
		}

	}
}

func handleClient(conn net.Conn) {

	//make sure we close the listener
	defer conn.Close()

	for {
		//read the data
		buf := make([]byte, 1024)

		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error", err)
		}

		res := []byte("+PONG\r\n")
		fmt.Print(res)

		//write the same data back to the user.
		_, err = conn.Write(res)
		if err != nil {
			fmt.Println("Error:", err)
		}

	}
}
