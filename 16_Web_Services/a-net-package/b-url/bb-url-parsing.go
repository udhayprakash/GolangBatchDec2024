package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	// db connection string
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme)          // postgres
	fmt.Println(u.User)            // user:pass
	fmt.Println(u.User.Username()) // user
	p, _ := u.User.Password()      // pass
	fmt.Println(p)

	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host) // host.com
	fmt.Println(port) // 5432

	fmt.Println(u.Path)     // /path
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
