package main

import (
	"fmt"
	"testing"

	"github.com/rwcarlsen/goexif/exif"
)

func TestFilenameByExif_openFail(t *testing.T) {
	_, err := FilenameByExif("not-exist", &DateTimeExtractor{})
	if err == nil {
		t.Error("FilenameByExif should return error when the specified file is not found")
	}
}

func TestFilenameByExif_parseFail(t *testing.T) {
	_, err := FilenameByExif("renamer_test.go", &DateTimeExtractor{})
	if err == nil {
		t.Error("FilenameByExif should return error when it cannot parse EXIF")
	}
}

type FailExtractor struct{}

func (f *FailExtractor) Extract(exif *exif.Exif) (string, error) {
	return "", fmt.Errorf("error")
}
func TestFilenameByExif_extractFail(t *testing.T) {
	extractor := &FailExtractor{}
	_, err := FilenameByExif("sample/DateTimeOriginal.jpg", extractor)

	if err == nil {
		t.Error("FilenameByExif should return error when extraction fails")
	}
}

func TestFilenameByExif(t *testing.T) {
	testCases := []struct {
		imageFile string
		expected  string
	}{
		{imageFile: "sample/DateTimeOriginal.jpg", expected: "20180102-030405.jpg"},
		{imageFile: "sample/ModifyDate.JPEG", expected: "20181231-123450.JPEG"}}

	for _, tc := range testCases {
		actual, err := FilenameByExif(tc.imageFile, &DateTimeExtractor{})

		if err != nil {
			t.Errorf("error should be nil but get: %s", err)
		}
		if actual != tc.expected {
			t.Errorf("expected: %s, but get: %s", tc.expected, actual)
		}
	}
}

func Test_createRenamedFilePath(t *testing.T) {
	testCases := []struct {
		imageFile string
		expected  string
	}{
		{imageFile: "sample/DateTimeOriginal.jpg", expected: "fixed/20180102-030405.jpg"},
		{imageFile: "sample/ModifyDate.JPEG", expected: "fixed/20181231-123450.JPEG"}}

	for _, tc := range testCases {
		actual, err := createRenamedFilePath(
			tc.imageFile, NewFixedLocator("fixed"), &DateTimeExtractor{})

		if err != nil {
			t.Errorf("error should be nil but get: %s", err)
		}
		if actual != tc.expected {
			t.Errorf("expected: %s, but get: %s", tc.expected, actual)
		}
	}
}
