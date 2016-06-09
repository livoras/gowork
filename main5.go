package main

import "fmt"

func pipe(name string, duration int, c chan string) {
	for i := 0; i < duration; i++ {
	}
	c <- name
}

type Person struct {
	name string
	age  int
}

func main() {
	nameChan := make(chan string)
	go pipe("Jerry", 100000, nameChan)
	go pipe("Tomy", 300, nameChan)
	fmt.Println("ucuk")
	fmt.Println(<-nameChan, <-nameChan)
}
