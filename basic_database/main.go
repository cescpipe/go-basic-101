package main

import (
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("start function SQL")
	connectPostgres()
	fmt.Println("end function SQL")
	fmt.Println()

	fmt.Println("start function noSQL")
	connectMongo()
	fmt.Println("end function noSQL")
}
