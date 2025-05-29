package subscriber

import (
	"runtime"
	"strings"
	"sync"

	"github.com/duiyuan/godemo/internal/datasync/options"
	"github.com/duiyuan/godemo/internal/datasync/pkg/connection"
	"github.com/duiyuan/godemo/pkg/logger"
	"github.com/duiyuan/godemo/pkg/util"
)

func MakeSubscriber(opts *options.Options, tunnel string, wg *sync.WaitGroup, handler connection.Handler) *connection.SubscriberConn {
	runtimeOptions, loggerOptions := opts.RuntimeOptions, opts.Log

	wg.Add(1)
	path := strings.Replace(loggerOptions.OutputPaths[0], "<tunnel>", tunnel, -1)

	if runtime.GOOS == "windows" {
		path = util.FixWindowsPath(path)
	}

	logger := logger.NewLogger(path)
	sub := connection.NewSubscriberConn(runtimeOptions.WSS, tunnel, wg, logger)
	sub.SetHandler(handler)
	go sub.Connect()
	return sub
}
