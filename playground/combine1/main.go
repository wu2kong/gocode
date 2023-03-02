package main

import "fmt"

type PersonInterface interface {
	SayHi()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHi() {
	fmt.Printf("My name is %s, my age is %d\n", p.Name, p.Age)
}

func showInfo(p PersonInterface) {
	p.SayHi()
}

type Student struct {
	Person
}

type Teacher struct {
	Person
}

func main() {
	stu := Student{Person{"stu.wu2kong", 18}}
	showInfo(stu)

	tea := Teacher{Person{"teacher.ABC", 28}}
	showInfo(tea)
}
