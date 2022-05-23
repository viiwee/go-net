package main

import (
	"flag"
	"fmt"
	"net"
	"strings"
)

type resultPorts struct {
	Port int
	Open bool
}

func scan(host string, port string) {
	_, err := net.Dial("tcp", host+":"+port)
	if err == nil {
		fmt.Printf("Connection successful to %s:%s\n", host, port)
	}
}
func main() {
	portsPtr := flag.String("ports", "", "Comma separated list of ports")
	hostPtr := flag.String("host", "", "Hostname or IP address that you want to scan")
	flag.Parse()

	if *portsPtr == "" || *hostPtr == "" {
		fmt.Println("Please supply a value for ports and host")
	}
	ports := strings.Split(*portsPtr, ",")
	for i := 0; i < len(ports); i++ {
		go scan(*hostPtr, ports[i])
	}
}
