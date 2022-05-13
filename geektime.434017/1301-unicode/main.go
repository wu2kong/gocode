package main

import (
	"fmt"
	"os"
	"reflect"
	"unsafe"
)

func dumpBytesArray(arr []byte) {
	fmt.Printf("[")
	for _, b := range arr {
		fmt.Printf("%c ", b)
	}
	fmt.Printf("]\n")
}

func main() {
	hai := '海'
	fmt.Printf("0x%X %d\n", hai, rune(hai)) // 0xE6B5B7
	os.Exit(1)

	var s = "hello"
	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // 将string类型变量地址显式转型为reflect.StringHeader
	fmt.Printf("0x%x\n", hdr.Data)                     // 0x10a30e0
	p := (*[5]byte)(unsafe.Pointer(hdr.Data))          // 获取Data字段所指向的数组的指针
	dumpBytesArray((*p)[:])                            // [h e l l o ]   // 输出底层数组的内容
}
