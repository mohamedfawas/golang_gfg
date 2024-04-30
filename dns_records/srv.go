package main

import (
	"fmt"
	"net"
)

func main() {
	cname, srvs, err := net.LookupSRV("xmpp-server", "tcp", "golang.org")
	fmt.Printf("\ncname: %s \n\n", cname)
	if err != nil {
		panic(err)
	}
	for _, srv := range srvs {
		fmt.Println(srv)
	}

}
