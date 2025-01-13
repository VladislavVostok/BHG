package main

import (
	"io"
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
	b := make([]byte, 512)

	for {
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected error")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))

		// Отправляем данные через conn.Write
		log.Println("Writing data")

		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to write data")
		}
	}
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
