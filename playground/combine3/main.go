package main

import "fmt"

type Personer interface {
	SayHi()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) SayHi() {
	fmt.Printf("say from person, name = %s, age = %d\n", p.Name, p.Age)
}

type Student struct {
	Person
}

// func (p Student) SayHi() {
// 	fmt.Printf("say from student, name = %s, age = %d\n", p.Name, p.Age)
// }

func main() {
	stu := Student{Person{"stu.wu2kong", 18}}
	showInfo(stu)
}

func showInfo(p Personer) {
	p.SayHi()
}
