/*
	20210120 - 这里强调了defer对于先进后出的作用，不过用在并发场景，通道里面是什么不重要，重要的是否消耗了，理解起来可能比较费劲
*/
package main

import "fmt"

func main() {
	var users [5]struct{}
	for i := range users {
		defer fmt.Println(i)
	}
	for i := range users {
		defer func() { fmt.Println(i) }() //这个叫闭包，什么玩意
	}
	for i := range users {
		defer Print(i)
	}
}
func Print(i int) {
	fmt.Println(i)
}
