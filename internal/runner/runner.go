package runner

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/vibhuanand/trooper-cli/internal/config"
)

type Runner struct {
	Shell string
}

func New() *Runner {
	return &Runner{Shell: "sh"}
}

func (r *Runner) RunWorkflow(ctx context.Context, wf config.Workflow) error {
	if len(wf.Steps) == 0 {
		return fmt.Errorf("workflow %q has no steps", wf.Name)
	}

	for i, step := range wf.Steps {
		stepNum := i + 1
		workdir := effectiveWorkdir(wf.Workdir, step.Workdir)

		switch {
		case step.Run != "":
			if err := r.runShell(ctx, wf.Name, stepNum, step.Run, workdir); err != nil {
				return err
			}
		case step.Terraform != nil:
			if err := r.runTool(ctx, wf.Name, stepNum, "terraform", step.Terraform.Args, workdir); err != nil {
				return err
			}
		case step.Kubectl != nil:
			if err := r.runTool(ctx, wf.Name, stepNum, "kubectl", step.Kubectl.Args, workdir); err != nil {
				return err
			}
		default:
			return fmt.Errorf("workflow %q step %d has no supported action (expected one of: run, terraform, kubectl)", wf.Name, stepNum)
		}
	}

	return nil
}

func effectiveWorkdir(wfDir, stepDir string) string {
	if stepDir != "" {
		return stepDir
	}
	return wfDir
}

func (r *Runner) runShell(ctx context.Context, workflow string, stepNum int, cmdStr string, workdir string) error {
	fmt.Fprintf(os.Stdout, "==> %s: step %d: run: %s\n", workflow, stepNum, cmdStr)

	cmd := exec.CommandContext(ctx, r.Shell, "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if workdir != "" {
		cmd.Dir = filepath.Clean(workdir)
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("workflow %q step %d failed (run): %w", workflow, stepNum, err)
	}
	return nil
}

func (r *Runner) runTool(ctx context.Context, workflow string, stepNum int, tool string, args []string, workdir string) error {
	fmt.Fprintf(os.Stdout, "==> %s: step %d: %s %v\n", workflow, stepNum, tool, args)

	path, err := exec.LookPath(tool)
	if err != nil {
		return fmt.Errorf("%s not found in PATH (required for workflow %q step %d)", tool, workflow, stepNum)
	}

	cmd := exec.CommandContext(ctx, path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if workdir != "" {
		cmd.Dir = filepath.Clean(workdir)
	}

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("workflow %q step %d failed (%s): %w", workflow, stepNum, tool, err)
	}
	return nil
}
