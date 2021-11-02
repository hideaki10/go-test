package main

import (
	"fmt"
	"net"
)

func main() {
	// fmt.Println(net.JoinHostPort("127.0.0.1", "6666"))
	// fmt.Println(net.JoinHostPort("::1", "6666"))
	// fmt.Println(net.SplitHostPort("127.0.0.1:9999"))
	// fmt.Println(net.SplitHostPort("[::1]:9999"))
	// fmt.Println(net.LookupAddr("127.0.0.1"))

	// host, err := net.LookupIP("www.baidu.com")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//	fmt.Println(host)

	//ip, ipnet, err := net.ParseCIDR("192.168.2.1/24")
	// fmt.Println(ip, ipnet, err)

	// fmt.Println(ipnet.Contains(net.ParseIP("192.168.2.1")))
	// fmt.Println(ipnet.Network())

	// addrs, _ := net.InterfaceAddrs()

	// for _, addr := range addrs {
	// 	fmt.Println(addr.Network(), addr.String())
	// }

	interfaces, _ := net.Interfaces()

	for _, inter := range interfaces {
		fmt.Println(inter.Name, inter.HardwareAddr, inter.MTU)

		addr, _ := inter.Addrs()
		for _, a := range addr {
			fmt.Println(a.Network(), a.String())
		}
	}
}
