
//06-修改数组（√）.go

package main

import (
	"fmt"
)

func modifyArr_(a *[5] int){
	a[1] = 20
}

func main() {
	var arr =  [5] int {1, 2, 3, 4, 5}
	modifyArr_(&arr)
	fmt.Println(arr)
}


