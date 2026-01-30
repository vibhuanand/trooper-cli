package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadAndFindWorkflow(t *testing.T) {
	tmp := t.TempDir()
	p := filepath.Join(tmp, "trooper.yaml")

	err := os.WriteFile(p, []byte(`
version: v1
project:
  name: test
workflows:
  - name: validate
    steps:
      - run: echo "hello"
`), 0o644)
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(p)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	wf, ok := cfg.FindWorkflow("validate")
	if !ok || wf == nil {
		t.Fatalf("expected to find workflow validate")
	}
	if wf.Name != "validate" {
		t.Fatalf("expected workflow name validate, got %q", wf.Name)
	}
}
