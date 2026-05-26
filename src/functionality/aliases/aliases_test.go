package aliases

import (
	"regexp"
	"strings"
	"testing"

	"am/src/constants"
	"am/src/types"
)

var ansiRe = regexp.MustCompile(`\x1b\[[0-9;:]*[a-zA-Z]`)

func stripANSI(s string) string {
	return ansiRe.ReplaceAllString(s, "")
}

func TestParseLine_Valid(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want types.AliasEntry
	}{
		{
			name: "double quoted command",
			in:   `alias ll="ls -la"#category:git#description:list files`,
			want: types.AliasEntry{Name: "ll", Command: "ls -la", Category: "git", Desc: "list files"},
		},
		{
			name: "single quoted command",
			in:   `alias ll='ls -la'#category:git#description:list files`,
			want: types.AliasEntry{Name: "ll", Command: "ls -la", Category: "git", Desc: "list files"},
		},
		{
			name: "unquoted command",
			in:   `alias g=git#category:vcs#description:git shortcut`,
			want: types.AliasEntry{Name: "g", Command: "git", Category: "vcs", Desc: "git shortcut"},
		},
		{
			name: "empty category becomes Uncategorized",
			in:   `alias g="git"#category:#description:no cat`,
			want: types.AliasEntry{Name: "g", Command: "git", Category: constants.UnCategorizedCategory, Desc: "no cat"},
		},
		{
			name: "empty description allowed",
			in:   `alias g="git"#category:vcs#description:`,
			want: types.AliasEntry{Name: "g", Command: "git", Category: "vcs", Desc: ""},
		},
		{
			name: "leading and trailing whitespace on line is trimmed",
			in:   "   alias g=\"git\"#category:vcs#description:desc   ",
			want: types.AliasEntry{Name: "g", Command: "git", Category: "vcs", Desc: "desc"},
		},
		{
			name: "space before category marker (as produced by BuildLine)",
			in:   `alias ll="ls -la" #category:git#description:list`,
			want: types.AliasEntry{Name: "ll", Command: "ls -la", Category: "git", Desc: "list"},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, ok := ParseLine(tc.in)
			if !ok {
				t.Fatalf("ParseLine(%q) returned ok=false", tc.in)
			}
			if got != tc.want {
				t.Errorf("ParseLine(%q)\n got = %+v\nwant = %+v", tc.in, got, tc.want)
			}
		})
	}
}

func TestParseLine_Invalid(t *testing.T) {
	tests := []struct {
		name string
		in   string
	}{
		{"empty", ""},
		{"plain text", "hello world"},
		{"no alias prefix", `ll="ls -la"#category:git#description:x`},
		{"no category marker", `alias ll="ls -la"#description:x`},
		{"no description marker", `alias ll="ls -la"#category:git`},
		{"no equals sign", `alias ll#category:git#description:x`},
		{"comment line", `#category:git`},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, ok := ParseLine(tc.in)
			if ok {
				t.Errorf("ParseLine(%q) returned ok=true, want false", tc.in)
			}
		})
	}
}

func TestBuildLine(t *testing.T) {
	tests := []struct {
		name  string
		entry types.AliasEntry
		want  string
	}{
		{
			name:  "double quoted by default",
			entry: types.AliasEntry{Name: "ll", Command: "ls -la", Category: "git", Desc: "list files"},
			want:  `alias ll="ls -la"#category:git#description:list files`,
		},
		{
			name:  "single quotes when command contains double quote",
			entry: types.AliasEntry{Name: "g", Command: `echo "hi"`, Category: "misc", Desc: "test"},
			want:  `alias g='echo "hi"'#category:misc#description:test`,
		},
		{
			name:  "empty command",
			entry: types.AliasEntry{Name: "x", Command: "", Category: "c", Desc: "d"},
			want:  `alias x=""#category:c#description:d`,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := BuildLine(tc.entry)
			if got != tc.want {
				t.Errorf("BuildLine(%+v) = %q, want %q", tc.entry, got, tc.want)
			}
		})
	}
}

