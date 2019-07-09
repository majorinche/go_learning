package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	//	"strings"
)

func checkRegexp(cont string, reg string, style int) (result interface{}) {
	check := regexp.MustCompile(reg)
	switch style {
	case 0:
		result = check.FindString(cont)
	case 1:
		result = check.FindAllString(cont, -1)
	default:
		result = check.FindAll([]byte(cont), -1)
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

	f, _ := os.OpenFile("/data/url.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	defer f.Close()

	for _, url := range reg.FindAllString(string(body), -1) {

		url_file, _ := os.OpenFile("/data/url.txt", os.O_RDWR, 0666)
		file, _ := ioutil.ReadAll(url_file)
		/***
				dd := strings.Split(url, "")
				fmt.Print(dd)
				dddd := ""
				for _, ddd := range dd {
					if ddd == "?" {
						ddd = `\?`
					}
					dddd += ddd
				}
		***/
		if checkRegexp(string(file), url, 0).(string) == "" {
			io.WriteString(f, url+"\n")
			fmt.Print("\n收集地址：" + url + "\n")
			num++
		}
		url_file.Close()
	}
	fmt.Print("\n首次收集网络地址：" + strconv.Itoa(len(reg.FindAllString(string(body), -1))) + "\n")
	fmt.Print("\n去重后网络地址数：" + strconv.Itoa(num))
	fmt.Print("\n\n首次储存成功！\n")
}

func main() {
	fistStart()
}
