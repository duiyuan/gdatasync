package logger

import (
	"runtime"
	"strings"

	"github.com/duiyuan/godemo/pkg/util"
	"github.com/spf13/pflag"
)

type Options struct {
	OutputPaths    []string `json:"output_paths" mapstructure:"output_paths"`
	ErrOutputPaths []string `json:"error_output_paths" mapstructure:"error_output_paths"`
}

func NewOption() *Options {
	return &Options{}
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	if fs == nil {
		return
	}

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
