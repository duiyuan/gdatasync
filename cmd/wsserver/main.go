package main

import (
	"os"
	"runtime"

	"gitbub.com/duiyuan/godemo/internal/wss"
)

func main() {
	if len(os.Getenv("GOMAXPROCS")) == 0 {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	server := wss.NewWSServer()
	server.Serve()
}
