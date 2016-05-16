package main

import (
	"fmt"
	"math" // the package's name is the same as the last element of the path
	"strings"
)

// you can omit the type when conseutive parameters has the same type
func add(x, y int) int {
	return x + y
}

// of course you can have you function have multiplue returns
func addAndMul(x, y int) (int, int) {
	return x + y, x * y
}

// nacked arguments
func nacked(sum int) (x, y int) {
	var a, b int
	x = sum / 2
	y = sum / 3
	a = 30
	fmt.Println(a, b)
	return
}

// assignments between different variables need explict conversion
func typesMakeing() {
	var x int = 10
	var y float64 = 10
	var z uint = uint(y)
	fmt.Println(x, y, z, "...")
}

// constants
func circle(r float64) float64 {
	const PI = 3.14159
	return PI * r * r
}

func sayHiForNTimes(times int) {
	for i := 0; i < times; i++ {
		fmt.Println("Hello world")
	}
	// no while loop, for spells for it.
	i := times - 1
	for i < times {
		fmt.Println("Hello world n")
		i++
	}
}

// if and else
func forOdd(words string) {
	// l only lives in if scope
	if l := len(words) % 2; l != 0 {
		fmt.Println(words)
	} else if l := len(words); l%2 == 0 {
		fmt.Println("noop...")
	}
}

// switch statment
func swichStat(num int) {
	switch a := num + 1; a % 2 {
	case 0:
		fmt.Println("even")
	case 1:
		fmt.Println("odd")
	default:
		fmt.Println("nothing...")
	}
}

// switch for if and else equalent
func switchFor(num int) {
	switch {
	case num < 4:
		fmt.Println("less than 4")
	case num < 10:
		fmt.Println("less than 10")
	case num < 100:
		fmt.Println("less than  100")
	default:
		fmt.Println("Too much big..")
	}
}

// defer a function call, it actually evaluated immediately
// it actuall push function to a stack and pop it before return arguments.
func defersomething(num int) int {
	defer fmt.Println("say it", num)
	num = num % 2
	fmt.Println("say it now", num)
	return num
}

// Struct
type Retangle struct {
	width, height float64
}

func makeRet(width, height int) *Retangle {
	return &Retangle{float64(width), float64(height)}
}

// with array
func loop() {
	var arr []int // some magic here?
	arr = []int{0, 1, 2, 3, 4}
	arr = arr[1:3] // slices are just references to certain setions of an array
	// just checkout http://127.0.0.1:3999/moretypes/11 for `cap function`
	for i, len := 0, len(arr); i < len; i++ {
		fmt.Println(arr[i])
	}
}

// create a dynamic arry with make
func makeArrWithValue(num int, inside int) []int {
	arr := make([]int, num)
	for i := 0; i < num; i++ {
		arr[i] = inside
	}
	return arr
}

// using range to iterate an array
func iteration(arr []int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}

func WordCount(s string) map[string]int {
	m := map[string]int{}
	fileds := strings.Fields(s)
	for _, word := range fileds {
		count, ok := m[word]
		if ok {
			m[word] = count + 1
		} else {
			m[word] = 1
		}
	}
	return m
}

/* functions can be as normal values */
func functionAsValue() {
	clo := 3
	fn := func(a, b int) int {
		return a*b + clo
	}
	fmt.Println(fn(1, 2), "normall")
	set, get := closure(3)
	set(10000)
	fmt.Println(get(), "get it!")
}

/* try closure here..., and it actually works. */
func closure(x int) (func(int), func() int) {
	set := func(newX int) {
		x = newX
	}
	get := func() int {
		return x
	}
	return set, get

}

// with pointers
func doSomethingWithPointers() {
	i, j := 10, 20
	p := &i
	*p = *p + j
	fmt.Println(p, *p)
}

// pointer to struct
func pointToStruct() {
	rect := Retangle{100, 200}
	pRect := &rect
	fmt.Println(pRect.width == rect.width)
}

// slices: len and capacity
func playWithSlices() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1[:2]
	s3 := s1[2:]
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	fmt.Println(len(s3), cap(s3))
}

// append something to slice
func appendToSlice() {
	var s1 []int
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 2, 3, 4)
	fmt.Println(len(s1), cap(s1))

	s2 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(len(s2), cap(s2))

	s3 := s2[0:3]
	fmt.Println(len(s3), cap(s3), s3)

	// will cover the elements beneath
	s3 = append(s3, 20, 30, 40)
	fmt.Println(len(s3), cap(s3), s2, s3)

	// TODO: what's happended here?
	s3 = append(s3, 11, 22, 33, 44)
	fmt.Println(len(s3), cap(s3), s2, s3)

	fmt.Println(len(s2), cap(s2), s2, s3)
}

// making map
func doSomethingWithMap() {
	m := map[string]string{}
	m["name"] = "jerry"
	m["age"] = "13"

	m2 := map[string]string{
		"shit": "shit",
	}
	fmt.Println(m, m2)

	ele, ok := m["shit"]
	if ok {
		fmt.Println(ele)
	} else {
		fmt.Println("Fuck no shit.")
	}
}

func main() {
	// name begins with a capital letter is exported
	// otherwise is not exported.
	fmt.Println("Hello World\n", math.Pi, add(3, 4))
	// := is used to make a new variables decleration
	x, y := addAndMul(3, 4)
	fmt.Println("No the same", x, y)
	x, y = nacked(30)
	fmt.Println("Nacked ...", x, y)

	typesMakeing()
	fmt.Println(circle(30))

	sayHiForNTimes(5)
	forOdd("fuck you up.")

	swichStat(3)
	switchFor(101)

	fmt.Println(defersomething(10), "...")
	fmt.Println(makeRet(100, 200).width)

	loop()

	fmt.Println(makeArrWithValue(10, 3))
	fmt.Println(WordCount(" fuck you every day day by day"))
	functionAsValue()
	iteration([]int{1, 3, 4, 4})

	doSomethingWithPointers()
	pointToStruct()
	playWithSlices()
	appendToSlice()
	doSomethingWithMap()

	a := map[string]Retangle{
		"SHIT": {3, 4},
	}
	a["fuck"] = Retangle{2, 3}
	fmt.Println(a)
}
