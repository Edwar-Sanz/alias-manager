package amfile

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"am/src/types"
)

func TestGetFileContent(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "file.txt")
	want := "hello\nworld\n"
	if err := os.WriteFile(path, []byte(want), 0644); err != nil {
		t.Fatalf("setup write: %v", err)
	}

	got, err := GetFileContent(path)
	if err != nil {
		t.Fatalf("GetFileContent: %v", err)
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGetFileContent_Missing(t *testing.T) {
	_, err := GetFileContent(filepath.Join(t.TempDir(), "does-not-exist"))
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func TestConfigDir(t *testing.T) {
	dir, err := ConfigDir()
	if err != nil {
		t.Fatalf("ConfigDir: %v", err)
	}
	if dir == "" {
		t.Fatal("expected non-empty dir")
	}
	if filepath.Base(dir) != "am" {
		t.Errorf("expected base name 'am', got %q", filepath.Base(dir))
	}
}

func TestCreateIfNotExist_CreatesDirAndFile(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "nested", "am")

	path, err := CreateIfNotExist(dir)
	if err != nil {
		t.Fatalf("CreateIfNotExist: %v", err)
	}

	if _, err := os.Stat(dir); err != nil {
		t.Errorf("expected dir to exist: %v", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Errorf("expected file to exist: %v", err)
	}
	if filepath.Dir(path) != dir {
		t.Errorf("path %q not in dir %q", path, dir)
	}
}

func TestCreateIfNotExist_IdempotentPreservesContent(t *testing.T) {
	dir := t.TempDir()

	path, err := CreateIfNotExist(dir)
	if err != nil {
		t.Fatalf("first call: %v", err)
	}
	if err := os.WriteFile(path, []byte("keep me"), 0644); err != nil {
		t.Fatalf("write: %v", err)
	}

	path2, err := CreateIfNotExist(dir)
	if err != nil {
		t.Fatalf("second call: %v", err)
	}
	if path != path2 {
		t.Errorf("paths differ: %q vs %q", path, path2)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	if string(content) != "keep me" {
		t.Errorf("content was overwritten, got %q", content)
	}
}

func TestWriteAlias_AppendToEmptyFile(t *testing.T) {
	path := newAliasFile(t, "")

	entry := types.AliasEntry{Name: "ll", Command: "ls -la", Category: "fs", Desc: "list"}
	if err := WriteAlias(path, entry); err != nil {
		t.Fatalf("WriteAlias: %v", err)
	}

	content := readAll(t, path)
	want := `alias ll="ls -la"#category:fs#description:list` + "\n"
	if content != want {
		t.Errorf("content = %q, want %q", content, want)
	}
}

func TestWriteAlias_AppendsNewlineWhenMissing(t *testing.T) {
	path := newAliasFile(t, `alias g="git"#category:vcs#description:git`)

	entry := types.AliasEntry{Name: "ll", Command: "ls -la", Category: "fs", Desc: "list"}
	if err := WriteAlias(path, entry); err != nil {
		t.Fatalf("WriteAlias: %v", err)
	}

	content := readAll(t, path)
	lines := strings.Split(strings.TrimRight(content, "\n"), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d: %q", len(lines), content)
	}
	if !strings.Contains(lines[1], "ll=") {
		t.Errorf("expected second line to contain new alias, got %q", lines[1])
	}
}

func TestWriteAlias_UpdatesExisting(t *testing.T) {
	initial := `alias g="git"#category:vcs#description:old` + "\n" +
		`alias ll="ls"#category:fs#description:short` + "\n"
	path := newAliasFile(t, initial)

	updated := types.AliasEntry{Name: "g", Command: "git --no-pager", Category: "vcs", Desc: "new"}
	if err := WriteAlias(path, updated); err != nil {
		t.Fatalf("WriteAlias: %v", err)
	}

	content := readAll(t, path)
	if strings.Contains(content, "description:old") {
		t.Errorf("old description should be gone:\n%s", content)
	}
	if !strings.Contains(content, "git --no-pager") {
		t.Errorf("new command missing:\n%s", content)
	}
	if !strings.Contains(content, `alias ll="ls"`) {
		t.Errorf("untouched alias was lost:\n%s", content)
	}
	if got := strings.Count(content, "alias g="); got != 1 {
		t.Errorf("expected exactly 1 'alias g=' line, got %d", got)
	}
}

func TestWriteAlias_MissingFile(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "nope")
	err := WriteAlias(missing, types.AliasEntry{Name: "g", Command: "git"})
	if err == nil {
		t.Fatal("expected error when file is missing")
	}
}

func TestDeleteAlias_RemovesExisting(t *testing.T) {
	initial := `alias g="git"#category:vcs#description:git` + "\n" +
		`alias ll="ls"#category:fs#description:list` + "\n"
	path := newAliasFile(t, initial)

	if err := DeleteAlias(path, "g"); err != nil {
		t.Fatalf("DeleteAlias: %v", err)
	}

	content := readAll(t, path)
	if strings.Contains(content, "alias g=") {
		t.Errorf("alias g still present:\n%s", content)
	}
	if !strings.Contains(content, "alias ll=") {
		t.Errorf("alias ll should remain:\n%s", content)
	}
}

func TestDeleteAlias_NotFound(t *testing.T) {
	path := newAliasFile(t, `alias g="git"#category:vcs#description:git`+"\n")
	err := DeleteAlias(path, "missing")
	if err == nil {
		t.Fatal("expected error when alias not found")
	}
	if !strings.Contains(err.Error(), "missing") {
		t.Errorf("error should mention the alias name, got %v", err)
	}
}

func TestDeleteAlias_LastEntryProducesEmptyFile(t *testing.T) {
	path := newAliasFile(t, `alias g="git"#category:vcs#description:git`+"\n")
	if err := DeleteAlias(path, "g"); err != nil {
		t.Fatalf("DeleteAlias: %v", err)
	}
	content := readAll(t, path)
	if content != "" {
		t.Errorf("expected empty file, got %q", content)
	}
}

func TestDeleteAlias_MissingFile(t *testing.T) {
	err := DeleteAlias(filepath.Join(t.TempDir(), "nope"), "g")
	if err == nil {
		t.Fatal("expected error for missing file")
	}
}

func newAliasFile(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	path := filepath.Join(dir, ".amf")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	return path
}

func readAll(t *testing.T, path string) string {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read: %v", err)
	}
	return string(b)
}
