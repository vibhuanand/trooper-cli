package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	defaultConfigFile = "troop.yaml"
)

func newInitCmd() *cobra.Command {
	var force bool
	var withGitHub bool

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize Troop in the current repository",
		Long:  "Creates a YAML-first Troop project (troop.yaml). Optionally scaffolds GitHub Actions workflows.",
		RunE: func(cmd *cobra.Command, args []string) error {
			cfgPath := defaultConfigFile

			// Handle existing config
			if _, err := os.Stat(cfgPath); err == nil && !force {
				return fmt.Errorf("%s already exists (use --force to overwrite)", cfgPath)
			}

			// Write starter config (repo root)
			if err := os.WriteFile(cfgPath, []byte(defaultTroopYAML()), 0o644); err != nil {
				return fmt.Errorf("write %s: %w", cfgPath, err)
			}

			if withGitHub {
				// .github/workflows
				prPlan := filepath.Join(".github", "workflows", "pr-plan.yml")
				mainApply := filepath.Join(".github", "workflows", "main-apply.yml")
				if err := os.MkdirAll(filepath.Dir(prPlan), 0o755); err != nil {
					return fmt.Errorf("create workflows dir: %w", err)
				}

				// Don’t overwrite workflows unless --force
				if err := writeIfAllowed(prPlan, []byte(defaultGHAPlanYAML()), force); err != nil {
					return err
				}
				if err := writeIfAllowed(mainApply, []byte(defaultGHAApplyYAML()), force); err != nil {
					return err
				}

				// CODEOWNERS (optional but helpful)
				codeowners := filepath.Join(".github", "CODEOWNERS")
				_ = writeIfAllowed(codeowners, []byte("* @platform-team\n"), force) // safe default; users can edit
			}

			fmt.Fprintf(cmd.OutOrStdout(), "initialized %s\n", cfgPath)
			if withGitHub {
				fmt.Fprintln(cmd.OutOrStdout(), "scaffolded .github/workflows (pr-plan.yml, main-apply.yml)")
			}
			return nil
		},
	}

	cmd.Flags().BoolVar(&force, "force", false, "Overwrite existing configuration if present")
	cmd.Flags().BoolVar(&withGitHub, "with-github", true, "Scaffold GitHub Actions workflows")
	return cmd
}

func writeIfAllowed(path string, content []byte, force bool) error {
	if _, err := os.Stat(path); err == nil && !force {
		return fmt.Errorf("%s already exists (use --force to overwrite)", path)
	}
	if err := os.WriteFile(path, content, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}

func defaultTroopYAML() string {
	// YAML-first "catalog" starter. Keep minimal but extensible.
	return `apiVersion: troop.dev/v1
kind: Project

metadata:
  name: troop-demo
  owner: platform-team

spec:
  cloud: azure
  cicd: github-actions

  moduleSource:
    repo: git::ssh://github.com/<org>/trooper-modules.git
    ref: v0.1.0

  state:
    mode: troop-managed
    backend: azurerm

  environments:
    - name: dev
      region: canadacentral
    - name: prod
      region: canadacentral

  # Teams define resources here. Troop generates terraform/terragrunt at runtime.
  resources: []
`
}

func defaultGHAPlanYAML() string {
	// Minimal workflow: validate + plan on PR.
	// Install step is placeholder until you wire releases or go install pinned SHA.
	return `name: PR Plan

on:
  pull_request:

jobs:
  plan:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Install troop
        run: |
          echo "TODO: install troop (release asset or go install)."
          # go install github.com/<org>/trooper-cli/cmd/troop@latest

      - name: Plan (dev)
        run: |
          troop plan --env dev
`
}

func defaultGHAApplyYAML() string {
	// Minimal workflow: apply on main. Add approvals via GitHub Environments later.
	return `name: Main Apply

on:
  push:
    branches: [ main ]

jobs:
  apply:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Install troop
        run: |
          echo "TODO: install troop (release asset or go install)."

      - name: Apply (dev)
        run: |
          troop apply --env dev
`
}
