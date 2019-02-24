package utils

import "os"

var (
	tracing = false
)

func TracingEnabled() bool {
	return tracing
}

func init()  {
	if os.Getenv("tracing") == "true" {
		tracing = true
	}
}