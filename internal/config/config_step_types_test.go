package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadStepTypes(t *testing.T) {
	tmp := t.TempDir()
	p := filepath.Join(tmp, "trooper.yaml")

	yml := `
version: v1
project:
  name: test
workflows:
  - name: validate
    workdir: .
    steps:
      - run: echo "hello"
      - terraform:
          args: ["version"]
      - kubectl:
          args: ["version", "--client"]
      - workdir: infra
        run: echo "in infra"
`
	if err := os.WriteFile(p, []byte(yml), 0o644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(p)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	wf, ok := cfg.FindWorkflow("validate")
	if !ok {
		t.Fatalf("expected workflow validate")
	}

	if wf.Workdir != "." {
		t.Fatalf("expected workflow workdir '.', got %q", wf.Workdir)
	}

	if len(wf.Steps) != 4 {
		t.Fatalf("expected 4 steps, got %d", len(wf.Steps))
	}

	if wf.Steps[1].Terraform == nil || len(wf.Steps[1].Terraform.Args) == 0 {
		t.Fatalf("expected terraform args")
	}

	if wf.Steps[2].Kubectl == nil || len(wf.Steps[2].Kubectl.Args) == 0 {
		t.Fatalf("expected kubectl args")
	}

	if wf.Steps[3].Workdir != "infra" {
		t.Fatalf("expected step workdir 'infra', got %q", wf.Steps[3].Workdir)
	}
}
