package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print("fuc")
	c := make(chan int, 1)
	for i := 0; i < 10; i++ {
		go process(i, c)
	}
	for {
		select {
		case c <- rand.Int():
			fmt.Println("Now.,")
		case <-c:
			fmt.Println("Now.")
		default:
			fmt.Println("default", len(c))
		}
		time.Sleep(time.Millisecond * 100 * 1)
	}
}

func process(t int, c chan int) {
	c <- t
}
