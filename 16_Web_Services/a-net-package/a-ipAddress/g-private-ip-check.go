package main

import (
	"fmt"
	"net"
)

func isPrivateIP(ip net.IP) bool {
	privateNets := []string{
		"10.0.0.0/8",
		"172.16.0.0/12",
		"192.168.0.0/16",
	}
	for _, cidr := range privateNets {
		_, net, _ := net.ParseCIDR(cidr)
		if net.Contains(ip) {
			return true
		}
	}
	return false
}

func main() {
	ip := net.ParseIP("192.168.1.1")
	fmt.Println(isPrivateIP(ip)) // true
}