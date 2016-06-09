package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	fmt.Print(os.Getenv("NODE_DEV"))
	a := []int{}
	var b [10]int
	fmt.Println(len(a), cap(a), reflect.TypeOf(a), reflect.TypeOf(b))

	a = append(a, 1)
	fmt.Println(cap(a))

	a = append(a, 1)
	fmt.Println(cap(a))

	a = append(a, 1)
	fmt.Println(cap(a))

	a = append(a, 1)
	a = append(a, 1)
	fmt.Println(cap(a))

	c := make(chan interface{})
	go func() {
		c <- 1
	}()
	go func() {
		c <- "string"
	}()
	fmt.Print(<-c, <-c)
}
