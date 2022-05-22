package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

//Declare variables and set their type
var int1 int

func main() {
	textPtr := flag.String("text", "", "Text to parse.")
	metricPtr := flag.String("metric", "chars", "Metric {chars|words|lines}")
	uniquePtr := flag.Bool("unique", false, "Measure unique values of a metric")
	// Actually parse the vars inputted by the user
	flag.Parse()

	if *textPtr == "" {
		flag.PrintDefaults()
		fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *textPtr, *metricPtr, *uniquePtr)
		os.Exit(1)
	}
	fmt.Printf("textPtr: %s, metricPtr: %s, uniquePtr: %t\n", *textPtr, *metricPtr, *uniquePtr)

	int1 = 12
	fmt.Printf("Your int as an int: %d\n", int1)
	fmt.Printf("Your int as a string: %s\n", strconv.Itoa(int1))
}
