package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go sum(100, 200, c)
	go sum(1, 3, c)
	x, y := <- c, <- c
	fmt.Println(x, y)
}

func sum(x, y int, c chan int) {
	c <- x + y
}
