package main

import (
	"os"
	"runtime"

	"github.com/duiyuan/gotest/internal/datasync"
)

func main() {
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	datasync.NewApp("gotest-datasync").Run()
}
