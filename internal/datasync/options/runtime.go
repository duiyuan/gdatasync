package options

import (
	"strings"

	"github.com/spf13/pflag"
)

type RuntimeOption struct {
	WSS  string `json:"wss" mapstructure:"wss"`
	Http string `json:"http" mapstructure:"http"`
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
}

func (s *RuntimeOption) Validate() []error {
	err := []error{}
	return err
}

func (s *RuntimeOption) Complete() error {
	s.WSS = strings.Trim(s.WSS, " ")
	return nil
}
