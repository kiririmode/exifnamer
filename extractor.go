package main

import (
	"github.com/rwcarlsen/goexif/exif"
)

// Extractor extracts EXIF information.
type Extractor interface {
	Extract(exif *exif.Exif) (string, error)
}

// DateTimeExtractor extract creation time of the photo.
type DateTimeExtractor struct{}

// Extract of DateTimeExtractor returns EXIF's "DateTimeOriginal" field.
// If not found, it tries the "DateTime" field.
func (e *DateTimeExtractor) Extract(exif *exif.Exif) (string, error) {
	datetime, err := exif.DateTime()
	if err != nil {
		return "", err
	}
	return datetime.Format("20060102-150405"), nil
}
