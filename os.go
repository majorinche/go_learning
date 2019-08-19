/*
	20190819 - os主要作用就是调用系统命令
*/
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("test")
	} else {
		fmt.Printf("localhost hostname is: %s\n", hostname)
	}
	out, err := exec.Command("ls", "-l", "/root/").CombinedOutput() // do not write one command is the same quotes, need seperate as a args...
	fmt.Printf("output is: %s\n", out)
}
