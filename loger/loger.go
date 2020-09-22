package loger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	outfile, _ = os.OpenFile("logs/logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	logFile    = log.New(outfile, "", 0)
)

func ForLog(v ...interface{}) {
	logFile.Println(fmt.Sprintf("%v: %v", time.Now(), v))
}
