package datasync

import (
	"github.com/duiyuan/gotest/internal/datasync/options"
	"github.com/duiyuan/gotest/pkg/app"
)

const desc = "datasync"

func NewApp(basename string) *app.App {
	opts := options.NewOption()

	app := app.NewApp(
		"datasync",
		basename,
		app.WithDesc(desc),
		app.WithOptions(opts),
		app.WithRunFunc(run(opts)),
	)

	return app
}

func run(opts *options.Options) app.RunFunc {
	return func(basename string) error {
		return Start(opts)
	}
}
