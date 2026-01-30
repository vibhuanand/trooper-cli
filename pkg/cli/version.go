package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/vibhuanand/trooper-cli/internal/buildinfo"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Fprintf(cmd.OutOrStdout(),
			"trooper %s (commit=%s, date=%s)\n",
			buildinfo.Version, buildinfo.Commit, buildinfo.Date,
		)
		return nil
	},
}
