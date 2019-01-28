package main

import (
	"flag"
	"fmt"
)

func main() {
	aa := flag.String("s", "hw", "s h text") //first is flag,second is value, third maybe help message
	bb := flag.String("d", "dddd", "s h text")
	flag.Parse()
	fmt.Println("value of s:", aa)  // output is: value of s: 0xc00000e200, so start+variable has two meaning: 1, declare a memory address variable, 2, retrieve a real value from memory address
	fmt.Println("value of d:", *bb) // output is: value of d: dddd
}
