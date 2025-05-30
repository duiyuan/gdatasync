package subscriber

import (
	"strings"
	"sync"

	"github.com/duiyuan/gotest/internal/datasync/options"
	"github.com/duiyuan/gotest/internal/datasync/pkg/connection"
	"github.com/duiyuan/gotest/pkg/logger"
)

func MakeSubscriber(opts *options.Options, tunnel string, wg *sync.WaitGroup, handler connection.Handler) *connection.SubscriberConn {
	runtimeOptions, loggerOptions := opts.RuntimeOptions, opts.Log

	wg.Add(1)

	path := strings.Replace(loggerOptions.OutputPaths[0], "<tunnel>", tunnel, -1)

	logger := logger.NewLogger(path)
	sub := connection.NewSubscriberConn(runtimeOptions.WSS, tunnel, wg, logger)
	sub.SetHandler(handler)
	go sub.Connect()
	return sub
}
