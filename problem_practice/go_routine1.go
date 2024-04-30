package main

import (
	"fmt"
	"time"
)

func show(input string) {
	fmt.Println(input)
}

func main() {
	go show("1")
	go show("2")
	go show("3")

	time.Sleep(time.Second * 2)
	fmt.Println("Hello from 'main' go routine")
}
