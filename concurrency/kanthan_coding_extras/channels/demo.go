package main

import "fmt"

func main() {
	myChannel := make(chan string) // initialize a channel, data that pass through this is of type string

	go func() {
		myChannel <- "data" //send data to the channel
	}() // invoke the anon function

	//our main function reads from the channel
	msg := <-myChannel // main fn is going to take in the data from the channel

	fmt.Println(msg)

	// here the main function will wait for either the channel to get closed or wait till a msg is recieved from the channel

}
