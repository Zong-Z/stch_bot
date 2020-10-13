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

// ForLog this is a function to display the passed arguments
// in the console and save to a file.
func ForLog(v ...interface{}) {
	str := fmt.Sprintf("%v: %v", time.Now().Format("01-02-2006 15:04:05.0000"), v)
	logFile.Println(str)
	log.Println(str)
}
