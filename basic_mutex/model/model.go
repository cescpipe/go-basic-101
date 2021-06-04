package model

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type Animal struct {
	Name     string
	RunMeter int
	mu       sync.Mutex
}

func (a *Animal) RunByRoutines(routineIndex string, wg *sync.WaitGroup) {

	fmt.Println("Increment Meter By : ", routineIndex)
	a.mu.Lock()
	a.RunMeter += 1
	wg.Done()
}

func (a *Animal) RunByRoutinesWithoutMutex(routineIndex string, wg *sync.WaitGroup) {
	// uncomment to see race condition
	rti, _ := strconv.Atoi(routineIndex)
	fmt.Println("Increment Meter By : ", routineIndex)

	if rti%2 == 0 {
		time.Sleep(5 * time.Second)
		a.RunMeter += 1
	} else {
		fmt.Println("Increment Meter By : ", routineIndex)
		time.Sleep(1 * time.Second)
		a.RunMeter += 1
	}

	wg.Done()
}
