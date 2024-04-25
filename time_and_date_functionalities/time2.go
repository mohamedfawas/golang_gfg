package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("checking time taken for executing this program")
	end := time.Now()
	fmt.Println("time taken is : ", end.Sub(start))
}
