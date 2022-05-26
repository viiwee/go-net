package main

import (
	"bufio"
	"log"
	"net"
)

func echo(conn net.Conn) {
	//When this function ends, close the connection
	defer conn.Close()

	//Create a reader and writer for the data coming in
	reader := bufio.NewReader(conn) //net.Conn has both a Read() and Write() function, making it both a reader and listener
	writer := bufio.NewWriter(conn)
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln("Failed to read string")
		}
		log.Printf("Read %d bytes: %s", len(s), s)

		log.Println("Writing data")

		//Cast s to a byte, and write it out to the connection
		writer.Write
		if _, err := writer.Write([]byte(s)); err != nil {
			println("Error writing data")
		}
		err = writer.Flush()
		if err != nil {
			println("Error flushing data")
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
