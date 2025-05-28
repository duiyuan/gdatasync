package app

import "github.com/marmotedu/component-base/pkg/cli/flag"

type CliOptions interface {
	Flags() (fss flag.NamedFlagSets)
	Validate() []error
}

// CliOptions abstracts configuration options for reading parameters from the
type ConfigurableOptions interface {
	// AddFlags adds flags to the specified FlagSet object.
	// AddFlags(fs *pflag.FlagSet)
	ApplyFlags() []error
}

type CompleteableOptions interface {
	Complete() error
}

type PrintableOptions interface {
	String() string
}
