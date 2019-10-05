package main

import (
	"fmt"
	"time"
)

func main() {
	var a = 100
	var b = 200
	var c = a + b

	fmt.Println(c)
	duration := time.Duration(10) * time.Second
	time.Sleep(duration)
}
