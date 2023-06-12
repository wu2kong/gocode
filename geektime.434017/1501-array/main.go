package main

import "fmt"

func main() {
	var arr = [...]int{1, 2, 3}
	fmt.Printf("%T\n", arr)

	var sli = []int{1, 2, 3}
	fmt.Printf("%T\n", sli)

	arr2 := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sl := arr2[3:7:9]
	fmt.Println(sl)
	var sli2 []int
	var sli3 = []int{}
	fmt.Println(sli2, len(sli2), cap(sli2))
	fmt.Println(sli3, len(sli3), cap(sli3))
	fmt.Println(sli2 == nil)
	fmt.Println(sli3 == nil)
	fmt.Printf("%p\n", sli2)
	fmt.Printf("%p\n", sli3)

	u := [...]int{11, 12, 13, 14, 15}
	fmt.Println("array:", u) // [11, 12, 13, 14, 15]
	s := u[1:3]
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s) // [12, 13]
	s = append(s, 24)
	fmt.Println("after append 24, array:", u)
	fmt.Printf("after append 24, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
	s = append(s, 25)
	fmt.Println("after append 25, array:", u)
	fmt.Printf("after append 25, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
	s = append(s, 26)
	fmt.Println("after append 26, array:", u)
	fmt.Printf("after append 26, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)

	s[0] = 22
	fmt.Println("after reassign 1st elem of slice, array:", u)
	fmt.Printf("after reassign 1st elem of slice, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
}
