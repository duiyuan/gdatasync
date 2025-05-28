package options

import (
	"github.com/marmotedu/component-base/pkg/cli/flag"
)

type Options struct {
	RuntimeOption *RuntimeOption `json:"runtime" mapstructure:"runtime"`
}

func NewOption() *Options {
	return &Options{
		RuntimeOption: NewRuntimeOption(),
	}
}

func (o *Options) Flags() (fss flag.NamedFlagSets) {
	o.RuntimeOption.AddFlags(fss.FlagSet("runtime"))
	return
}

func (o *Options) Complete() error {
	return nil
}

func (o *Options) Validate() []error {
	return []error{}
}
