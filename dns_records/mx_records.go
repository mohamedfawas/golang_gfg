package main

import (
	"fmt"
	"net"
)

func main() {
	mx_recs, _ := net.LookupMX("facebook.com")
	for _, v := range mx_recs {
		fmt.Println(v)
	}

}
