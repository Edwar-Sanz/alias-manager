package root

import (
	"strings"
	"testing"
)

func TestCreateCommand(t *testing.T) {
	cmd := CreateCommand()
	if cmd == nil {
		t.Fatal("CreateCommand returned nil")
	}
	if cmd.Use != "am" {
		t.Errorf("Use = %q, want %q", cmd.Use, "am")
	}
	if !cmd.SilenceErrors {
		t.Error("SilenceErrors should be true")
	}
	if !cmd.SilenceUsage {
		t.Error("SilenceUsage should be true")
	}
}

func TestBuildLongDescription(t *testing.T) {
	out := BuildLongDescription()
	for _, s := range []string{
		"Alias Manager",
		"List aliases",
		"Add or update an alias",
		"Remove an alias",
		"am l",
		"am a",
		"am r",
	} {
		if !strings.Contains(out, s) {
			t.Errorf("output missing %q\n---\n%s", s, out)
		}
	}
}
