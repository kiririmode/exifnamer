package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

// jpegSuffix is Regexp for jpeg extension
var jpegSuffix = regexp.MustCompile(`^(?i)\.jpe?g$`)

// heicSuffix is Regexp for heic extension
var heicSuffix = regexp.MustCompile(`^(?i)\.heic$`)

// isJpeg returns whether a path has JPEG extension.
func isJpeg(path string) bool {
	return jpegSuffix.MatchString(filepath.Ext(path))
}

// isHeic returns whether a path has HEIC extension.
func isHeic(path string) bool {
	return heicSuffix.MatchString(filepath.Ext(path))
}

// IsTargetImage returns whether this program can treat a specified file.
func IsTargetImage(path string) bool {
	return isJpeg(path) || isHeic(path)
}

func main() {
	// log is directed to stderr, without timestamp
	log.SetOutput(os.Stderr)
	log.SetFlags(0)

	dryRun := flag.Bool("n", false, "dry-run")
	targetDir := flag.String("d", "", "copy destination directory")
	flag.Parse()

	// when the targetDir is specified, create it
	if !*dryRun && *targetDir != "" {
		if err := os.MkdirAll(*targetDir, 0755); err != nil {
			log.Fatal(err)
		}
	}

	renamer := &CopyRenamer{}
	extractor := &DateTimeExtractor{}
	// TODO: OriginalFileLocator should be replaced with FixedLocator
	var locator Locator = &OriginalFileLocator{}

	// targetDir is specified, all files should be located in that directory.
	if *targetDir != "" {
		locator = NewFixedLocator(*targetDir)
	}

	for _, dir := range flag.Args() {
		filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}
			if info.IsDir() || !IsTargetImage(path) {
				// skip when the path is directory, or is not a jpeg file.
				return nil
			}

			e := renamer.Rename(path, locator, extractor, *dryRun)
			if e != nil {
				// when error occured during rename, only copy to destination directory.
				log.Printf("error: %s. only copy to %s\n", e, *targetDir)
				if !*dryRun {
					os.Link(path, filepath.Join(*targetDir, filepath.Base(path)))
				}
			}
			// ignore error
			return nil
		})
	}
}
