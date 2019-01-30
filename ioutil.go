package main

import (
	"flag"
	"fmt"
	"io/ioutil"
)

var (
	configFile = flag.String(
		"config.file", "ipmi.yml", //only ipmi.yml is the real value of configFile
		"Path to configuration file.",
	)
)

func main() {
	tt, _ := ioutil.ReadFile(*configFile) // return of flag.String is a memory address, so need star to retrieve real value, whichi should be ipmi.yml here
	fmt.Printf("%s", tt)                  //here will print out the content in the file, ipmi.yml
}
