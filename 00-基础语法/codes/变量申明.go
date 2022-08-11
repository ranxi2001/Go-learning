package main

import (
	"fmt"
)

func main() {
	//const常量申明①const 名称 类型 = ‘常量’
	const name string = "Tom"
	fmt.Println(name)
	//const常量申明②const 名称 = ‘常量’
	const age = 30
	fmt.Println(age)

	const name_1, name_2 string = "Tom", "Jay"
	fmt.Println(name_1, name_2)

	const name_3, age_1 = "Tom", 30
	fmt.Println(name_3, age_1)
}
