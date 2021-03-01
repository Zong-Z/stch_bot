package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// TimeFormat logger time format.
const TimeFormat = "01-02-2006 15:04:05.0000"

var (
	outfile, _ = os.OpenFile("logs/logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	logFile    = log.New(outfile, "", 0)
)

// ForInfo writes the data transferred to the function in the file with the information tag.
func ForInfo(v ...interface{}) {
	str := fmt.Sprintf("%s |INFO|: %v", time.Now().Format(TimeFormat), v)
	logFile.Println(str)
	log.Println(str)
}

// ForWarning writes the data transferred to the function in the file with the warning tag.
func ForWarning(v ...interface{}) {
	str := fmt.Sprintf("%s |WARNING|: %v\n", time.Now().Format(TimeFormat), v)
	logFile.Println(str)
	log.Println(str)
}

// ForError writes the data transferred to the function in the file with the error tag and and stops execution of the
// program.
func ForError(v ...interface{}) {
	str := fmt.Sprintf("%s |ERROR|: %v\n", time.Now().Format(TimeFormat), v)
	logFile.Println(str)
	log.Println(str)
	panic(str)
}
