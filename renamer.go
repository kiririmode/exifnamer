package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// FilenameByExif returns a EXIF's field embedded in an image file described by filepath.
// Which EXIF's field is used depends on an extractor implementation passed as arguments.
func FilenameByExif(path string, extractor Extractor) (string, error) {

	var reader ExifReader
	switch {
	case isJpeg(path):
		reader = &JpegExifReader{}
	case isHeic(path):
		reader = &HeicExifReader{}
	default:
		return "", fmt.Errorf("%s is not a target file format", path)
	}

	decoded, err := reader.Read(path)
	if err != nil {
		return "", fmt.Errorf("Reading %s for EXIF info: %s", path, err)
	}

	extracted, err := extractor.Extract(decoded)
	if err != nil {
		return "", fmt.Errorf("Extracting EXIF info from %s: %s", path, err)
	}

	return extracted + filepath.Ext(path), nil
}

// Renamer renames and moves/copies original file to an another directory.
// String origPath is a path of original file, locator specifies the destination
// directory, extractor generates the file name after move/copy.
type Renamer interface {
	Rename(origPath string, locator Locator, extractor Extractor, dryRun bool) error
}

// ReplaceRenamer replaces name of original file directly.
type ReplaceRenamer struct{}

// Rename renames and moves/copies an original file to an another directory.
// String origPath is a path of original file, locator specifies the destination
// directory, extractor generates a new file name.
func (r *ReplaceRenamer) Rename(origPath string, locator Locator, extractor Extractor, dryRun bool) error {
	newPath, err := createRenamedFilePath(origPath, locator, extractor)
	if err != nil {
		return err
	}
	log.Printf("%s --> %s\n", origPath, newPath)

	if dryRun {
		return nil
	}
	return os.Rename(origPath, newPath)
}

// CopyRenamer copies an original file to new directory with an new name.
// Original file is kept untouched.
type CopyRenamer struct{}

// Rename copies an original file to an another directory.
// String origPath is a path of original file, locator specifies the destination
// directory, extractor generates a new file name.
func (c *CopyRenamer) Rename(origPath string, locator Locator, extractor Extractor, dryRun bool) error {
	newPath, err := createRenamedFilePath(origPath, locator, extractor)

	if err != nil {
		return err
	}
	log.Printf("%s --> %s\n", origPath, newPath)

	if dryRun {
		return nil
	}
	return os.Link(origPath, newPath)
}

// createRenamedFilePath creates and returns new file path as a string.
func createRenamedFilePath(origPath string, locator Locator, extractor Extractor) (string, error) {
	destDir, err := locator.Dest(origPath)
	if err != nil {
		return "", err
	}

	fileName, err := FilenameByExif(origPath, extractor)
	if err != nil {
		return "", err
	}

	return filepath.Join(destDir, fileName), nil
}
