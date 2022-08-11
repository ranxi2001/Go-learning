package main

import (
	"fmt"
)

func main() {
	var arr =  [5] int {1, 2, 3, 4, 5}
	modifyArr(arr)
	fmt.Println(arr)
}
//定义一个修改数组的函数，定义传入参数的类型为5位数组，不定义返回，直接修改内存
func modifyArr(a [5] int){
	a[1] = 20
}
