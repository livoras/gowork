package main

import (
	"fmt"
	"reflect"
)

func main() {
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
}
