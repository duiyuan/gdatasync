package app

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/marmotedu/component-base/pkg/util/homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const configFlagName = "config"

var cfgFile string

// nolint: gochecknoinits
func init() {
	pflag.StringVarP(&cfgFile, "config", "c", cfgFile, "Read configuration from specified `FILE`, "+
		"support JSON, TOML, YAML, HCL, or Java properties formats.")
}

// addConfigFlag adds flags for a specific server to the specified FlagSet object.
func addConfigFlag(basename string, fs *pflag.FlagSet) {
	fs.AddFlag(pflag.Lookup(configFlagName))

	viper.AutomaticEnv()
	viper.SetEnvPrefix(strings.Replace(strings.ToUpper(basename), "-", "_", -1))
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	cobra.OnInitialize(func() {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			viper.AddConfigPath(".")

			if names := strings.Split(basename, "-"); len(names) > 1 {
				dir := names[0]
				p0 := filepath.Join(homedir.HomeDir(), "."+dir)
				p1 := filepath.Join(homedir.HomeDir(), dir, "etc")
				p2 := filepath.Join("/etc", dir)

				viper.AddConfigPath(p0)
				viper.AddConfigPath(p1)
				viper.AddConfigPath(p2)

				log.Printf("Add config path: %s \n", p0)
				log.Printf("Add config path: %s \n", p1)
				log.Printf("Add config path: %s \n", p2)

			}

			viper.SetConfigName(basename)
		}

		if err := viper.ReadInConfig(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read configuration file(%s): %v\n", cfgFile, err)
			os.Exit(1)
		}

		v1 := viper.GetString("provider.wss")
		v2 := viper.GetString("provider.http")

		fmt.Println(v1, v2)
	})
}
