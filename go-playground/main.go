package main

import (
	"fmt"

	dog "vv/dog"
)

func main() {
	// nChan := new(chan int)

	fmt.Println("new")

	// testDog()

	var a, b int = 1, 3
	fmt.Println(a, b)

	variableValue := 12
	fmt.Print(variableValue)
}

func testDog() {
	d := new(dog.Dog)
	d.Age = 19
	d.Name = "wukong"
	d.Run()
	d.Eat()

	fmt.Println(d)
}
