package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 没有初始化的slice其值为nil
	var names []string
	fmt.Println(names)        // []
	fmt.Println(names == nil) // true

	// nil 切片是可以直接使用append操作的
	names = append(names, "joe")

	ages := make([]int, 3)
	ages[1] = 1
	ages[2] = 2
	// ages[3] = 3 // error
	ages = append(ages, 18)
	println(ages) // 自动扩容了

	// slice 还可以 `copy`。这里我们创建一个空的和 `ages` 有相同长度的 slice——`ages_backup`，
	ages_backup := make([]int, len(ages))
	copy(ages_backup, ages)
	fmt.Println("ages_backup:", ages_backup)
	fmt.Println("ages:", ages_backup)
	ages_backup[2] = 9 // ages,ages_backup 两个变量的底层数据同时变了，说明共用了底层数组
	fmt.Println("ages_backup:", ages_backup)
	fmt.Println("ages:", ages_backup)

	// 切片操作
	seqs := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("seqs[:] %v\n", seqs[:])
	fmt.Printf("seqs[3:5] %v\n", seqs[3:5])
	fmt.Printf("seqs[:5] %v\n", seqs[:5])
	fmt.Printf("seqs[3:] %v\n", seqs[3:])
	/*
		seqs[:] [0 1 2 3 4 5 6 7 8 9]
		seqs[3:5] [3 4]
		seqs[:5] [0 1 2 3 4]
		seqs[3:] [3 4 5 6 7 8 9]
		取前不取后，没有取的表示全部
	*/

	a := unsafe.Sizeof(seqs)
	println(a)

	arr1 := [3]int{1, 2, 3}
	arr2 := [4]int{1, 2, 3, 4}
	arr3 := [4]int{1, 2, 3, 4}
	println(reflect.TypeOf(arr1).String())
	println(reflect.TypeOf(arr2).String())
	println(reflect.TypeOf(arr3).String())
}
