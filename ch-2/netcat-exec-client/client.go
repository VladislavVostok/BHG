package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// Подключение к серверу
	serverAddress := "192.168.1.104:20080"
	conn, err := net.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Не удалось подключиться к серверу: %v\n", err)
	}
	defer conn.Close()

	fmt.Printf("Подключение к серверу %s установлено.\n", serverAddress)
	fmt.Println("Введите команды для выполнения на сервере (exit для выхода):")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(">> ")
		// Читаем команду от пользователя
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Ошибка ввода: %v\n", err)
			break
		}

		// Отправляем команду серверу
		_, err = conn.Write([]byte(command))
		if err != nil {
			log.Printf("Ошибка отправки команды: %v\n", err)
			break
		}

		// Читаем ответ от сервера
		response := make([]byte, 4096)
		n, err := conn.Read(response)
		if err != nil {
			log.Printf("Ошибка получения ответа: %v\n", err)
			break
		}

		// Выводим результат
		fmt.Printf("Ответ сервера:\n%s\n", string(response[:n]))
	}

	fmt.Println("Клиент завершил работу.")
}
