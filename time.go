package main
import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("Go launched at %s\n", t.Local())
}
