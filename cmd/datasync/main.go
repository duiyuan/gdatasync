package main

import (
	"os"
	"runtime"

	"gitbub.com/duiyuan/godemo/internal/datasync"
)

func main() {
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	datasync.Start()
}
