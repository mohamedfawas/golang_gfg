package main

import (
	"fmt"
	"net"
)

func main() {
	iprecords, _ := net.LookupIP("www.facebook.com")
	fmt.Println(iprecords)
	for _, ip := range iprecords {
		fmt.Println(ip)
	}

	//canonical name
	cname, _ := net.LookupCNAME("www.facebook.com")
	fmt.Println(cname)
}
