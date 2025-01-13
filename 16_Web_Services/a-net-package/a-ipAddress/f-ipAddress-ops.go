package main

import (
	"fmt"
	"net"
)

func main() {
	ip1 := net.ParseIP("192.168.1.1")
	ip2 := net.ParseIP("192.168.1.1")
	ip3 := net.ParseIP("192.168.1.0")

	fmt.Println("Equality Check")
	// fmt.Println(ip1 == ip2)
	fmt.Println(ip1.Equal(ip2)) // true
	fmt.Println(ip1.Equal(ip3)) // false

	fmt.Println("\nComparing ipv4 and ipv6")
	ipv4 := net.ParseIP("192.168.1.1")
	ipv6 := net.ParseIP("::1")

	fmt.Println(ipv4.To4() != nil)      // true
	fmt.Println(ipv6.To4() == nil)      // true
	fmt.Println(ipv4.Equal(ipv6.To4())) // false

	fmt.Println("Inequality Check")
	fmt.Println(!ip1.Equal(ip2)) // true

	fmt.Println("\nLess Than / Greater Than Checks")
	fmt.Println(ip1.String() < ip2.String()) // false
	fmt.Println(ip1.String() > ip2.String()) // false

	fmt.Println("Network and Host Address Checks")
	// 192.168.0.0 is ending in zero
	fmt.Println(isNetworkAddress(ip1)) // true
	fmt.Println(isNetworkAddress(ip3)) // false

	fmt.Println("CIDR check")
	_, cidr, _ := net.ParseCIDR("192.168.1.0/24")
	ip := net.ParseIP("192.168.1.10")
	fmt.Println(cidr.Contains(ip)) // true

	fmt.Println("Subnet comparision")
	_, parentCIDR, _ := net.ParseCIDR("192.168.1.0/24")
	_, childCIDR, _ := net.ParseCIDR("192.168.1.128/25")

	fmt.Println(parentCIDR.Contains(childCIDR.IP)) // true
	
}

func isNetworkAddress(ip net.IP) bool {
	return ip != nil && ip.To4() != nil && ip[len(ip)-1] == 0
}
