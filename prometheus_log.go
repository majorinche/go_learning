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
