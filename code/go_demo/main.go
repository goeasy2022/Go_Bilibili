package main

import (
	"fmt"

	"github.com/goeasy2022/go_demo/sum"
)

func main() {
	fmt.Println("Hello World!")
	var a int = 1
	var b int = 2
	var c int = sum.Add(a, b)
	fmt.Println(c)

}
