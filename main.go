package main

import "fmt"

type Dog struct {
	Name string
}

type Animal interface {
	speak() string
}

func getNewDog() *Dog {
	return &Dog{}
}

func (d *Dog) speak() string {
	return "Woof!"
}

type Cat struct {
	Name string
}

func getNewCat() *Cat {
	return &Cat{}
}

func (c *Cat) speak() string {
	return "Meow!"
}

type Lion struct {
	Name string
}

func getNewLion() *Lion {
	return &Lion{}
}

func (l *Lion) speak() string {
	return "Arrr!"
}

func main() {
	getAnimalFactory(1)
	getAnimalFactory(2)
	getAnimalFactory(3)
}

func getAnimalFactory(animalType int) {
	var animal Animal
	switch animalType {
	case 1:
		animal = getNewDog()
	case 2:
		animal = getNewCat()
	case 3:
		animal = getNewLion()
	}
	animalSound := animal.speak()
	fmt.Println(animalSound)
}
