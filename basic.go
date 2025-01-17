package main

import "fmt"

func main() {
	fmt.Println("Hello, GoPhers!")
	x := 10
	//x = 20
	fmt.Println(x)
	// x = 30
	//var i int = 20
	//var f float32 = float32(i)
	//fmt.Println("i = ", i)
	//fmt.Println("f = ", f)

	const value = 10
	var i int = value
	var f float32 = value
	fmt.Println("i = ", i)
	fmt.Println("f = ", f)

}
