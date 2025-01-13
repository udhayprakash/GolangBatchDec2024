package main

import (
	"fmt"
	"net"
	"reflect"
)

func main() {
	// ipObj:= net.ParseIP("192.0.2.1")
	// if ipObj != nil {
	// 	fmt.Println("Parsed IP:", ipObj)
	// } else {
	// 	fmt.Println("Invalid IP address")
	// }
	// fmt.Println()

	for _, ipAddr := range []string{
		"192.0.2.2",
		"292.0.2.2",
		"192.0.2.2456",
		"192.0.2.245",
		"142.250.195.238",
		"0.0.0.0",
		"255.255.255.255",
		"127.0.0.1",
	} {
		// String to ip address parsing
		ip := net.ParseIP(ipAddr)
		if ip != nil {
			fmt.Println("Ip address: ", ip)
			fmt.Println("Type      : ", reflect.TypeOf(ip))
			fmt.Println("DefaultMask:", ip.DefaultMask())
			fmt.Println("isLoopBack: ", ip.IsLoopback())

		} else {
			fmt.Println("Invalid IP address:", ipAddr)
		}
		fmt.Println()
	}
}
