package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(net.JoinHostPort("127.0.0.1", "6666"))
	fmt.Println(net.JoinHostPort("::1", "6666"))
	fmt.Println(net.SplitHostPort("127.0.0.1:9999"))
	fmt.Println(net.SplitHostPort("[::1]:9999"))
	fmt.Println(net.LookupAddr("127.0.0.1"))
}
