// Package utils @Author Bing
// @Date 2023/11/9 10:38:00
// @Desc
package utils

import (
	"log"
	"os"
)

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	DebugLog *log.Logger
)

func InitLogger() {
	InfoLog = log.New(os.Stdout, "\033[1;34m[info]\t\033[0m", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(os.Stdout, "\033[1;31m[error]\t\033[0m", log.Ldate|log.Ltime|log.Lshortfile)
	DebugLog = log.New(os.Stdout, "\033[1;37m[debug]\t\033[0m", log.Ldate|log.Ltime|log.Lshortfile)
}
