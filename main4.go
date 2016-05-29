package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)
	// Channel is like a stack
	go sum(100, 200, c)
	go sum(1, 3, c)
	go sum(1, 34, c)
	x, y := <-c, <-c
	fmt.Println(x, y, <-c)
	fmt.Println(runtime.NumCPU(), "fuck")
}

func sum(x, y int, c chan int) {
	c <- x + y
}
