package utils

import (
	"fmt"
	"os"
)

type LogFun func(msg string, args ...interface{})

var (
	Log LogFun
)

func init()  {
	if os.Getenv("tracing") == "true" {
		Log = func(msg string, args ...interface{}) {
			fmt.Printf(msg + "\n", args...)
		}
	} else {
		Log = func(msg string, args ...interface{}) {
			
		}
	}
}