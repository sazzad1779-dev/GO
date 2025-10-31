package main

import "fmt"

func main() {
	var a int = 10
	var b float32 = 11.1
	var c int = a + int(b)
	var d float32 = float32(a) + b
	fmt.Println("a: ", a, "\nb: ", b, "\nint total :", c, "\nfloat32 total", d)
}
