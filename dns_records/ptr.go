package main

import (
	"fmt"
	"net"
)

func main() {
	//fmt.Println(net.LookupAddr("6.8.8.8"))
	dom_name, _ := net.LookupAddr("6.8.8.8")
	for _, v := range dom_name {
		fmt.Println(v)
	}
}
