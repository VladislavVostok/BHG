package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		fmt.Println(address)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// порт закрыт или отфильтрован
			continue
		}

		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
