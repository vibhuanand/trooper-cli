package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestHealthCommand(t *testing.T) {
	cmd := NewRootCmd()

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"health"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	got := strings.TrimSpace(buf.String())
	if got != "ok" {
		t.Fatalf("expected 'ok', got %q", got)
	}
}

func TestVersionCommand(t *testing.T) {
	cmd := NewRootCmd()

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"version"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	got := strings.TrimSpace(buf.String())
	if !strings.HasPrefix(got, "trooper ") {
		t.Fatalf("expected output to start with 'trooper ', got %q", got)
	}
}
