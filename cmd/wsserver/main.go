package main

import (
	"os"
	"runtime"

	"github.com/duiyuan/gotest/internal/wss"
)

func main() {
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	server := wss.NewWSServer()
	server.Serve()
}
