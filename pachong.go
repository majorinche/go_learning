package main

/***
	20190709, bugfix, finish dedup
	1, function is capture sensetive
	2, result of function must define interface{}


***/

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkRegexp(file string, url string, style int) (result interface{}) {
	index_num := strings.Index(url, "=") + 1 // using url to do regular expression is not useurl, we mush extract info following equal signal
	url_link := url[index_num:]
	check_url := regexp.MustCompile(url_link)
	switch style {
	case 0:
		result = check_url.FindString(file)
	case 1:
		result = check_url.FindAllString(file, -1)
	default:
		result = check_url.FindAll([]byte(file), -1)
	}
	return
}

func fistStart() {
	var num int
	url := "http://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=39042058_20_oem_dg&wd=golang%E5%AE%9E%E7%8E%B0&oq=golang%2520%25E5%2588%25A0%25E9%2599%25A4%25E6%2595%25B0%25E7%25BB%2584&rsv_pq=d9be28ec0002df1b&rsv_t=8017GWpSLPhDmKilZQ1StC04EVpUAeLEP90NIm%2Bk5pRh5R9o57NHMO8Gaxm1TtSOo%2FvtJj%2B98%2Fsc&rqlang=cn&rsv_enter=1&inputT=3474&rsv_sug3=16&rsv_sug1=11&rsv_sug7=100&rsv_sug2=0&rsv_sug4=4230"
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	reg := regexp.MustCompile(`((ht|f)tps?)://[w]{0,3}.baidu.com/link\?[a-zA-z=0-9-\s]*`)

	os.Truncate("/data/url.txt", 0) //clear file first, cool method
	url_file, _ := os.OpenFile("/data/url.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	defer url_file.Close()

	for _, url := range reg.FindAllString(string(body), -1) {
		url_file, _ := os.OpenFile("/data/url.txt", os.O_RDWR, 0666)
		file, _ := ioutil.ReadAll(url_file)
		if checkRegexp(string(file), url, 0).(string) == "" {
			io.WriteString(url_file, url+"\n")
			fmt.Print("\n收集地址：" + url + "\n")
			num++
		}
		url_file.Close()
	}
	fmt.Print("\n首次收集网络地址：" + strconv.Itoa(len(reg.FindAllString(string(body), -1))) + "\n")
	fmt.Print("\n去重后网络地址数：" + strconv.Itoa(num))
}

func main() {
	fistStart()
}
