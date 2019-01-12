package main
import (
	"fmt"
	"time"
	"runtime"
	"github.com/sirupsen/logrus"
)

func main() {
	t := time.Now()
	fmt.Printf("Go launched at %s\n", t.Local())
	logrus.Infof("Go Version: %s", runtime.Version())
}
