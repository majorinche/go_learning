package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	//"os/exec"
)

var (
	configFile = flag.String(
		"config.file", "ipmi.yml", //
		"Path to configuration file.",
	)
)

func main() {
	tt, _ := ioutil.ReadFile(*configFile)                     // 因为flag.string的返回值是指针，所以这里需要*号,ioutil.ReadFile直接读取文件的内容
	fmt.Printf("%s", tt)                                      //here will print out the content in the file, ipmi.yml
	args := []string{}                                        //前面是数组，后面看似是结构体，到底是啥？
	fmt.Printf("%s", args)                                    //here will print out the content in the file, ipmi.yml
	args = append(args, "-h 192.168.2.130 -u admin -p admin") //如果前面定义过args，这里再使用，就可以直接使用等于号了
	fqcmd := "ipmimonitoring"
	out, err := exec.Command(fqcmd, args...).CombinedOutput()
	fmt.Printf("%s %s", out, err)

}
