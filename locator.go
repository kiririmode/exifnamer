package main

import (
	"path/filepath"
)

// Locator returnes directory path where a file should be located.
// String path represents the path of the file.
type Locator interface {
	Dest(path string) (string, error)
}

// FixedLocator locates files to fixed directory destDir.
type FixedLocator struct {
	destDir string
}

// NewFixedLocator returns FixedLocator object.
func NewFixedLocator(destDir string) *FixedLocator {
	return &FixedLocator{destDir: destDir}
}

// Dest only returns destDir
func (l *FixedLocator) Dest(path string) (string, error) {
	return l.destDir, nil
}

// OriginalFileLocator locates the same direcory in which the files are.
type OriginalFileLocator struct{}

// Dest returns the same direcory in which the files are.
func (l *OriginalFileLocator) Dest(path string) (string, error) {
	return filepath.Dir(path), nil
}
