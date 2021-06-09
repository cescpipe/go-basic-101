package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	number := 100.3875

	fmt.Println("start fmt package")
	fmt.Println(fmt.Sprintf("Hello %s", "World"))
	fmt.Println(fmt.Sprintf("%.2f", number))

	name := struct {
		Name string
		Age  int
	}{
		Name: "Gopher",
		Age:  6,
	}

	fmt.Println(fmt.Sprintf("%v", name))
	fmt.Println(fmt.Sprintf("%+v", name))
	fmt.Println(fmt.Sprintf("%#v", name))

	fmt.Println(fmt.Sprintf("%010d", 9))
	fmt.Println(fmt.Sprintf("%010d", 170))
	fmt.Println(fmt.Sprintf("%010d", 2000))
	fmt.Println("end fmt package")
	fmt.Println("")

	fmt.Println("start math package")
	fmt.Println(fmt.Sprintf("default rand 1: %d", rand.Intn(1000)))
	fmt.Println(fmt.Sprintf("default rand 2: %d", rand.Intn(1000)))
	fmt.Println(fmt.Sprintf("default rand 1: %f", rand.Float64()))
	fmt.Println(fmt.Sprintf("default rand 2: %f", rand.Float64()))

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Println(fmt.Sprintf("custom seed rand: %d", r1.Intn(1000)))
	fmt.Println(fmt.Sprintf("customer seed rand: %f", r1.Float64()))
	fmt.Println("end math package")
	fmt.Println("")

	fmt.Println("start strconv package")
	v := "3.1415926535"
	if s, err := strconv.ParseFloat(v, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	if s, err := strconv.ParseFloat("NaN", 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

}
