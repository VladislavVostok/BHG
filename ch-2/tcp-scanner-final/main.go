package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, results chan int) {
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func main() {
	ports := make(chan int, 100) // Создается канал с помощью make(). В качестве второго параметра передается значение 100.
	// Это добавляет каналу буферизацию, то есть в него можно будет отправлять элемент и не ждать,
	// пока получатель этот элемент прочтет.
	// Буферизованные каналы идеально подходят для поддержания и отслеживания работы нескольких
	// производителей и потребителей. Емкость канала определяется как 100.
	// Значит, он может вместить 100 элементов, до того как отправитель будет заблокирован.
	// Это дает небольшой прирост производительности, поскольку все воркеры смогут запускаться сразу.

	results := make(chan int)
	var openports []int

	for i := 0; i <= cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 1; i <= 1024; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)

	sort.Ints(openports)

	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}

}
