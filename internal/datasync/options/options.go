package options

import (
	"github.com/duiyuan/godemo/pkg/logger"
	"github.com/marmotedu/component-base/pkg/cli/flag"
)

type Options struct {
	RuntimeOption *RuntimeOption  `json:"runtime" mapstructure:"runtime"`
	LogOption     *logger.Options `json:"log" mapstruction:"log"`
}

func NewOption() *Options {
	return &Options{
		RuntimeOption: NewRuntimeOption(),
		LogOption:     logger.NewOption(),
	}
}

func (o *Options) Flags() (fss flag.NamedFlagSets) {
	o.RuntimeOption.AddFlags(fss.FlagSet("runtime"))
	o.LogOption.AddFlags(fss.FlagSet("log"))
	return
}

func (o *Options) Complete() error {
	return nil
}

func (o *Options) Validate() []error {
	return []error{}
}
