package main

import (
	"fmt"
)

func main() {
	myChannel := make(chan string) // initialize a channel, data that pass through this is of type string
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "data" //send data to the channel
	}() // invoke the anon function

	go func() {
		anotherChannel <- "cow" // puts data to the another channel
	}()

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}

}
