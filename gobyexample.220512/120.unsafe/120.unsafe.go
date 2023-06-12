package main

import (
	"fmt"
	"unsafe"
)

type Programmer struct {
	name     string
	age      int
	language string
}

func main() {
	// p := Programmer{"stefno11", "go"}
	// fmt.Println(p)

	// name := (*string)(unsafe.Pointer(&p))
	// *name = "qcrao11"

	// lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Offsetof(p.language)))
	// *lang = "Golang"

	// fmt.Println(p)

	// fmt.Printf("%T, %v\n", unsafe.Pointer(&p), unsafe.Pointer(&p))
	// fmt.Printf("%T, %v\n", uintptr(unsafe.Pointer(&p)), uintptr(unsafe.Pointer(&p)))
	// fmt.Printf("%T, %v\n", unsafe.Offsetof(p.language), unsafe.Offsetof(p.language))
}

func main1() {
	p := Programmer{"stefno", 18, "go"}
	fmt.Println(p)

	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&p)) + unsafe.Sizeof(int(0)) + unsafe.Sizeof(string(""))))
	*lang = "Golang"

	fmt.Println(p)
}
