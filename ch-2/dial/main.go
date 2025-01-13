package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanmy.nmap.org:80")

	if err == nil {
		fmt.Println("Connections successful")
	}
}
