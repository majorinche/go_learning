//运行方式：go run http_server.go
//访问方式：curl http://localhost:8000/fff/

package main

import "net/http"
import "fmt"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r) //output is: &{GET /fff/ HTTP/1.1 1 1 map[Accept:[*/*] User-Agent:[curl/7.29.0]] {} <nil> 0 [] false localhost:8000 map[] map[] <nil> map[] [::1]:43424 /fff/ <nil> <nil> <nil> 0xc42005a300}
	fmt.Println("打印指针变量 r 的内存地址 &r")
	var aa = &r
	fmt.Println(aa)                  // output is: 0xc42000e058
	w.Write([]byte("Hello World\n")) // []byte 切片字节
}

func main() {
	http.HandleFunc("/", helloWorld) // 第一个是模式/路径，第二个是访问这个路径对应的响应函数
	http.ListenAndServe(":8000", nil)
}
