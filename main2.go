package main

import (
	"fmt"
	"io"
	"strings"
)

type Rectange struct {
	witdh  int
	height int
}

/**
 * Structs implement these methods are automatically became
 * a Move type.
 */

type Move interface {
	painting(int) *Rectange
	drawing(int) *Rectange
}

func (v *Rectange) painting(scale int) *Rectange {
	fmt.Println(v, scale)
	v.height = 900
	return v
}

func (v *Rectange) drawing(size int) *Rectange {
	v.witdh = 30
	fmt.Println("Drawing..", v, size)
	return v
}

/**
 * Generic Type
 */
func typeSwitch(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("float", "fuck...")
	}
}

/**
 * Implementing Stringer interface
 */
func (v *Rectange) String() string {
	return fmt.Sprintf("~%v %v~", v.witdh, v.height)
}

/**
 * IP Address
 */
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

/**
 * Returning customize error
 */
type MyError struct {
	name  string
	times int
}

func (e *MyError) Error() string {
	return "FUCK you up!"
}

func runWithError() error {
	var e error // implemented Error interface.
	e = &MyError{"Jerry", 23}
	return e
}

func ip() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

/* Default an error type with float value */
type ErrSqrNeg float64

func (e ErrSqrNeg) Error() string {
	return fmt.Sprintf("%v fuck it!", float64(e))
}

func sqr(num float64) (float64, error) {
	if num < 0 {
		return num, ErrSqrNeg(num)
	} else {
		return num, nil
	}
}

/**
 * Play with go routine
 */
func printName(name string) {
	fmt.Println(name)
	fmt.Println(name)
	fmt.Println(name)
	if name == "Jerry" {
		fmt.Println(name)
		fmt.Println(name)
		fmt.Println(name)
		fmt.Println(name)
		fmt.Println(name)
	}
}

func main() {
	fmt.Println(runWithError())

	var v Move = &Rectange{2, 3}
	v.painting(3).drawing(20).painting(300)
	fmt.Println("fuck", v)

	// empty interface can refer to any types
	var i interface{}
	i = 3
	i = 3.232
	i = "fuck"
	fmt.Println(i)

	// type assertion
	t := i.(string)
	fmt.Println(t, "...")
	typeSwitch(0)
	typeSwitch(0.1)
	typeSwitch("0.1")

	v2 := &Rectange{12, 34}
	fmt.Println(v2)
	ip()

	_, err := sqr(-3)
	// fmt.Println(ret, float64(err))
	if err != nil {
		fmt.Println("YES>", err)
	}

	b := []byte{33, 44, 4}
	fmt.Println(b)
	go printName("Jerry")
	printName("Lucy")
	printName("Tomy")
	read8bytes()
}

func read8bytes() {
	r := strings.NewReader("Hello, Reader........")
	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Println(n, err, b, b[:n])
		if err == io.EOF {
			fmt.Println("fuck")
			break
		}
	}
}
