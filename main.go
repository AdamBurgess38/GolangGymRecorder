package main

import (
	"fmt"
	"packages/inputters"
)

func main() {
    fmt.Println("hello world")
    test := inputters.FetchInteger("hello wolrd!",2)
    fmt.Println(test)
}