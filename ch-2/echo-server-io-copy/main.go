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

	// Копируем данные из io.Reader в io.Writer через io.Copy()
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("Unable to read/write data")
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
