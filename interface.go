/*
	20190819 - 指针体现的是灵活，而普通变量体现的是简单，
	20190821 - 如果是这样，指针的使用倒不是必须的，至少对于小项目来说，不过有些方法的返回值就是指针，你不得不用*号来获取原始值，为什么返回值要设置为指针

*/

package main

import (
	"fmt"
)

type Books struct { //结构体本身
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var Books1 Books
	var Books2 Books

	Books1.title = "高性能网站构建实战"
	Books1.author = "叶毓睿等"
	Books1.subject = "website"
	Books1.book_id = 11

	Books2.title = "高性能网站构建实战"
	Books2.author = "叶毓睿等"
	Books2.subject = "website"
	Books2.book_id = 22
	fmt.Println(Books1)
	printBook(&Books1) //把Books1的内存地址作为参数
	var aa *int
	var bb *int

	cc := 123
	dd := 222
	aa = &cc
	bb = &dd
	//ff = aa + bb, 指针不能进行简单的加操作，那怎么管理指针变量呢？

	fmt.Println(aa)
	fmt.Println(bb)
	//fmt.Print(ff)

}

func printBook(book *Books) { //这里的Books是一个结构体，然后通过*号调用结构体的实际值
	fmt.Printf("Book title: %s\n", book.title)
}
