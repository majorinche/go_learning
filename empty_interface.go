/*
	20210120 - 明确空接口，空接口通道相关知识点
*/
package main

import (
	"fmt"
)

func main() {

	slice := make([]interface{}, 7) //这里7理解为队列的长度，元素的长度等等
	slice2 := make([]interface{}, 7)
	slice3 := make(chan []interface{}, 7)
	slice4 := make(chan int, 7)
	map1 := make(map[string]string)
	map2 := make(map[string]int)

	map1["Command"] = "ping" //字符串类型的映射
	map2["TaskID"] = 1       //整数型的映射
	//map3 := make(map[string]map[string]string)
	//map3["mapvalue"] = map1

	slice[0] = map1
	slice[1] = map2
	//slice[3] = map3
	slice3 <- nil //给通道怎么赋值？
	slice3 <- nil
	slice3 <- nil
	slice3 <- nil
	aaa := <-slice3 //通道给变量赋值可以，但是怎么给通道赋值，只会一个nil了，其他不会

	slice4 <- 10 //这个通道是整数型，很好赋值，主要给空接口通道赋值不会，或者压根就不需要赋值
	origin := <-slice4
	//fmt.Println(slice[3])
	fmt.Println("正常的映射：", slice, "nil代表映射没有用满，还剩余几个名额")
	fmt.Println("空接口，没有赋值，所以都是nil:", slice2)
	fmt.Println("给空接口通道赋了4次，但是用了一次，所以结果是3：", len(slice3))
	fmt.Println("从空接口通道中取了一次值出来，值本身是赋值过的nil：", aaa)
	fmt.Println("给一个整数通道付了一次值，通道本身打印出来是内存地址：", slice4, "转换一下就是原始值：", origin)
}
