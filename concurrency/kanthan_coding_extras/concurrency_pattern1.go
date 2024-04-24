package main

import "fmt"

func main() {
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	//loop over that slice of characters and for each character we are going to put it on to our char channel.
	for _, s := range chars {
		select {
		case charChannel <- s:

		}
	}
	close(charChannel) //after putting all the data we close our channel

	// loop over the channel and recieve the data that  is put into the channle
	for result := range charChannel {
		fmt.Println(result)
	}
}
