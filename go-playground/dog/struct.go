package dog

import "fmt"

type Animal struct {
	Age int
}

type Dog struct {
	Animal
	ID   int
	Name string
}

func (d *Dog) Run() {
	fmt.Println("i am a dog: run ", d.Name)
}

func (d *Animal) Eat() {
	fmt.Println("i am a animal: eat ", d.Age)
}
