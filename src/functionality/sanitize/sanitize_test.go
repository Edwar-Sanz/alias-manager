package sanitize

import (
	"testing"

	"am/src/types"
)

func TestName(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"  ", ""},
		{"hello", "hello"},
		{"  hello  ", "hello"},
		{"my-alias_1", "my-alias_1"},
		{"my alias", "myalias"},
		{"my@alias!", "myalias"},
		{"a.b.c", "abc"},
		{"name123", "name123"},
		{"___---", "___---"},
		{"\t\nname\t\n", "name"},
		{"ñame", "ame"},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			got := Name(tc.in)
			if got != tc.want {
				t.Errorf("Name(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestCommand(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"only spaces", "   ", ""},
		{"plain command", "ls -la", "ls -la"},
		{"trim outer spaces", "  ls -la  ", "ls -la"},
		{"double quoted", `"ls -la"`, "ls -la"},
		{"single quoted", `'ls -la'`, "ls -la"},
		{"quoted with inner spaces", `"  ls -la  "`, "ls -la"},
		{"mismatched quotes left", `"ls -la`, `"ls -la`},
		{"mismatched quotes right", `ls -la"`, `ls -la"`},
		{"mismatched mixed", `"ls -la'`, `"ls -la'`},
		{"single char", "x", "x"},
		{"two single quotes", "''", ""},
		{"two double quotes", `""`, ""},
		{"quote inside", `"echo \"hi\""`, `echo \"hi\"`},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Command(tc.in)
			if got != tc.want {
				t.Errorf("Command(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestCategory(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"  git  ", "git"},
		{"git tools", "gittools"},
		{"git/utils", "gitutils"},
		{"a-b_c", "a-b_c"},
		{"1234", "1234"},
		{"!@#", ""},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			got := Category(tc.in)
			if got != tc.want {
				t.Errorf("Category(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestDesc(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"   ", ""},
		{"hello", "hello"},
		{"  hello world  ", "hello world"},
		{"a description!", "a description!"},
		{"\n\thello\n", "hello"},
		{"keeps internal   spaces", "keeps internal   spaces"},
	}
	for _, tc := range tests {
		t.Run(tc.in, func(t *testing.T) {
			got := Desc(tc.in)
			if got != tc.want {
				t.Errorf("Desc(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestEntry(t *testing.T) {
	in := types.AliasEntry{
		Name:     "  my alias!  ",
		Command:  `"ls -la"`,
		Category: "  git tools  ",
		Desc:     "  list files  ",
	}
	want := types.AliasEntry{
		Name:     "myalias",
		Command:  "ls -la",
		Category: "gittools",
		Desc:     "list files",
	}
	got := Entry(in)
	if got != want {
		t.Errorf("Entry(%+v) = %+v, want %+v", in, got, want)
	}
}

func TestEntry_EmptyFields(t *testing.T) {
	got := Entry(types.AliasEntry{})
	if got != (types.AliasEntry{}) {
		t.Errorf("Entry(zero) = %+v, want zero", got)
	}
}
