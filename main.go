package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func scanWorker(c chan int, r chan int, host string) {
	//For each integer (port) in the channel, do the following
	for port := range c {
		//Try to connect over the given port
		conn, err := net.Dial("tcp", host+":"+strconv.Itoa(port))
		if err != nil {
			//Failed connection
			//fmt.Println("Failed connection to " + host + ":" + strconv.Itoa(port) + "\n")
			fmt.Print(".")
			r <- 0
		} else {
			//Successful connection
			//fmt.Println("Connection successful to " + host + ":" + strconv.Itoa(port) + "\n")
			fmt.Print("!")
			r <- port
			conn.Close()
		}
		//Lower the WaitGroup count by one. No longer necessary
		//w.Done()
	}
}
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func expandEntry(entryPtr *string) []int {
	portsRaw := strings.Split(*entryPtr, ",")
	var ports []int
	for _, entry := range portsRaw {
		//Define regex patterns
		singlePattern := "^[0-9]+$"
		rangePattern := "^[0-9]+-[0-9]+$"
		//Try and match our entry to a range pattern
		rangeMatch, _ := regexp.MatchString(rangePattern, entry)
		singleMatch, _ := regexp.MatchString(singlePattern, entry)
		if singleMatch {
			port, _ := strconv.Atoi(entry)
			ports = append(ports, port)
			//fmt.Printf("Adding Port: %d (%s)\n", port, entry)
		} else if rangeMatch {
			//Matched a range. Turn into a list, and append
			start, _ := strconv.Atoi(strings.Split(entry, "-")[0])
			end, _ := strconv.Atoi(strings.Split(entry, "-")[1])
			portList := makeRange(start, end)
			ports = append(ports, portList...)
		} else {
			//Error condition, this string does not match our format
			fmt.Printf("Error: '%s' is not in a correct format.")
			os.Exit(1)
		}

	}
	return ports
}
func main() {
	portsPtr := flag.String("ports", "", "Comma separated list of ports")
	hostPtr := flag.String("host", "", "Hostname or IP address that you want to scanWorker")
	workersPtr := flag.Int("workers", 100, "Number of ports to scan simultaneously")
	flag.Parse()

	//Validate the
	if *portsPtr == "" || *hostPtr == "" {
		fmt.Println("Please supply a value for ports and host")
	}

	//Expand the entered ports into a list of all ports to run
	ports := expandEntry(portsPtr)

	//Make channel for the ports list and the returning results
	portChan := make(chan int, *workersPtr)
	portResults := make(chan int)
	//TODO: Make channel for the return of open ports

	//Create a wait group and add the # ports to it so that we can wait to close the program until the WG is finished
	//var scanWait sync.WaitGroup

	//Start workers for the scanning function
	for i := 0; i < cap(portChan); i++ {
		go scanWorker(portChan, portResults, *hostPtr)
	}

	//Add ports to the channel so workers can process them. This is a go func because we need to run the receiving end while this is still going
	go func() {
		for i := 0; i < len(ports); i++ {
			//scanWait.Add(1)
			curPort := ports[i]
			portChan <- curPort
		}
	}()

	//Receive Results
	var openPorts []int
	for i := 0; i < len(ports); i++ {
		//scanWait.Add(1)
		result := <-portResults
		if result != 0 {
			openPorts = append(openPorts, result)
		}
	}
	fmt.Println("\n")
	close(portChan)    //Close our channel, stopping our workers
	close(portResults) //Close our results channel
	//scanWait.Wait() //Wait until all of the go functions complete.
	sort.Ints(openPorts)
	for i := 0; i < len(openPorts); i++ {
		fmt.Println(*hostPtr + ":" + strconv.Itoa(openPorts[i]) + ": Open\n")
	}

}
