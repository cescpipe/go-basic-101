package main

import (
	"fmt"
	"time"
)

func say1(s string) {
	for i := 0; i < 15; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}
}

func say2(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Second)
		fmt.Println(s)
	}
}

func main() {
	go say1("world")
	go say2("hello")
	time.Sleep(60 * time.Second)
}
