// fns, a take 2, b t 3s,
// a assigns val to channel , and b prints the value
package main

import (
	"fmt"
	"time"
)

func a(myChannel chan string) {
	//time.Sleep(time.Second * 2)
	//fmt.Println("from fun A")
	myChannel <- "from fun A"
}

func b(myChannel chan string) {
	//time.Sleep(time.Second * 3)
	//fmt.Println("from fun B")
	out := <-myChannel
	fmt.Println(out)
}

func main() {
	myChannel := make(chan string)
	go a(myChannel)
	go b(myChannel)
	time.Sleep(time.Second * 3)

	// start := time.Now()
	// go a()
	// go b()

	// end := time.Now()
	//end.Sub(start)
	//time.Now().Sub(end.)
	//time_taken := end - start
	// fmt.Println(time.Since(start))

	// fmt.Println(end.Sub(start))
}
