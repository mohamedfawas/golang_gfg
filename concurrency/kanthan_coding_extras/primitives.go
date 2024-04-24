package main

import (
	"fmt"
	"time"
)

func someFunc(num string) {
	fmt.Println(num)
}

func main() {
	go someFunc("1") //this is a go routine now
	go someFunc("2")
	go someFunc("3")
	// all the above codes are asynchronous

	//main is not associated with our go routines, it's not gonna wait for go routine to finish execution
	time.Sleep(time.Second * 2) // sleep for two seconds
	fmt.Println("hai")
}
