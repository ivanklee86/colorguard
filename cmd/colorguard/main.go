package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"os"
	"strings"

	"github.com/ivanklee86/colorguard/pkg/colorguard"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Build information (injected by goreleaser).
	version = "dev"
)

const (
	defaultConfigFilename = "bouncer"
	envPrefix             = "BOUNCER"
)

// main function.
func main() {
	command := NewRootCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}

func NewRootCommand() *cobra.Command {
	colorguard := colorguard.New()

	cmd := &cobra.Command{
		Use:     "bouncer",
		Short:   "Enforce minimum versions in CI/CD.",
		Long:    "A CLI to enforce minimum versions for packages in CI/CD.",
		Version: version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			colorguard.Out = cmd.OutOrStdout()
			colorguard.Err = cmd.ErrOrStderr()

			return initializeConfig(cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprint(colorguard.Out, cmd.UsageString())
		},
	}

	cmd.PersistentFlags().BoolVar(&colorguard.NoExitCode, "no-exit-on-fail", false, "Don't return a non-zero exit code on failure.")

	return cmd
}

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()

	v.SetConfigName(defaultConfigFilename)
	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	v.SetEnvPrefix(envPrefix)
	v.AutomaticEnv()
	bindFlags(cmd, v)

	return nil
}

func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			if err := v.BindEnv(f.Name, fmt.Sprintf("%s_%s", envPrefix, envVarSuffix)); err != nil {
				os.Exit(1)
			}
		}

		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			if err := cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val)); err != nil {
				os.Exit(1)
			}
		}
	})
}
