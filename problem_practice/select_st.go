package main

import (
	"fmt"
	"time"
)

func main() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "channel 1 data"
		time.Sleep(time.Second * 2)
	}()

	go func() {
		//time.Sleep(time.Second * 5)
		anotherChannel <- "channel 2 data"

	}()

	select {
	case msgFromMyChannel := <-myChannel:
		fmt.Println(msgFromMyChannel)
		fmt.Println("got data from channel 1")
	case msgFromAnotherChannel := <-anotherChannel:
		fmt.Println(msgFromAnotherChannel)
		fmt.Println("data from channel 2")
	}

}
