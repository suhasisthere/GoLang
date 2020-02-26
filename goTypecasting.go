package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 10
	b := strconv.Itoa(a)
	fmt.Println("felli")

	fmt.Println("value of b is: " + b)
	fmt.Printf("type of B is %T\n", b)

	for i := 0;i<10;i++ {
		fmt.Println(i)
	}

}
