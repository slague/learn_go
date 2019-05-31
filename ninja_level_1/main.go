package main

import "fmt"

// EX 1

func exercise1() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

// EX 2

var a int
var b string
var c bool

func exercise2() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// EX 3

var num = 42
var name = "James Bond"
var isItTrue = true

func exercise3() {
	s := fmt.Sprintf("%v %v %v", num, name, isItTrue)
	s2 := fmt.Sprintf("%v\t%v\t%v", num, name, isItTrue) // with \n
	fmt.Println(s)
	fmt.Println(s2)
}

// EX 4

func exercise4() {
	type pizza int
	var p pizza
	fmt.Println(p)
	fmt.Printf(`Value of "p" is: %v Type of "p" is: %T`, p, p)
	fmt.Println()
}

// EX 5

type hotdog int

var h hotdog
var y int

func exercise5() {
	fmt.Printf(`Value of "y" is: %v Type of "y" is: %T`, y, y)
	fmt.Println()
	fmt.Printf(`Value of "h" is: %v Type of "h" is: %T`, h, h)
	fmt.Println()
	y = int(h)
	fmt.Printf(`NOW the value of "y" is: %v Type of "y" is: %T`, y, y)
	fmt.Println()
}

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
}
