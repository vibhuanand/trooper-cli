package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	defaultDir  = ".trooper"
	defaultFile = "trooper.yaml"
)

func newInitCmd() *cobra.Command {
	var force bool

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize Trooper in the current repository",
		Long:  "Creates a .trooper folder with a starter trooper.yaml configuration file.",
		RunE: func(cmd *cobra.Command, args []string) error {
			dirPath := defaultDir
			cfgPath := filepath.Join(dirPath, defaultFile)
			gitkeepPath := filepath.Join(dirPath, ".gitkeep")

			// Create directory if missing
			if err := os.MkdirAll(dirPath, 0o755); err != nil {
				return fmt.Errorf("create %s: %w", dirPath, err)
			}

			// Handle existing config
			if _, err := os.Stat(cfgPath); err == nil && !force {
				return fmt.Errorf("%s already exists (use --force to overwrite)", cfgPath)
			}

			// Write starter config
			content := []byte(defaultTrooperYAML())
			if err := os.WriteFile(cfgPath, content, 0o644); err != nil {
				return fmt.Errorf("write %s: %w", cfgPath, err)
			}

			// Ensure .gitkeep exists
			if _, err := os.Stat(gitkeepPath); os.IsNotExist(err) {
				if err := os.WriteFile(gitkeepPath, []byte(""), 0o644); err != nil {
					return fmt.Errorf("write %s: %w", gitkeepPath, err)
				}
			}

			fmt.Fprintf(cmd.OutOrStdout(), "initialized %s\n", cfgPath)
			return nil
		},
	}

	cmd.Flags().BoolVar(&force, "force", false, "Overwrite existing configuration if present")
	return cmd
}

func defaultTrooperYAML() string {
	// Keep this minimal. We'll evolve schema later.
	return `version: v1
project:
  name: trooper-demo
workflows:
  - name: validate
    steps:
      - run: echo "hello from trooper"
`
}
