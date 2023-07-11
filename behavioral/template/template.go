/*
	The Template Pattern defines the skeleton of an algorithm in a method,
	deferring some steps to subclasses. Template Method lets subclasses
	redefine certain steops of an algorithm without changing the algorithm's
	structure
*/

package main

import (
	"fmt"
)

type Drink interface {
	Brew()
}

type Coffee struct{}

type Tea struct{}

func BoilWater() {
	fmt.Println("Boiling water...")
}

func Serve() {
	fmt.Println("Drink is ready!")
}

func (c *Coffee) Brew() {
	fmt.Println("grinding coffee bean...")
	fmt.Println("brewing coffee...")
}

func (t *Tea) Brew() {
	fmt.Println("setting up teapot...")
	fmt.Println("brewing tea... ")
}

func MakeDrink(d Drink) {
	BoilWater()
	if coffee, ok := d.(*Coffee); ok {
		coffee.Brew()
	} else {
		d.Brew()
	}
	Serve()
}

func main() {
	coffee := &Coffee{}
	tea := &Tea{}

	MakeDrink(coffee)
	MakeDrink(tea)
}
