package main

import (
	"fmt"
	"net"
)

func main() {
	ns_recs, _ := net.LookupNS("facebook.com")
	//fmt.Println(ns_recs)
	for _, v := range ns_recs {
		fmt.Println(v)
	}

}
