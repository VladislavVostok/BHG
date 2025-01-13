package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup //	выступающую в качестве синхронизированного счетчика
	for i := 1; i <= 1024; i++ {
		wg.Add(1) // увеличиваем этот счетчик при каждом создании горутины для сканирования порта
		go func(j int) {
			defer wg.Done() // уменьшает этот счетчик при завершении каждой единицы работы
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait() // который блокирует выполнение, пока не будет выполнена вся работа и счетчик не достигнет нуля
}
