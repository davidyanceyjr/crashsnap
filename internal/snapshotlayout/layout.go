package snapshotlayout

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var errEmptyRoot = errors.New("snapshot root is required")

// Layout is the internal draft directory contract for snapshot storage.
// It is intentionally limited to path layout and does not imply manifest shape.
type Layout struct {
	root string
}

// Open validates a snapshot root and returns the draft layout rooted there.
func Open(root string) (Layout, error) {
	if root == "" {
		return Layout{}, errEmptyRoot
	}

	cleanRoot := filepath.Clean(root)
	info, err := os.Stat(cleanRoot)
	if err == nil && !info.IsDir() {
		return Layout{}, fmt.Errorf("snapshot root %q is not a directory", cleanRoot)
	}
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return Layout{}, fmt.Errorf("stat snapshot root %q: %w", cleanRoot, err)
	}

	return Layout{root: cleanRoot}, nil
}

func (l Layout) Root() string {
	return l.root
}

func (l Layout) MetadataDir() string {
	return filepath.Join(l.root, "metadata")
}

func (l Layout) RawDir() string {
	return filepath.Join(l.root, "raw")
}

func (l Layout) NormalizedDir() string {
	return filepath.Join(l.root, "normalized")
}

func (l Layout) ReportsDir() string {
	return filepath.Join(l.root, "reports")
}

func (l Layout) ArtifactsDir() string {
	return filepath.Join(l.root, "artifacts")
}

func (l Layout) Paths() []string {
	return []string{
		l.root,
		l.MetadataDir(),
		l.RawDir(),
		l.NormalizedDir(),
		l.ReportsDir(),
		l.ArtifactsDir(),
	}
}

// Materialize creates the draft snapshot directory structure on disk.
func (l Layout) Materialize() error {
	for _, path := range l.Paths() {
		if err := os.MkdirAll(path, 0o755); err != nil {
			return fmt.Errorf("create snapshot path %q: %w", path, err)
		}
	}

	return nil
}
