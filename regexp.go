package main

/***
	20190709, using strings to get index and then slice the string

***/

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	aa := "http://www.baidu.com/link?url=SLlWn5tCJr5JzlF5NPwxIuMVkmligcjaF-MULXr_uKxql9QpwSIdQXbMIcjQ9gci"
	bb := strings.Index(aa, "=") + 1
	cc := aa[bb:]
	fmt.Print(cc)
	check_url := regexp.MustCompile(cc)

	url_file, _ := os.OpenFile("/data/url.txt", os.O_RDWR, 0666)
	file, _ := ioutil.ReadAll(url_file)

	if check_url.FindString(string(file)) == "" {
		fmt.Print("nothing found")
	} else {
		fmt.Print("we can find")
	}
}
