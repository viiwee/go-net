package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	//When this function ends, close the connection
	defer conn.Close()

	//Buffer for data
	b := make([]byte, 512)
	for {
		//Receive data from the connection
		//Standard read format (size of data, error)
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			//EOF means the client closed the connection
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			//Unknown error
			log.Println("Unexpected Error!")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))

		//Send the data back to the client
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data!")
		}
	}
}

func main() {
	//Configure a listener
	listener, err := net.Listen("tcp", ":25565")
	if err != nil {
		log.Fatalln("Unable to bind to port, is port already listening?")
	}
	log.Println("Listening on 0.0.0.0:25565")
	for {
		//Loop, waiting for connections and then setting up the echo server for each
		conn, err := listener.Accept() //This will wait until a client connects. And then it will return a conn instance
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		//Handle the connection using the echo function we created earlier
		go echo(conn)
	}
}
