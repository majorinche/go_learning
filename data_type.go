package main

import (
	"fmt"
)

func main() {
	var b bool
	fmt.Println(b)
	fmt.Println(&b)
	test_tt(&b) //&b就代表你需要传递参数的函数必须是*bool的形式，才可以使用&b
}

func test_tt(ccc *bool) {
	fmt.Print(ccc)
}
