package main

import (
	"fmt"
	"time"
)

func printRoutine(rt string, dc chan string) {
	for i := 0; i < 20; i++ {
		time.Sleep(2 * time.Second)
		tn := time.Now().Format("2006/01/02 15:04:05")
		st := fmt.Sprintf("%s_%s_%d", tn, rt, i+1)
		dc <- st
	}
}

func main() {

	dc := make(chan string)
	defer close(dc)

	go func() {
		for v := range dc {
			fmt.Println(v)
		}
	}()

	go printRoutine("R1", dc)
	time.Sleep(2 * time.Second)
	go printRoutine("R2", dc)
	time.Sleep(3 * time.Second)
	go printRoutine("R3", dc)

	time.Sleep(60 * time.Second)
}
