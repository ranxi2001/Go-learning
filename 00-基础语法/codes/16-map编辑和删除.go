package main

import (
	"fmt"
)

func main() {
	person := map[int]string{
		1 : "Tom",
		2 : "Aaron",
		3 : "John",
	}
	fmt.Println("data :",person)
	//删除
	delete(person, 2)
	fmt.Println("data :",person)
	//编辑
	person[2] = "Jack"
	person[3] = "Kevin"
	fmt.Println("data :",person)
}
