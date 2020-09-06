package loger

import (
	"log"
	"os"
)

var (
	outfile, _ = os.OpenFile("logs/logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	LogFile    = log.New(outfile, "", 0)
)
