package snapshotlayout

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestOpenReturnsExpectedPaths(t *testing.T) {
	root := filepath.Join(t.TempDir(), "incident-001")

	layout, err := Open(root)
	if err != nil {
		t.Fatalf("Open returned error: %v", err)
	}

	if layout.Root() != root {
		t.Fatalf("expected root %q, got %q", root, layout.Root())
	}

	wantPaths := []string{
		root,
		filepath.Join(root, "metadata"),
		filepath.Join(root, "raw"),
		filepath.Join(root, "normalized"),
		filepath.Join(root, "reports"),
		filepath.Join(root, "artifacts"),
	}

	gotPaths := layout.Paths()
	if len(gotPaths) != len(wantPaths) {
		t.Fatalf("expected %d paths, got %d", len(wantPaths), len(gotPaths))
	}

	for i := range wantPaths {
		if gotPaths[i] != wantPaths[i] {
			t.Fatalf("expected path %d to be %q, got %q", i, wantPaths[i], gotPaths[i])
		}
	}
}

func TestMaterializeCreatesDraftFixtureShape(t *testing.T) {
	root := filepath.Join(t.TempDir(), "incident-001")

	layout, err := Open(root)
	if err != nil {
		t.Fatalf("Open returned error: %v", err)
	}

	if err := layout.Materialize(); err != nil {
		t.Fatalf("Materialize returned error: %v", err)
	}

	for _, path := range layout.Paths() {
		info, err := os.Stat(path)
		if err != nil {
			t.Fatalf("expected %q to exist: %v", path, err)
		}
		if !info.IsDir() {
			t.Fatalf("expected %q to be a directory", path)
		}
	}
}

func TestOpenRejectsEmptyRoot(t *testing.T) {
	_, err := Open("")
	if !errors.Is(err, errEmptyRoot) {
		t.Fatalf("expected errEmptyRoot, got %v", err)
	}
}

func TestOpenRejectsFileRoot(t *testing.T) {
	tempDir := t.TempDir()
	root := filepath.Join(tempDir, "snapshot-root")
	if err := os.WriteFile(root, []byte("not a directory"), 0o644); err != nil {
		t.Fatalf("WriteFile returned error: %v", err)
	}

	_, err := Open(root)
	if err == nil {
		t.Fatal("expected Open to reject file root")
	}
}
