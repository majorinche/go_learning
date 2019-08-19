/*
	20190819 - 给结构体类型的变量定义函数
*/
package main

import (
	"fmt"
)

type Triangle struct {
	base   float64
	height float64
}

func (test Triangle) area() float64 { //相当于给结构体定义函数
	return 0.5 * (test.base * test.height)
}

func (aaa Triangle) ccc() { //给结构体Triangle定义函数ccc(),而aaa我们叫什么呢，其实叫结构体的变量，访问时，需要在定义相同的变量名访问这个函数
	fmt.Print("ccc")
}

func main() {
	test := Triangle{base: 3, height: 1}
	fmt.Println(test.area())
	aaa := Triangle{} //给结构体定义函数，可以不用自定义结构体的值
	aaa.ccc()
}
