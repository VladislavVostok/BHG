package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Dog struct{}

func (d *Dog) SayHello() {
	fmt.Println("Woof woof")
}

type Friend interface {
	SayHello()
}

func Greet(f Friend) {
	f.SayHello()
}

/*
В структурах отсутствуют модификаторы области, такие как закрытая (private),
публичная (public) или защищенная (protected), которые обычно присутствуют
в других языках и служат для управления доступом к их членам.
Вместо этого в Go область доступности определяется величиной регистра: типы и поля,
начинающиеся с `прописной` буквы, экспортируются и являются доступными вне пакета,
в то время как начинающиеся со `строчной` — закрытые и доступны только внутри пакета.
*/
type Person struct {
	Name string
	Age  int
}

func (p *Person) SayHello() {
	fmt.Println("Hello,", p.Name)
}

func main() {
	fmt.Println("Hello, Black Hat Gophers!")

	//Примитивные типы данных

	/*
		bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, byte, rune, float32, float64, complex64 и complex128
	*/

	var x = "Htllo Wrld"
	z := int(42)

	fmt.Println(x)
	fmt.Println(z)

	// Срезы и карты
	var s = make([]string, 0)
	var m = make(map[string]string)
	s = append(s, "some string")
	m["some key"] = "some value"
	fmt.Println(s)
	fmt.Println(m)

	//Указатели, структуры и интерфейсы
	ptr := &z
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr = 100
	fmt.Println(z)

	var guy = new(Person)
	guy.Name = "Dave"
	guy.SayHello()

	Greet(guy)

	var dog = new(Dog)
	Greet(dog)

	y := int(1)
	if y == 1 {
		fmt.Println("X is equal to 1")
	} else {
		fmt.Println("X is not equal to 1")
	}

	switch x {
	case "foo":
		fmt.Println("Found foo")
	case "bar":
		fmt.Println("Found bar")
	default:
		fmt.Println("Default case")
	}

	other(42)          // Вывод: I'm an integer!
	other("hello")     // Вывод: I'm a string!
	other(3.14)        // Вывод: Unknown type!
	other([]int{1, 2}) // Вывод: Unknown type!

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	nums := []int{2, 4, 6, 8}
	for idx, val := range nums {
		fmt.Println(idx, val)
	}

	// Многопоточность
	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("main function")

	// Каналы
	c := make(chan int)
	go strlen("Salutatuins", c)
	go strlen("World", c)
	a, b := <-c, <-c
	fmt.Println(a, b, a+b)

}

//type switch

func other(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("I'm an integer!")
	case string:
		fmt.Println("I'm a string!")
	default:
		fmt.Println("Unknown type!", v)
	}

	// Обработка ошибок
	if err := foo1(); err != nil {
		fmt.Println("Здесь обработка ошибки!")
	}

	f := Foo{"Joe Junior", "Hello Shabado"}
	b, _ := json.Marshal(f)
	fmt.Println(string(b))
	json.Unmarshal(b, &f)

}

// Горутины
func f() {
	fmt.Println("f function")
}

// Каналы
func strlen(s string, c chan int) {
	c <- len(s)
}

// Обработка ошибок
type MyError string

func (e MyError) Error() string {
	return string(e)
}

func foo1() error {
	return errors.New("Some Error Occurred")
}

type Foo struct {
	Bar string
	Baz string
}
