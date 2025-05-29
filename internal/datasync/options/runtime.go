package options

import (
	"strings"

	"github.com/spf13/pflag"
)

type RuntimeOption struct {
	WSS         string   `json:"wss" mapstructure:"wss"`
	Http        string   `json:"http" mapstructure:"http"`
	OutputPaths []string `json:"output-paths" mapstructure:"output-paths"`
}

func NewRuntimeOption() *RuntimeOption {
	return &RuntimeOption{}
}

func (s *RuntimeOption) AddFlags(fs *pflag.FlagSet) {
	if fs == nil {
		return
	}

	fs.StringVar(&s.WSS, "runtime.wss", s.WSS, "Websocket endpoint")
	fs.StringVar(&s.Http, "runtime.http", s.Http, "http endpoint")
	fs.StringSliceVar(&s.OutputPaths, "runtime.output-paths", s.OutputPaths, "http endpoint")
}

func (s *RuntimeOption) Validate() []error {
	err := []error{}
	return err
}

func (s *RuntimeOption) Complete() error {
	s.WSS = strings.Trim(s.WSS, " ")
	return nil
}
