package main

import "fmt"

func main() {
	const value = 10
	var i int = value
	var f float32 = float32(value)
	fmt.Println("i",i)
	fmt.Println("f",f)
}