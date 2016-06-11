package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	ONE = iota
	TWO
	THREE
)

func main() {
	fmt.Print(ONE, TWO, THREE)
	fmt.Print("fuc")
	c := make(chan int)
	for i := 0; i < 10; i++ {
		go process(i, c)
	}
	for {
		select {
		case c <- rand.Int():
			fmt.Println("Now.,")
		case i := <-c:
			fmt.Println("Now.", i)
		default:
			fmt.Println("default", len(c))
		}
		time.Sleep(time.Millisecond * 100 * 1)
	}
}

func process(t int, c chan int) {
	c <- t
}
