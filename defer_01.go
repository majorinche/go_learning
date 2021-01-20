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
