package main

import (
	"fmt"
	"net"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Errorf("error")
	}
	for _, addr := range addrs { // 这里必须有一个下划线，否则结果是index，而不是实际值
		fmt.Println(addr)
	}
}
