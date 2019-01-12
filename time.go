package main
import (
	"fmt"
	"time"
	"runtime"
	"github.com/sirupsen/logrus"
	sdkVersion "github.com/operator-framework/operator-sdk/version"
)

func main() {
	t := time.Now()
	fmt.Printf("Go launched at %s\n", t.Local())
	logrus.Infof("Go Version: %s", runtime.Version())
	logrus.Infof("operator-sdk Version: %v", sdkVersion.Version)
}
