/*
	20190819 - flag是解析命令含参数的包
	使用方法：
F:\go\src\github.com\majorinche\go_learning>go run go_flag.go -s ccc -d ffff
value of s: 0xc000030210
value of d: ffff

*/
package main

import (
	"flag"
	"fmt"
)

func main() {
	aa := flag.String("s", "hw", "s h text")   //first is flag,second is value, third maybe help message
	bb := flag.String("d", "dddd", "s h text") //第二个是一个默认值
	flag.Parse()
	fmt.Println("value of s:", aa)  // 默认flag的返回值是内存地址，即指针，如果需要返回实际值，需要加*号
	fmt.Println("value of d:", *bb) // output is: value of d: dddd
}
