package main

import (
	"log"
	"net"
	"os"
)

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		log.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	//ensure the teardown of the server on application termination.
	defer listener.Close()
	log.Println("The server is up and listening on port 6379")

	for {
		//Block until we recieve an incoming request
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error:", err)
			continue
		}

		//handler method.
		handleClient(conn)

	}
}

func handleClient(conn net.Conn) {

	//make sure we close the listener
	defer conn.Close()

	//read the data
	buf := make([]byte, 1024)

	input, err := conn.Read(buf)
	if err != nil {
		log.Println("Error", err)
	}

	log.Println("Received Data:", buf[:input])

	//write the same data back to the user.
	_, err = conn.Write(buf[:input])
	if err != nil {
		log.Println("Error:", err)
	}

}
