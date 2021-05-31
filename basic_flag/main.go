package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {

	functionName := flag.String("function", "", "function_name")
	name := flag.String("name", "function not found", "function_name")
	round := flag.Int("round", 1, "round loop")

	flag.Parse()

	fmt.Println("functionName:", *functionName)
	//fmt.Println("Name:", *Name)
	//fmt.Println("Round:", *Round)

	for i:=0;i<*round;i++{
		if *functionName != ""{
			fmt.Printf("%s | %s | %s | round %d", time.Now().Format("01/02/2006 15:04:05"), *functionName, *name, i+1)
			fmt.Println()
			time.Sleep(1 * time.Second)
		}else{
			fmt.Printf("%s | %s", time.Now().Format("01/02/2006 15:04:05"), *name)
		}
	}

}