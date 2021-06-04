package main

import (
	"fmt"
	"go-tutorial/basic_mutex/model"
	"sync"
)

func main() {

	gopher := model.Animal{
		Name:     "Gopher",
		RunMeter: 0,
	}

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		rt := fmt.Sprintf("%d", i+1)
		//go gopher.RunByRoutines(rt, &wg)
		go gopher.RunByRoutinesWithoutMutex(rt, &wg)

	}

	wg.Wait()
	fmt.Println("total gopher run meter : ", gopher.RunMeter)
}
