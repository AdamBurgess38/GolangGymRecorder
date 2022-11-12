package main

import (
	"fmt"
	"packages/pkg"
)

func main() {
    fmt.Println("hello world")
    test := pkg.FetchInteger("hello wolrd!",2)
    fmt.Println(test)
}