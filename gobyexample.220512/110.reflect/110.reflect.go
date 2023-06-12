package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Student struct {
	Name string `json:"name" w2k:"a=2,b=3"`
	Age  int
}

func (s Student) SayHi() {
	fmt.Print(s.Name)
}
func main() {
	var i int = 0
	itype := reflect.TypeOf(i)
	fmt.Println(itype)

	stu := Student{"wu2kong", 18}
	var ii interface{}
	ii = stu
	stutype := reflect.TypeOf(ii)
	meth, ok := stutype.MethodByName("SayHi")
	if ok != true {
		fmt.Print(meth)
		os.Exit(-1)
	}

	fmt.Println(meth.Name)

	filed, _ := stutype.FieldByName("Age")
	fmt.Println(filed.Tag)

	nameFiled, _ := stutype.FieldByName("Name")
	fmt.Printf("Tag: %v\n", nameFiled.Tag)
	fmt.Printf("Tag: w2k: %v\n", strings.Split(nameFiled.Tag.Get("w2k"), ","))
}
