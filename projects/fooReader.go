package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FooReader struct{} //Create an emtpy type.

//Assign a read function to the FooReader struct.
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in >")
	return os.Stdin.Read(b)
}

type FooWriter struct{}

func (FooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out >")
	return os.Stdout.Write(b)
}

func main() {
	//Create an instance of the structs above.
	var (
		reader FooReader
		writer FooWriter
	)
	//The next many lines can be replaced with the following solution
	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Error reading/writing data")
	}

	//Text buffer
	//buffer := make([]byte, 4096)
	//s, err := reader.Read(buffer)
	//if err != nil {
	//	log.Fatalln("Unable to read data")
	//}
	//fmt.Printf("Read %d bytes from stdin\n", s)
	//s, err = writer.Write(buffer)
	//if err != nil {
	//	log.Fatal("Failed to write data")
	//}
	//fmt.Printf("Wrote %d bytes to stdout\n", s)
}
