package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/rwcarlsen/goexif/exif"
)

func TestDateTimeExtractor_Extract_succeed(t *testing.T) {
	testCases := []struct {
		imageName string
		expected  string
	}{
		{imageName: "sample/DateTimeOriginal.jpg", expected: "20180102-030405"},
		// ModifyDate.jpg has no "DateTimeOriginal", so exif uses "ModifyDate"
		{imageName: "sample/ModifyDate.JPEG", expected: "20181231-123450"}}

	for _, tc := range testCases {
		decoded, err := loadExif(tc.imageName)
		if err != nil {
			t.Errorf("error should be nil but: %s", err)
		}

		sut := &DateTimeExtractor{}
		actual, err := sut.Extract(decoded)
		if err != nil {
			t.Errorf("error should be nil but: %s", err)
		}
		if actual != tc.expected {
			t.Errorf("expected: %s, but get: %s", tc.expected, actual)
		}
	}
}

func TestDateTimeExtractor_Extract_fail(t *testing.T) {
	testCases := []struct {
		imageName string
		expected  string
	}{
		{imageName: "sample/noDate.jpg", expected: "20181231-123450"}}

	for _, tc := range testCases {
		decoded, err := loadExif(tc.imageName)
		if err != nil {
			t.Errorf("error should be nil but: %s", err)
		}

		sut := &DateTimeExtractor{}
		_, err = sut.Extract(decoded)
		if !exif.IsTagNotPresentError(err) {
			t.Errorf("error should be nil but caught %T: %s", err, err)
		}
	}
}

func loadExif(path string) (*exif.Exif, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Errorf("error should be nil but: %s", err)
	}
	defer f.Close()

	return exif.Decode(f)
}
