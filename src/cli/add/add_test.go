package add

import (
	"strings"
	"testing"

	"github.com/spf13/cobra"
)

func TestCreateCommand(t *testing.T) {
	cmd := CreateCommand()
	if cmd == nil {
		t.Fatal("CreateCommand returned nil")
	}
	if !strings.HasPrefix(cmd.Use, "a ") {
		t.Errorf("Use should start with 'a ', got %q", cmd.Use)
	}
	if len(cmd.Aliases) == 0 || cmd.Aliases[0] != "add" {
		t.Errorf("expected 'add' alias, got %v", cmd.Aliases)
	}
	if cmd.Args == nil {
		t.Fatal("Args validator should be set")
	}
}

func TestCreateCommand_ArgsValidation(t *testing.T) {
	cmd := CreateCommand()
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"too few", []string{"only-one"}, true},
		{"min", []string{"name", "cmd"}, false},
		{"with category", []string{"name", "cmd", "cat"}, false},
		{"with desc", []string{"name", "cmd", "cat", "desc"}, false},
		{"too many", []string{"a", "b", "c", "d", "e"}, true},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := cmd.Args(&cobra.Command{}, tc.args)
			if (err != nil) != tc.wantErr {
				t.Errorf("Args(%v) err=%v, wantErr=%v", tc.args, err, tc.wantErr)
			}
		})
	}
}

func TestBuildDescription(t *testing.T) {
	out := buildDescription()
	if out == "" {
		t.Fatal("expected non-empty description")
	}
}
