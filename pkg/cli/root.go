package cli

import (
	"github.com/spf13/cobra"
)

// NewRootCmd returns a fresh root command instance.
// This avoids global state and makes the CLI easier to test and extend.
func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:           "trooper",
		Short:         "Trooper: modular DevOps/Infra automation CLI.",
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.AddCommand(newHealthCmd())
	rootCmd.AddCommand(newVersionCmd())

	return rootCmd
}

func Execute() error {
	return NewRootCmd().Execute()
}
