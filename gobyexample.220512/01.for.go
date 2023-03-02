package main

import "fmt"

func main() {
	// 第一种循环
	forv1()

	// 第二种循环
	forv2()

	// 第三种循环
	forv3()
}

func forv1() {
	i := 1
	for i <= 3 {
		fmt.Print(i, " ")
		i++
	}

	fmt.Println()
}

func forv2() {
	for i := 9; i <= 12; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()
}

func forv3() {
	i := 1
	for {
		if i <= 3 {
			print("hello", " ")
		} else {
			break
		}
		i++
	}

	fmt.Println()
}
