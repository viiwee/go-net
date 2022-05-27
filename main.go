package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

const remoteHost string = "example.com"
const webProtocol string = "80"

func relayHTTP(src net.Conn) {
	//Connect to our remote host
	fmt.Printf("Dialing %s:%s\n", remoteHost, webProtocol)
	dst, err := net.Dial("tcp", remoteHost+":"+webProtocol)
	if err != nil {
		println("Error dialing remote host")
		dst.Close()
		return
	}
	defer dst.Close()

	//Copy traffic from source to remote
	go func() {
		fmt.Println("source -> dest")
		written, err := io.Copy(dst, src)
		if err != nil {
			log.Fatalln("Error reading/writing")
		}
		fmt.Printf(
			"Sending %d bytes from %s to %s\n",
			written,
			src.RemoteAddr(),
			dst.RemoteAddr())
	}()

	//Copy responses from remote to source
	fmt.Println("dest -> source")
	written, err := io.Copy(src, dst)
	if err != nil {
		log.Fatalln("Error reading/writing")
	}
	fmt.Printf(
		"Sending %d bytes from %s to %s\n",
		written,
		dst.RemoteAddr(),
		src.RemoteAddr())
}

func main() {
	//Setup listener for incoming traffic on 443
	//Accept new connections, and then forward the connection into the relay
	//Loop back
	fmt.Println("Listening on :80 for new connections")
	listener, err := net.Listen("tcp", ":"+webProtocol)
	if err != nil {
		log.Fatalln("Could not setup listener on :80, already listening?")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("Unable to accept new connection")
		}
		fmt.Printf("Started new connection, %s\n", conn.RemoteAddr())
		go relayHTTP(conn)
	}
}
