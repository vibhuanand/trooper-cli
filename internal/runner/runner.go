package runner

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/vibhuanand/trooper-cli/internal/config"
)

type Runner struct {
	Shell string
}

func New() *Runner {
	// Cross-platform enough for now; we can add OS-specific shell selection later.
	return &Runner{Shell: "sh"}
}

func (r *Runner) RunWorkflow(ctx context.Context, wf config.Workflow) error {
	if len(wf.Steps) == 0 {
		return fmt.Errorf("workflow %q has no steps", wf.Name)
	}

	for i, step := range wf.Steps {
		if step.Run == "" {
			return fmt.Errorf("workflow %q step %d has empty 'run'", wf.Name, i+1)
		}

		fmt.Fprintf(os.Stdout, "==> %s: step %d: %s\n", wf.Name, i+1, step.Run)

		cmd := exec.CommandContext(ctx, r.Shell, "-c", step.Run)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("workflow %q step %d failed: %w", wf.Name, i+1, err)
		}
	}

	return nil
}
