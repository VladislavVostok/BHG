package main

import (
	"bufio"
	"log"
	"net"
)

//echo - это функция-обработчик, просто отражающая полученные данные

func echo(conn net.Conn) {
	/*
	   Defer — это ключевое слово в Go, которое позволяет
	   отложить выполнение функции до момента завершения выполнения текущей функции.
	   Это относительно простой способ управлять ресурсами.
	*/

	defer conn.Close()

	//создаём буфер для хранения полученных данных
	reader := bufio.NewReader(conn)

	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	log.Printf("Read %d bytes: %s", len(s), s)

	// Отправляем данные через conn.Write
	log.Println("Writing data")

	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}
	writer.Flush()
}

func main() {
	// Привязываемся к TCP-порту 20080 во всех интерфейсах
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		// Ожидаем соединения и при его установке создаём net.Conn
		conn, err := listener.Accept()
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accept connection")
		}
		// обрабатываем соединение, используя горутины для многопоточности

		go echo(conn)
	}
}
