package options

import (
	"github.com/duiyuan/godemo/pkg/logger"
	"github.com/marmotedu/component-base/pkg/cli/flag"
)

type Options struct {
	RuntimeOptions *RuntimeOption  `json:"runtime" mapstructure:"runtime"`
	Log            *logger.Options `json:"log" mapstructure:"log"`
}

func NewOption() *Options {
	return &Options{
		RuntimeOptions: NewRuntimeOption(),
		Log:            logger.NewOptions(),
	}
}

func (o *Options) Flags() (fss flag.NamedFlagSets) {
	o.RuntimeOptions.AddFlags(fss.FlagSet("runtime"))
	o.Log.AddFlags(fss.FlagSet("log"))
	return
}

func (o *Options) Complete() error {
	return o.Log.Complete()
}

func (o *Options) Validate() []error {
	return []error{}
}
