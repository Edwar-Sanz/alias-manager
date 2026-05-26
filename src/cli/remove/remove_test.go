package remove

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
	if !strings.HasPrefix(cmd.Use, "r ") {
		t.Errorf("Use should start with 'r ', got %q", cmd.Use)
	}
	wantAliases := map[string]bool{"remove": true, "rm": true, "delete": true, "del": true}
	for _, a := range cmd.Aliases {
		delete(wantAliases, a)
	}
	if len(wantAliases) != 0 {
		t.Errorf("missing aliases: %v (have %v)", wantAliases, cmd.Aliases)
	}
}

func TestCreateCommand_RequiresExactlyOneArg(t *testing.T) {
	cmd := CreateCommand()
	if cmd.Args == nil {
		t.Fatal("Args validator should be set")
	}
	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"zero", nil, true},
		{"one", []string{"foo"}, false},
		{"two", []string{"foo", "bar"}, true},
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
	for _, s := range []string{"Remove alias", "am r"} {
		if !strings.Contains(out, s) {
			t.Errorf("description missing %q", s)
		}
	}
}
