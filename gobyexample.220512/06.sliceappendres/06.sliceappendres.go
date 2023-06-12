package main

import (
	"fmt"
)

func showAppendRes() {
	s := []int{5}    // [5] len=1 cap=1
	s = append(s, 7) // [5,7] len=2, cap=2
	s = append(s, 9) // [5,7,9] len=3, cap=4

	x := append(s, 11)
	fmt.Println(s, x) // [5 7 9] [5 7 9 11]

	y := append(s, 12)
	fmt.Println(s, x, y) // [5 7 9] [5 7 9 12] [5 7 9 12]

	s = append(s, 13)
	fmt.Println(s, x, y) // [5 7 9 13] [5 7 9 13] [5 7 9 13]
}

func showAppendPointer() {
	s := []int{5}
	fmt.Printf("%p\n", s) // 0x1400012c008

	s = append(s, 7)
	fmt.Printf("%p\n", s) // 0x1400012c020

	s = append(s, 9)
	fmt.Printf("%p\n", s) // 0x14000130000

	x := append(s, 11)
	fmt.Printf("%p, %p\n", s, x)         // 0x14000130000, 0x14000130000 // 只要发生扩容，底层就会变化；不扩容，底层就不变化，会相互影响
	fmt.Printf("%p, %p\n", &s[0], &x[0]) // 0x14000130000, 0x14000130000

	y := append(s, 12)
	fmt.Printf("%p, %p\n", s, y)         // 0x14000130000, 0x14000130000 // 只要发生扩容，底层就会变化；不扩容，底层就不变化，会相互影响
	fmt.Printf("%p, %p\n", &s[0], &y[0]) // 0x14000130000, 0x14000130000

	s = append(s, 13)
	fmt.Printf("%p\n", s) // 0x14000130000
}

func showAppendPointer2() {
	s := []int{5}
	fmt.Printf("%p\n", &s) // 0x1400009c018

	s = append(s, 7)
	fmt.Printf("%p\n", &s) // 0x1400009c018

	s = append(s, 9)
	fmt.Printf("%p\n", &s) // 0x1400009c018

	x := append(s, 11)
	fmt.Printf("%p, %p\n", &s, &x) // 0x1400009c018, 0x1400009c030 // 只要发生扩容，底层就会变化；不扩容，底层就不变化，会相互影响

	y := append(s, 12)
	fmt.Printf("%p, %p\n", &s, &y) // 0x1400009c018, 0x1400009c048 // 只要发生扩容，底层就会变化；不扩容，底层就不变化，会相互影响

	s = append(s, 13)
	fmt.Printf("%p\n", &s) // 0x1400009c018
}

func main() {
	// showAppendPointer()

	var b []string
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", b)
	b = append(b, "lihao")
	fmt.Printf("%p\n", &b)
	fmt.Printf("%p\n", b)

	// print(b == nil)
	// a := new([]string)
	// print(a == nil)
	// fmt.Printf("%v\n", a)
	// fmt.Printf("%p\n", a)
}
