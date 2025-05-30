package logger

import (
	"runtime"
	"strings"

	"github.com/duiyuan/godemo/pkg/util"
	"github.com/spf13/pflag"
)

type Options struct {
	Level          string   `json:"level" mapstructure:"level"`
	OutputPaths    []string `json:"output-paths" mapstructure:"output-paths"`
	ErrOutputPaths []string `json:"error-output-paths" mapstructure:"error-output-paths"`
}

func NewOptions() *Options {
	return &Options{
		Level:          "",
		OutputPaths:    []string{},
		ErrOutputPaths: []string{},
	}
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	if fs == nil {
		return
	}

	fs.StringVar(&o.Level, "log.level", o.Level, "level of logs")
	fs.StringSliceVar(&o.OutputPaths, "log.output-paths", o.OutputPaths, "output paths of log")
	fs.StringSliceVar(&o.ErrOutputPaths, "log.error-output-paths", o.ErrOutputPaths, "error output paths of log")
}

func (o *Options) Validate() []error {
	return []error{}
}

func (o *Options) Complete() error {
	if runtime.GOOS == "windows" {
		if len(o.OutputPaths) > 0 {
			index := util.FindIndex(o.OutputPaths, func(s string) bool { return strings.HasSuffix(s, ".log") })
			if index > -1 {
				file := o.OutputPaths[index]
				o.OutputPaths[index] = util.FixWindowsPath(file)
			}
		}
		if len(o.ErrOutputPaths) > 0 {
			index := util.FindIndex(o.ErrOutputPaths, func(s string) bool { return strings.HasSuffix(s, ".log") })
			if index > -1 {
				file := o.ErrOutputPaths[index]
				o.ErrOutputPaths[index] = util.FixWindowsPath(file)
			}
		}
	}
	return nil
}
