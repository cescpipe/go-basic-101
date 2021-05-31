package main

import (
	"fmt"
	"go-tutorial/basic_interface/model"
)

type Bear struct {
	Name string
	Color string
}

func (b *Bear) GetName() string{
	return b.Name
}

func (b *Bear) GetColor() string{
	return b.Color
}

type Dog struct {
	Name string
	Color string
}

func (d *Dog) GetName() string{
	return d.Name
}

func (d *Dog) GetColor() string{
	return d.Color
}


func main(){

	wb := Bear{
		Name:  "Polar Bear",
		Color: "white",
	}

	dbm := Dog{
		Name:  "Doberman",
		Color: "black",
	}


	fmt.Println("All Animals")
	var animals []model.Animal

	animals = append(animals, &wb, &dbm)

	for _, v := range animals{
		ta := v.(model.Animal)
		fmt.Printf("%s/%s \n", ta.GetName(), ta.GetColor())

	}

}
