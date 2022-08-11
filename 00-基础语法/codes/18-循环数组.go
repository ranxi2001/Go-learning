package main

import (
	"fmt"
)

func main() {
	//person是一个数组
	person := [3] string {"Tom", "Aaron", "John"}
	fmt.Printf("len=%d cap=%d array=%v\n",len(person),cap(person),person)

	fmt.Println("")

	//range 对象  这种格式的循环可以对字符串、数组、切片等进行迭代输出元素，k为迭代计数器。
	for k, v := range person {
		fmt.Printf("person[%d]: %s\n", k, v)
	}

	fmt.Println("")
	//读取key，对于数组而言，key就是索引
	for i := range person {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}

	fmt.Println("")
	//传统的数组索引
	for i := 0; i < len(person); i++ {
		fmt.Printf("person[%d]: %s\n", i, person[i])
	}

	fmt.Println("")

	//使用空白符，只读取value
	for _, name := range person {
		fmt.Println("name :", name)
	}
}
