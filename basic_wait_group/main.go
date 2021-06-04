package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	defer func() {
		fmt.Println("func main end")
	}()
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go worker(i, &wg)
		time.Sleep(1 * time.Second)
	}

	wg.Wait()
}
