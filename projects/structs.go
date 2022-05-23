package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Height int
	Weight float32
	Name   string
}

func main() {
	matt := Person{74, 180.4, "Matt Fisher"}
	jsonMatt, _ := json.Marshal(matt)
	fmt.Println(string(jsonMatt))
	fmt.Println("test")
	err := json.Unmarshal(jsonMatt, &matt)
	if err != nil {
		return
	}
}
