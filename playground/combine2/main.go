package main

import "fmt"

type eater interface {
	eat()
}

type People struct {
	食物 string
}

func (p People) eat() {
	fmt.Printf("i eat %s every day.\n", p.食物)
}

type Cat struct {
	零食 string
}

func (c Cat) eat() {
	fmt.Printf("i like %s.\n", c.零食)
}

func main() {
	people := People{食物: "米饭"}
	showInfo(people)

	cat := Cat{零食: "🐟"}
	showInfo(cat)
}

func showInfo(e eater) {
	e.eat()
}
