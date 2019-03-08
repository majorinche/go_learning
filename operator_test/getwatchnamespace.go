package main

import (
	"fmt"
	"os"
)

const(
	WatchNamespaceEnvVar="WATCH_NAMESPACE"
)

func main(){
	ns, found := os.LookupEnv(WatchNamespaceEnvVar)
	fmt.Println(ns)
	fmt.Println(found)
}

