/*
	20190819 - 测试prometheus的日志功能
	20210120 - go的依赖关系让我很头疼，不明显，我这里明明没有prometheus，居然运行时成功，log方法也正常调用
*/
package main

import (
	"github.com/prometheus/common/log"
)

func main() {
	log.Infoln("Starting ipmi_exporter mc_20190128")
	var test *string
	test = new(string)
	*test = "majorin"
	log.Infof("Starting ipmi_exporter mc_20190128: %s", test)
	log.Infof("Starting ipmi_exporter mc_20190128: %s", *test)
	log.Infof("Starting ipmi_exporter mc_20190128: %s , %s", *test, test)
}
