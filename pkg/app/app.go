package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	"github.com/marmotedu/component-base/pkg/cli/globalflag"
	"github.com/marmotedu/component-base/pkg/version"
	"github.com/marmotedu/component-base/pkg/version/verflag"
	"github.com/marmotedu/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	progressMessage = color.GreenString("==>")
)

type RunFunc func(basename string) error
type SetOption func(*App)

type App struct {
	basename    string
	name        string
	cliOptions  CliOptions
	description string
	run         RunFunc
	silence     bool
	commands    []*Command
	cmd         *cobra.Command
	args        cobra.PositionalArgs
	noVersion   bool
	noConfig    bool
}

func WithOptions(opt CliOptions) SetOption {
	return func(a *App) {
		a.cliOptions = opt
	}
}

func WithRunFunc(run RunFunc) SetOption {
	return func(a *App) {
		a.run = run
	}
}

func WithDesc(desc string) SetOption {
	return func(a *App) {
		a.description = desc
	}
}

func WithSilence(silence bool) SetOption {
	return func(a *App) {
		a.silence = silence
	}
}

func NewApp(name string, basename string, setOptions ...SetOption) *App {
	app := &App{
		basename: basename,
		name:     name,
	}

	for _, s := range setOptions {
		s(app)
	}

	app.buildCommand()
	return app
}

func (a *App) buildCommand() {
	cmd := cobra.Command{
		Use:           FormatBaseName(a.basename),
		Short:         a.name,
		Long:          a.description,
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          a.args,
	}

	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	cmd.Flags().SortFlags = true

	cliflag.InitFlags(cmd.Flags())

	if len(a.commands) > 0 {
		for _, c := range a.commands {
			cmd.AddCommand(c.cobraCommand())
		}
		cmd.SetHelpCommand(helpCommand(FormatBaseName(a.basename)))
	}
	// cmd := cobra.Command -> cmd.Excute() -> cobra.OnInitialize(read cli by viper) -> cmd.RunE
	if a.run != nil {
		cmd.RunE = a.runCommand
	}

	var namedFlagSets cliflag.NamedFlagSets
	if a.cliOptions != nil {
		namedFlagSets = a.cliOptions.Flags()
		pfs := cmd.Flags()
		for _, f := range namedFlagSets.FlagSets {
			pfs.AddFlagSet(f)
		}
	}

	if !a.noVersion {
		verflag.AddFlags(namedFlagSets.FlagSet("global"))
	}

	if !a.noConfig {
		addConfigFlag(a.basename, namedFlagSets.FlagSet("global"))
	}

	globalflag.AddGlobalFlags(namedFlagSets.FlagSet("global"), cmd.Name())

	a.cmd = &cmd
}

func (a *App) runCommand(cmd *cobra.Command, args []string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	log.Infof("%v workdir: %s", progressMessage, pwd)

	cliflag.PrintFlags(cmd.Flags())

	if !a.noVersion {
		verflag.PrintAndExitIfRequested()
	}

	if !a.noConfig {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			return err
		}
		if err := viper.Unmarshal(a.cliOptions); err != nil {
			return err
		}
	}

	if !a.silence {
		log.Infof("%s Starting %s...", progressMessage, a.name)
		if !a.noVersion {
			log.Infof("%v Version: `%s`", progressMessage, version.Get().ToJSON())
		}
		if !a.noConfig {
			log.Infof("%s Config file used: %s", progressMessage, viper.ConfigFileUsed())
		}
	}

	if a.cliOptions != nil {
		if err := a.applyCliOptionRules(); err != nil {
			return err
		}
	}

	if a.run != nil {
		return a.run(a.basename)
	}

	return nil
}

func (a *App) applyCliOptionRules() error {
	if completeableOptions, ok := a.cliOptions.(CompleteableOptions); ok {
		if err := completeableOptions.Complete(); err != nil {
			return err
		}
	}

	if errs := a.cliOptions.Validate(); len(errs) != 0 {
		errlist := []string{}
		for _, e := range errs {
			errlist = append(errlist, e.Error())
		}
		return fmt.Errorf("errors: %s", strings.Join(errlist, ","))
		// return errors.NewAggregate(errs)
	}

	if printableOptions, ok := a.cliOptions.(PrintableOptions); ok && !a.silence {
		log.Infof("%v Config: `%s`", progressMessage, printableOptions.String())
	}

	return nil
}

func (a *App) Run() {
	if err := a.cmd.Execute(); err != nil {
		log.Errorf("%v %v", color.RedString("Error"), err)
		os.Exit(1)
	}
}
