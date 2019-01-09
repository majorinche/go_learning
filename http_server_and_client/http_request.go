package main

import "net/http"
import "io/ioutil"
import "log"
import "fmt"

func main() {
	response, err := http.Get("http://127.0.0.1:8000/") //请求数据
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body) //读取响应的Body信息，说明response不止Body信息
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
}