func TestBuildLineParseLineRoundTrip(t *testing.T) {
	entries := []types.AliasEntry{
		{Name: "ll", Command: "ls -la", Category: "git", Desc: "list files"},
		{Name: "g", Command: "git", Category: "vcs", Desc: "shortcut"},
		{Name: "say", Command: `echo "hi"`, Category: "misc", Desc: "with quotes"},
	}
	for _, e := range entries {
		t.Run(e.Name, func(t *testing.T) {
			line := BuildLine(e)
			got, ok := ParseLine(line)
			if !ok {
				t.Fatalf("ParseLine failed for built line %q", line)
			}
			if got != e {
				t.Errorf("round-trip mismatch:\n got = %+v\nwant = %+v\nline = %q", got, e, line)
			}
		})
	}
}

func TestParseAliases_Empty(t *testing.T) {
	parsed := ParseAliases("")
	if len(parsed.Categories) != 0 {
		t.Errorf("expected no categories, got %v", parsed.Categories)
	}
	if len(parsed.ByCategory) != 0 {
		t.Errorf("expected empty ByCategory, got %v", parsed.ByCategory)
	}
}

func TestParseAliases_MultipleEntries(t *testing.T) {
	content := strings.Join([]string{
		`alias ll="ls -la"#category:fs#description:long list`,
		`alias g="git"#category:vcs#description:git`,
		`# random comment`,
		`alias gs="git status"#category:vcs#description:status`,
		``,
	}, "\n")

	parsed := ParseAliases(content)

	if got, want := parsed.Categories, []string{"fs", "vcs"}; !equalSlices(got, want) {
		t.Errorf("Categories = %v, want %v", got, want)
	}
	if got := len(parsed.ByCategory["fs"]); got != 1 {
		t.Errorf("fs entries = %d, want 1", got)
	}
	if got := len(parsed.ByCategory["vcs"]); got != 2 {
		t.Errorf("vcs entries = %d, want 2", got)
	}
	if parsed.ByCategory["vcs"][1].Name != "gs" {
		t.Errorf("expected second vcs entry to be 'gs', got %q", parsed.ByCategory["vcs"][1].Name)
	}
}

func TestParseAliases_PreservesCategoryOrder(t *testing.T) {
	content := strings.Join([]string{
		`alias a="x"#category:zeta#description:`,
		`alias b="x"#category:alpha#description:`,
		`alias c="x"#category:zeta#description:`,
	}, "\n")
	parsed := ParseAliases(content)
	if got, want := parsed.Categories, []string{"zeta", "alpha"}; !equalSlices(got, want) {
		t.Errorf("Categories = %v, want %v (insertion order)", got, want)
	}
}

func TestFormatCategory_ContainsEntries(t *testing.T) {
	entries := []types.AliasEntry{
		{Name: "ll", Command: "ls -la", Desc: "list"},
		{Name: "gs", Command: "git status", Desc: "status"},
	}
	out := stripANSI(FormatCategory("tools", entries))
	for _, s := range []string{"tools", "ll", "ls -la", "list", "gs", "git status", "status"} {
		if !strings.Contains(out, s) {
			t.Errorf("FormatCategory output missing %q\noutput: %s", s, out)
		}
	}
}

func TestFormatCategories_ContainsAll(t *testing.T) {
	out := stripANSI(FormatCategories([]string{"git", "fs", "misc"}))
	for _, s := range []string{"git", "fs", "misc"} {
		if !strings.Contains(out, s) {
			t.Errorf("FormatCategories output missing %q", s)
		}
	}
}

func TestFormatCategories_Empty(t *testing.T) {
	out := FormatCategories(nil)
	if out != "" {
		t.Errorf("expected empty string for nil categories, got %q", out)
	}
}

func TestFormatAllAliases_ContainsSeparatorsAndCategories(t *testing.T) {
	parsed := types.ParsedAliases{
		Categories: []string{"a", "b"},
		ByCategory: map[string][]types.AliasEntry{
			"a": {{Name: "n1", Command: "c1", Desc: "d1"}},
			"b": {{Name: "n2", Command: "c2", Desc: "d2"}},
		},
	}
	rawOut := FormatAllAliases(parsed)
	out := stripANSI(rawOut)
	if !strings.Contains(out, constants.Separator) {
		t.Errorf("expected separator in output")
	}
	for _, s := range []string{"a", "b", "n1", "c1", "d1", "n2", "c2", "d2"} {
		if !strings.Contains(out, s) {
			t.Errorf("FormatAllAliases output missing %q", s)
		}
	}
	if got := strings.Count(out, constants.Separator); got != 3 {
		t.Errorf("expected 3 separators (1 header + 1 after each of 2 categories), got %d", got)
	}
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
