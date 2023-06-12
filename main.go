package main

import "fmt"

type Stu struct {
	Name string
}

func main() {
	// s1 := make([]Stu, 2)
	// println(s1)
	// s1 = append(s1, Stu{Name: ""})
	s1 := Stu{}
	fmt.Printf("%T\n", s1)

	s2 := new(Stu)
	println(s2)
	fmt.Printf("%T\n", s2)
	fmt.Printf("%v\n", s2)
}
