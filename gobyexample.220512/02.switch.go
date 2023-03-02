package main

import (
	"fmt"
	"time"
)

// 值测试
func switchv1() {
	value := 4
	switch value {
	case 1:
		fmt.Println(value, " ")
	case 2:
		fmt.Println(value, " ")
	case 3:
		fmt.Println(value, " ")
	default:
		fmt.Println("default")
	}
}

// 多值测试
func switchv2() {
	switch time.Now().Weekday() {
	case time.Saturday:
		println("saturday")
	case time.Sunday:
		println("sun day")
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		println("week day")
	}
}

// 条件判断
func switchv3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		println("it is before noon.")
	case t.Hour() < 18:
		println("it is afternoon.")
	default:
		println("it is night.")
	}
}

// 类型判断，类型断言
func switchv4() {
	whatIAm := func(v interface{}) {
		switch v.(type) {
		case int:
			fmt.Printf("it is int value: %d\n", v)
		case bool:
			fmt.Printf("it is bool value: %v\n", v)
		default:
			fmt.Printf("i do not know this type value.\n")
		}
	}

	whatIAm(true)
	whatIAm(10)
	whatIAm(1.11)
}

func main() {
	switchv1()

	switchv2()

	switchv3()

	switchv4()
}
