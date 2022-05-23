package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

//Declare variables and set their type
var int1 int

//Testing with channels
func Runlist_normal(n int) int {
	if n < 2 {
		println(n)
		return n
	}
	return Runlist_normal(n-1) + Runlist_normal(n-2)
}

func Runlist_chan(c chan int, quit chan bool) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Fin Fib")
			return
		}
	}
}

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
	str1 := strconv.Itoa(int1)
	fmt.Printf("Your int as a string: %s\n", str1)
	var int2, err = strconv.Atoi(str1)
	if err != nil {
		log.Fatalf("Could not convert string, failed with %s", str1)
	}
	fmt.Printf("Your int as a string as an int: %d\n", int2)

	//Channel Testing
	//Make the channel
	c := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- true
	}()
	Runlist_chan(c, quit)

	//Make a number that the list can add to
	Runlist_normal(10)

}
