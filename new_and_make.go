package main

import (
	"fmt"
)

func main() {
	var i *int
	i = new(int) // we must assign memory address to type, we heard about value need memory, but the type still need it, or it will report error: panic: runtime error: invalid memory address or nil pointer dereference, another point is: return value is a pointer value, so next we define *i=10, mean the real value of pointer is 10
	*i = 10
	fmt.Println(*i)
	// for make, it only assign memeory for three type: slice, map, and channel
	// ch: = make(chan string), it mean this channel will used to communicate the string data
}
