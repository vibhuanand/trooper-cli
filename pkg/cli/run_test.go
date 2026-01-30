package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRunWorkflowNotFound(t *testing.T) {
	tmp := t.TempDir()
	cfgDir := filepath.Join(tmp, ".trooper")
	if err := os.MkdirAll(cfgDir, 0o755); err != nil {
		t.Fatal(err)
	}

	cfgPath := filepath.Join(cfgDir, "trooper.yaml")
	if err := os.WriteFile(cfgPath, []byte(`
version: v1
project:
  name: test
workflows:
  - name: validate
    steps:
      - run: echo "hello"
`), 0o644); err != nil {
		t.Fatal(err)
	}

	cmd := NewRootCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)

	cmd.SetArgs([]string{"run", "does-not-exist", "--config", cfgPath})
	if err := cmd.Execute(); err == nil {
		t.Fatalf("expected error for missing workflow, got nil")
	}
}
