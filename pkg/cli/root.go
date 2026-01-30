package cli

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "trooper",
	Short:         "Trooper: modular DevOps/Infra automation CLI.",
	SilenceUsage:  true,
	SilenceErrors: true,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(healthCmd)
	rootCmd.AddCommand(versionCmd)
}
