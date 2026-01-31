package cli

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/vibhuanand/trooper-cli/internal/config"
	"github.com/vibhuanand/trooper-cli/internal/runner"
)

func newRunCmd() *cobra.Command {
	var cfgPath string
	var dryRun bool
	var shell string

	cmd := &cobra.Command{
		Use:   "run <workflow>",
		Short: "Run a workflow from the Troop config",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			workflowName := args[0]

			if cfgPath == "" {
				cfgPath = filepath.Join(".trooper", "trooper.yaml")
			}

			cfg, err := config.Load(cfgPath)
			if err != nil {
				return err
			}

			wf, ok := cfg.FindWorkflow(workflowName)
			if !ok {
				return fmt.Errorf("workflow %q not found in %s", workflowName, cfgPath)
			}

			r := runner.New()
			r.DryRun = dryRun
			if shell != "" {
				r.Shell = shell
			}

			return r.RunWorkflow(context.Background(), *wf)
		},
	}

	cmd.Flags().StringVar(&cfgPath, "config", "", "Path to config file (default .trooper/trooper.yaml)")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Print steps without executing them")
	cmd.Flags().StringVar(&shell, "shell", "", "Shell to use for run steps (default sh)")
	return cmd
}
