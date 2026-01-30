package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestCompletionCommandBash(t *testing.T) {
	cmd := NewRootCmd()

	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs([]string{"completion", "bash"})

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	out := buf.String()
	if !strings.Contains(out, "bash completion") && len(out) < 50 {
		// completion output is large; basic sanity check
		t.Fatalf("expected non-trivial completion output, got length %d", len(out))
	}
}
