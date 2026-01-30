package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestInitCreatesFiles(t *testing.T) {
	tmp := t.TempDir()
	old, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(old) })

	if err := os.Chdir(tmp); err != nil {
		t.Fatal(err)
	}

	cmd := NewRootCmd()
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"init"})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	cfg := filepath.Join(tmp, ".trooper", "trooper.yaml")
	if _, err := os.Stat(cfg); err != nil {
		t.Fatalf("expected config to exist: %v", err)
	}

	out := strings.TrimSpace(buf.String())
	if !strings.Contains(out, "initialized") {
		t.Fatalf("expected output to include 'initialized', got %q", out)
	}
}

func TestInitRefusesOverwriteWithoutForce(t *testing.T) {
	tmp := t.TempDir()
	old, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.Chdir(old) })

	if err := os.Chdir(tmp); err != nil {
		t.Fatal(err)
	}

	// First init
	cmd1 := NewRootCmd()
	cmd1.SetArgs([]string{"init"})
	if err := cmd1.Execute(); err != nil {
		t.Fatalf("expected no error on first init, got %v", err)
	}

	// Second init should fail without --force
	cmd2 := NewRootCmd()
	cmd2.SetArgs([]string{"init"})
	if err := cmd2.Execute(); err == nil {
		t.Fatalf("expected error on second init without --force, got nil")
	}
}
