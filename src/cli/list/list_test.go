package list

import (
	"strings"
	"testing"
)

func TestCreateCommand(t *testing.T) {
	cmd := CreateCommand()
	if cmd == nil {
		t.Fatal("CreateCommand returned nil")
	}
	if cmd.Use != "l" {
		t.Errorf("Use = %q, want %q", cmd.Use, "l")
	}
	wantAliases := map[string]bool{"list": true, "ls": true}
	for _, a := range cmd.Aliases {
		delete(wantAliases, a)
	}
	if len(wantAliases) != 0 {
		t.Errorf("missing aliases: %v (have %v)", wantAliases, cmd.Aliases)
	}
}

func TestCreateCommand_Flags(t *testing.T) {
	cmd := CreateCommand()

	cFlag := cmd.Flags().Lookup("category")
	if cFlag == nil {
		t.Fatal("--category flag missing")
	}
	if cFlag.Shorthand != "c" {
		t.Errorf("category shorthand = %q, want %q", cFlag.Shorthand, "c")
	}

	bigCFlag := cmd.Flags().Lookup("categories")
	if bigCFlag == nil {
		t.Fatal("--categories flag missing")
	}
	if bigCFlag.Shorthand != "C" {
		t.Errorf("categories shorthand = %q, want %q", bigCFlag.Shorthand, "C")
	}
}

func TestBuildDescription(t *testing.T) {
	out := buildDescription()
	if out == "" {
		t.Fatal("expected non-empty description")
	}
	for _, s := range []string{"List aliases", "am l", "--category", "--categories"} {
		if !strings.Contains(out, s) {
			t.Errorf("description missing %q", s)
		}
	}
}
