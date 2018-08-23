package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"go4.org/media/heif"
)

// ExifReader reads EXIF from a file described by path.
type ExifReader interface {
	Read(path string) (*exif.Exif, error)
}

// JpegExifReader reads EXIF from a JPEG file.
type JpegExifReader struct{}

// Read reads EXIF from a path, and returns that EXIF data.
func (r *JpegExifReader) Read(path string) (*exif.Exif, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoded, err := exif.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("JPEG: Parsing %s for EXIF info: %s", path, err)
	}
	return decoded, nil
}

// HeicExifReader reads EXIF from a HEIC (HEIF) file.
type HeicExifReader struct{}

// Read reads EXIF from a path, and returns that EXIF data.
func (r *HeicExifReader) Read(path string) (*exif.Exif, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	b, err := heif.Open(f).EXIF()
	if err != nil {
		return nil, fmt.Errorf("HEIC: Extracting %s for EXIF info: %s", path, err)
	}

	decoded, err := exif.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("HEIC: Parsing %s for EXIF info: %s", path, err)
	}

	return decoded, nil
}
