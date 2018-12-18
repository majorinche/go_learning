package main
import (
	"fmt"
	"os"
)

func main(){
	hostname,err := os.Hostname()
	if err != nil {
		fmt.Printf("test")
	}else{
		fmt.Printf("localhost hostname is: %s\n", hostname)
	}
}
